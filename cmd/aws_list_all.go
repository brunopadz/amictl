package cmd

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/brunopadz/amictl/commons"
	aws2 "github.com/brunopadz/amictl/providers/aws"
	"github.com/spf13/cobra"
)

func init() {
	awsCmd.AddCommand(listAll)
}

// listAmiCmd represents the listAmi command
var listAll = &cobra.Command{
	Use:     "list-all",
	Short:   "List all AMIs",
	Long:    `List all AMIs for a given region and account.`,
	Example: `  amictl aws list-all --account 123456789012 --region us-east-1`,
	RunE:    runAll,
}

func runAll(cmd *cobra.Command, args []string) error {
	account, err := cmd.Flags().GetString("account")
	if err != nil {
		return err
	}

	// Creates a input filter to get AMIs
	f := &ec2.DescribeImagesInput{
		Owners: []*string{
			aws.String(account),
		},
	}

	region, err := cmd.Flags().GetString("region")
	if err != nil {
		return err
	}

	// Establishes new authenticated session to AWS
	sess, err := aws2.NewSession(region)
	if err != nil {
		return err
	}

	// Filter AMIs based on input filter
	a, err := sess.DescribeImages(f)
	if err != nil {
		return err
	}

	l := aws2.ListAll(a)
	r := strings.Join(l, "\n")

	cost, err := cmd.Flags().GetBool("cost")
	if err != nil {
		return err
	}

	if cost == true {
		var total float64
		for _, ami := range a.Images {
			s := aws.Int64Value(ami.BlockDeviceMappings[0].Ebs.VolumeSize)
			i := aws.StringValue(ami.ImageId)
			p := aws2.GetAmiPriceBySize(s, region)

			total += p
			cmd.Println("AMI-ID:", i, "Size:", s, "GB", "Estimated cost monthly: U$", commons.Round(p))
		}
		rt := commons.Round(total)
		cmd.Println("\nEstimated cost monthly: U$", rt, "for", len(l), "AMIs")
	} else {
		cmd.Println(r)
		cmd.Println("Total of", len(l), "AMIs")
	}

	return nil
}

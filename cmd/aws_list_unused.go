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
	awsCmd.AddCommand(listUnused)
}

// listAmiCmd represents the listAmi command
var listUnused = &cobra.Command{
	Use:     "list-unused",
	Short:   "List unused AMIs",
	Long:    `List not used AMIs for a given region and account.`,
	Example: `  amictl aws list-unused --account 123456789012 --region us-east-1`,
	RunE:    runUnused,
}

func runUnused(cmd *cobra.Command, args []string) error {
	account, err := cmd.Flags().GetString("account")
	if err != nil {
		return err
	}

	// Creates a input filter to get AMIs
	f := &ec2.DescribeImagesInput{
		Owners: []*string{
			&account,
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

	// Compare AMI list
	l, u, err := aws2.ListNotUsed(a, sess)
	if err != nil {
		return err
	}

	n := commons.Compare(l, u)
	r := strings.Join(n, "\n")

	cost, err := cmd.Flags().GetBool("cost")
	if err != nil {
		return err
	}

	if cost == true {
		var total float64
		for _, i := range n {
			for _, ami := range a.Images {
				if aws.StringValue(ami.ImageId) == i {
					s := aws.Int64Value(ami.BlockDeviceMappings[0].Ebs.VolumeSize)
					p := aws2.GetAmiPriceBySize(s, region)
					total += p
					cmd.Println("AMI-ID:", i, "Size:", s, "GB", "Estimated cost monthly: U$", commons.Round(p))
				}
			}
		}
		rt := commons.Round(total)
		cmd.Println("\nEstimated cost monthly: U$", rt, "for", len(n), "Unused AMIs")
	} else {
		cmd.Println(r)
		cmd.Println("Total of", len(n), "not used AMIs")
	}

	return nil
}

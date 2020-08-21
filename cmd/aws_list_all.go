package cmd

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
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

func runAll(cmd *cobra.Command, _ []string) error {
	account, err := cmd.Flags().GetString("account")
	if err != nil {
		return err
	}

	// Creates describeImagesOutput input filter to get AMIs
	criteria := &ec2.DescribeImagesInput{
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
	describeImagesOutput, err := sess.DescribeImages(criteria)
	if err != nil {
		return err
	}

	imageList := aws2.ListAll(describeImagesOutput)

	cost, err := cmd.Flags().GetBool("cost")
	if err != nil {
		return err
	}

	if cost == true {
		var partial float64
		for _, ami := range imageList {
			p := aws2.GetAmiPriceBySize(ami.Size, region)
			cmd.Println("ami-id:", ami.ID, "size:", ami.Size, "GB", "Estimated cost monthly: U$", aws2.Round(p))
			partial += p
		}

		total := aws2.Round(partial)
		cmd.Println("\nEstimated cost monthly: U$", total, "for", len(imageList), "AMIs")
	} else {
		for _, ami := range imageList {
			cmd.Println("ami-id:", ami.ID)
		}

		cmd.Println("Total of", len(imageList), "AMIs")
	}

	return nil
}

package cmd

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	provider "github.com/brunopadz/amictl/providers/aws"
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
	cost, err := cmd.Flags().GetBool("cost")
	if err != nil {
		return err
	}

	account, err := cmd.Flags().GetString("account")
	if err != nil {
		return err
	}

	region, err := cmd.Flags().GetString("region")
	if err != nil {
		return err
	}

	// Establishes new authenticated session to AWS
	sess, err := provider.NewSession(region)
	if err != nil {
		return err
	}

	// Creates describeImagesOutput input filter to get AMIs
	var criteria = &ec2.DescribeImagesInput{
		Owners: []*string{
			aws.String(account),
		},
	}

	// Filter AMIs based on input filter
	describeImagesOutput, err := sess.DescribeImages(criteria)
	if err != nil {
		return err
	}

	var volume int64

	for _, ami := range describeImagesOutput.Images {
		cmd.Printf("ami-id: %s ", aws.StringValue(ami.ImageId))
		if cost {
			var volumeSize 	= aws.Int64Value(ami.BlockDeviceMappings[0].Ebs.VolumeSize)
			volume += volumeSize

			cmd.Printf("size: %d GB ", volumeSize)
			cmd.Printf("monthly cost: U$ %g", provider.GetAmiPriceBySize(volumeSize, region))
		}
		cmd.Println()
	}

	cmd.Printf("Total of AMIs: %d \n", len(describeImagesOutput.Images))
	if cost {
		cmd.Printf("Estimated cost monthly: U$ %g", provider.GetAmiPriceBySize(volume, region))
	}


	return nil
}

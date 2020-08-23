package cmd

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	provider "github.com/brunopadz/amictl/providers/aws"
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

func runUnused(cmd *cobra.Command, _ []string) error {
	account, err := cmd.Flags().GetString("account")
	if err != nil {
		return err
	}

	// Creates DescribeImagesInput to get AMIs
	var criteria = &ec2.DescribeImagesInput{
		Owners: []*string{
			&account,
		},
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

	// Filter AMIs based on criteria filter
	output, err := sess.DescribeImages(criteria)
	if err != nil {
		return err
	}

	// Filter AMI with reservations
	err = provider.FilterAmiInUse(sess, output)
	if err != nil {
		return err
	}

	err = provider.Render(cmd, region, output)
	if err != nil {
		return err
	}

	return nil
}


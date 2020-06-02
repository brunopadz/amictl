package cmd

import (
	"amictl/providers"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
)

// listAmiCmd represents the listAmi command
var listUnused = &cobra.Command{
	Use:   "list-unused",
	Short: "List unused AMIs",
	Long:  `List not used AMIs for a given region and account.`,
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// Creates a input filter to get AMIs
		f := &ec2.DescribeImagesInput{
			Owners: []*string{
				aws.String(args[0]),
			},
		}

		// Establishes new authenticated session to AWS
		s := providers.AwsSession(args[1])

		// Filter AMIs based on input filter
		a, err := s.DescribeImages(f)
		if err != nil {
			fmt.Println(err)
		}

		// Compare AMI list
		l := providers.AwsListNotUsed(a, s)
		fmt.Println(l)

	},
}

func init() {
	awsCmd.AddCommand(listUnused)
}

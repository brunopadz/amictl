package cmd

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/brunopadz/amictl/providers"
	"github.com/spf13/cobra"
)

// listAmiCmd represents the listAmi command
var listAll = &cobra.Command{
	Use:   "list-all",
	Short: "List all AMIs",
	Long:  `List all AMIs for a given region and account.`,
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

		l := providers.AwsListAll(a, s)

		r := strings.Join(l, "\n")
		fmt.Println(r)
		fmt.Println("Total of", len(l), "AMIs")

	},
}

func init() {
	awsCmd.AddCommand(listAll)
}

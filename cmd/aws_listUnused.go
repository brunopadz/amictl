package cmd

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/brunopadz/amictl/commons"
	"github.com/brunopadz/amictl/providers"
	"github.com/spf13/cobra"
)

var listUsedCommand = &cobra.Command{
	Use:   "list-unused",
	Short: "List unused AMIs",
	Long:  "List not used AMIs for a given region and account.",
	Args:  	cobra.MaximumNArgs(2),
	Run: 	ListUnusedCommand(),
}

func init() {
	awsCmd.AddCommand(listUsedCommand)
}

// ListUnusedCommand return a callable to list unused ami
func ListUnusedCommand() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		var accountID = aws.String(args[0])
		var region = aws.String(args[1])

		// Creates a input filter to get AMIs
		f := &ec2.DescribeImagesInput{
			Owners: []*string{
				accountID,
			},
		}

		// Establishes new authenticated session to AWS
		s := providers.AwsSession(region)

		// Filter AMIs based on input filter
		a, err := s.DescribeImages(f)
		if err != nil {
			fmt.Println(err)
		}

		// Compare AMI list
		l, u := providers.AwsListNotUsed(a, s)

		n := commons.Compare(l, u)
		r := strings.Join(n, "\n")

		fmt.Println(r)
		fmt.Println("Total of", len(n), "not used AMIs")
	}
}

package cmd

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/brunopadz/amictl/providers"
	"github.com/spf13/cobra"
	"strings"
)

var listAllCommand = &cobra.Command{
	Use:   "list-all",
	Short: "List all AMIs",
	Long:    `List all AMIs for a given region and account.`,
	Example: `  amictl aws list-all --account 123456789012 --region us-east-1`,
	Args:  cobra.MaximumNArgs(2),
	RunE:   ListAllCommand(),
}

func init() {
	awsCmd.AddCommand(listAllCommand)
}

// ListAllCommand return a callable to list all ami
func ListAllCommand() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
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
      return err
		}

		l := providers.AwsListAll(a)
		r := strings.Join(l, "\n")

		fmt.Println(r)
		fmt.Println("Total of", len(l), "AMIs")
    
    return nil
	}
}

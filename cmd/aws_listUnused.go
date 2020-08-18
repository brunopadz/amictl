package cmd

import (
	"fmt"
	aws2 "github.com/brunopadz/amictl/providers/aws"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/brunopadz/amictl/commons"
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
	// Creates a input filter to get AMIs
	f := &ec2.DescribeImagesInput{
		Owners: []*string{
			aws.String(account),
		},
	}

	// Establishes new authenticated session to AWS
	s := aws2.Session(region)

	// Filter AMIs based on input filter
	a, err := s.DescribeImages(f)
	if err != nil {
		fmt.Println(err)
	}

	// Compare AMI list
	l, u := aws2.ListNotUsed(a, s)

	n := commons.Compare(l, u)
	r := strings.Join(n, "\n")

	if cost == true {
		var total float64
		for _, i := range n {
			for _, ami := range a.Images {
				if aws.StringValue(ami.ImageId) == i {
					s := aws.Int64Value(ami.BlockDeviceMappings[0].Ebs.VolumeSize)
					p := aws2.GetAmiPriceBySize(s, region)
					total += p
					fmt.Println("AMI-ID:", i, "Size:", s, "GB", "Estimated cost monthly: U$", commons.Round(p))
				}
			}
		}
		rt := commons.Round(total)
		fmt.Println("\nEstimated cost monthly: U$", rt, "for", len(n), "Unused AMIs")
	} else {
		fmt.Println(r)
		fmt.Println("Total of", len(n), "not used AMIs")
	}

	return nil
}

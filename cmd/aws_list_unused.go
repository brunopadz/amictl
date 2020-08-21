package cmd

import (
	"github.com/aws/aws-sdk-go/service/ec2"
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

func runUnused(cmd *cobra.Command, _ []string) error {
	account, err := cmd.Flags().GetString("account")
	if err != nil {
		return err
	}

	// Creates describeImagesOutput input filter to get AMIs
	criteria := &ec2.DescribeImagesInput{
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
	describeImagesOutput, err := sess.DescribeImages(criteria)
	if err != nil {
		return err
	}

	imageList := aws2.ListAll(describeImagesOutput)

	// compare AMI list
	imageInUseList, err := aws2.ListNotUsed(sess, imageList)
	if err != nil {
		return err
	}

	imageNotInUseList := aws2.Compare(imageList, imageInUseList)

	cost, err := cmd.Flags().GetBool("cost")
	if err != nil {
		return err
	}

	if cost == true {
		var total float64
		for _, imageNotInUse := range imageNotInUseList {
			for _, imageInUse := range imageInUseList {
				if imageNotInUse.ID == imageInUse.ID {
					p := aws2.GetAmiPriceBySize(imageNotInUse.Size, region)
					total += p
					cmd.Println("ami-id:", imageNotInUse.ID, "size:", imageNotInUse.Size, "GB", "Estimated cost monthly: U$", aws2.Round(p))
				}
			}
		}

		rt := aws2.Round(total)
		cmd.Println("\nEstimated cost monthly: U$", rt, "for", len(imageNotInUseList), "Unused AMIs")
	} else {
		for _, image := range imageNotInUseList {
			cmd.Println("ami-id:", image.ID)
		}

		cmd.Println("Total of", len(imageNotInUseList), "not used AMIs")
	}

	return nil
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func Render(cmd *cobra.Command, region string, describeImagesOutput *ec2.DescribeImagesOutput) error {
	var volume int64

	red := color.New(color.FgRed)
	yellow := color.New(color.FgYellow)
	green := color.New(color.FgGreen)
	title := color.New(color.Bold)

	okmsg := "\nEverything looks good."
	warnmsg := "\nYou could be saving some money."
	alertmsg := "\nGo ahead and delete them to save some money."
	getcost := "Try running it again with --cost flag."

	cost, err := cmd.Flags().GetBool("cost")
	if err != nil {
		return err
	}

	if cost {
		title.Println("AMI ID\t\t\tSize\tMonthly Cost")
		for _, ami := range describeImagesOutput.Images {
			var volumeSize = aws.Int64Value(ami.BlockDeviceMappings[0].Ebs.VolumeSize)
			volume += volumeSize

			cmd.Printf("%s", aws.StringValue(ami.ImageId))
			cmd.Printf("\t%d GB ", volumeSize)
			cmd.Printf("\tUS$ %g\n", GetAmiPriceBySize(volumeSize, region))
		}
		cmd.Println("Estimated cost monthly: US$", GetAmiPriceBySize(volume, region))
	} else {
		for _, ami := range describeImagesOutput.Images {
			cmd.Println(aws.StringValue(ami.ImageId))
		}
		if len(describeImagesOutput.Images) == 0 {
			green.Println("Total of ", len(describeImagesOutput.Images), "AMIs.", okmsg, getcost)
		} else if len(describeImagesOutput.Images) >= 1 && len(describeImagesOutput.Images) <= 10 {
			yellow.Println("Total of ", len(describeImagesOutput.Images), "AMIs.", warnmsg, getcost)
		} else {
			red.Println("Total of ", len(describeImagesOutput.Images), "AMIs.", alertmsg, getcost)
		}
	}

	return nil
}

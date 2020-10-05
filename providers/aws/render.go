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

	okmsg := "\nEverything looks good."
	warnmsg := "\nYou could be saving some money."
	alertmsg := "\nGo ahead and delete them to save some money."
	getcost := "Try running it again with --cost flag."

	cost, err := cmd.Flags().GetBool("cost")
	if err != nil {
		return err
	}

	for _, ami := range describeImagesOutput.Images {
		cmd.Printf("%s", aws.StringValue(ami.ImageId))
		if cost {
			var volumeSize = aws.Int64Value(ami.BlockDeviceMappings[0].Ebs.VolumeSize)
			volume += volumeSize

			cmd.Printf("\t size: %d GB ", volumeSize)
			cmd.Printf("\t monthly cost: U$ %g", GetAmiPriceBySize(volumeSize, region))
		}
		cmd.Println()
	}

	if len(describeImagesOutput.Images) == 0 {
		green.Println("Total of ", len(describeImagesOutput.Images), "AMIs.", okmsg, getcost)
	} else if len(describeImagesOutput.Images) >= 1 && len(describeImagesOutput.Images) <= 10 {
		yellow.Println("Total of ", len(describeImagesOutput.Images), "AMIs.", warnmsg, getcost)
	} else {
		red.Println("Total of ", len(describeImagesOutput.Images), "AMIs.", alertmsg, getcost)
	}

	if cost {
		cmd.Printf("Estimated cost monthly: U$ %g \n", GetAmiPriceBySize(volume, region))
	}

	return nil
}

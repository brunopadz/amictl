package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
)

func Render(cmd *cobra.Command, region string, describeImagesOutput *ec2.DescribeImagesOutput) error {
	var volume int64

	cost, err := cmd.Flags().GetBool("cost")
	if err != nil {
		return err
	}

	for _, ami := range describeImagesOutput.Images {
		cmd.Printf("ami-id: %s ", aws.StringValue(ami.ImageId))
		if cost {
			var volumeSize = aws.Int64Value(ami.BlockDeviceMappings[0].Ebs.VolumeSize)
			volume += volumeSize

			cmd.Printf("size: %d GB ", volumeSize)
			cmd.Printf("monthly cost: U$ %g", GetAmiPriceBySize(volumeSize, region))
		}
		cmd.Println()
	}

	cmd.Println("Total of AMIs: ", len(describeImagesOutput.Images))
	if cost {
		cmd.Printf("Estimated cost monthly: U$ %g \n", GetAmiPriceBySize(volume, region))
	}

	return nil
}

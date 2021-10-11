package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/pterm/pterm"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	cfg "github.com/brunopadz/amictl/config"
	aaws "github.com/brunopadz/amictl/pkg/providers/aws"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	awsCmd.AddCommand(inspect)
}

var inspect = &cobra.Command{
	Use:     "inspect",
	Short:   "Inspect AMI",
	Long:    `Inspect command shows additional info about AMIs`,
	Example: `  amictl aws inspect`,
	RunE:    runInspect,
}

var (
	ami    string
	region string
)

func runInspect(cmd *cobra.Command, _ []string) error {

	var c cfg.Config

	err := viper.Unmarshal(&c)
	if err != nil {
		fmt.Println(err)
	}

	s, err := aaws.New(region)
	if err != nil {
		log.Fatalln("Couldn't create a session to AWS.")
	}

	a := ec2.NewFromConfig(s)

	input := &ec2.DescribeImagesInput{
		ImageIds: []string{
			ami,
		},
		Owners: []string{
			viper.GetString("aws.account"),
		},
	}

	o, err := a.DescribeImages(context.TODO(), input)
	if err != nil {
		log.Fatalln("Couldn't get AMI data.")
	}

	for _, v := range o.Images {
		pterm.FgLightCyan.Println("Displaying info for:", pterm.NewStyle(pterm.Bold).Sprint(aws.ToString(v.ImageId)))
		pterm.FgDarkGray.Println("----------------------------------------------")
		fmt.Println("Name:", aws.ToString(v.Name))
		fmt.Println("Description:", aws.ToString(v.Description))
		fmt.Println("Creation Date:", aws.ToString(v.CreationDate))
		fmt.Println("Deprecation Time:", aws.ToString(v.DeprecationTime))
		fmt.Println("Root Device Name:", aws.ToString(v.RootDeviceName))
		fmt.Println("RAM Disk ID:", aws.ToString(v.RamdiskId))
		fmt.Println("Kernel ID:", aws.ToString(v.KernelId))
		fmt.Println("Platform Details:", aws.ToString(v.PlatformDetails))
		fmt.Println("Public:", aws.ToBool(v.Public))
		fmt.Println("Tags:")
		for _, t := range v.Tags {
			fmt.Println(" ", aws.ToString(t.Key), "=", aws.ToString(t.Value))
		}

	}

	return nil
}

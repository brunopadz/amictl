package cmd

import (
	"context"
	"fmt"

	"github.com/pterm/pterm"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	cfg "github.com/brunopadz/amictl/config"
	aaws "github.com/brunopadz/amictl/pkg/providers/aws"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	awsCmd.AddCommand(deregister)
}

var deregister = &cobra.Command{
	Use:     "deregister",
	Short:   "Deregister a single AMI.",
	Long:    `Deregister command deregisters / deletes a single AMI.`,
	Example: `  amictl aws deregister --region us-east-1 --ami ami-0x00000000f`,
	RunE:    runDelete,
}

func runDelete(cmd *cobra.Command, _ []string) error {

	var c cfg.Config

	err := viper.Unmarshal(&c)
	if err != nil {
		fmt.Println(err)
	}

	s, err := aaws.New(region)
	if err != nil {
		fmt.Println("Couldn't create a session to AWS. Please check your credentials.")
	}

	a := ec2.NewFromConfig(s)

	i := &ec2.DeregisterImageInput{
		ImageId: aws.String(ami),
	}

	_, err = a.DeregisterImage(context.TODO(), i)
	if err != nil {
		fmt.Println("Couldn't describe AMIs.")
	}

	pterm.FgLightCyan.Println("Image deregistered:", pterm.NewStyle(pterm.Bold).Sprint(aws.ToString(i.ImageId)))

	return err
}

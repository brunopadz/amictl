package cmd

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	cfg "github.com/brunopadz/amictl/config"
	aaws "github.com/brunopadz/amictl/pkg/providers/aws"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	awsCmd.AddCommand(delete)
}

var delete = &cobra.Command{
	Use:     "delete",
	Short:   "Delete a single AMI.",
	Long:    `Delete command deregister a single AMI. Check the docs for more info.`,
	Example: `  amictl aws delete --region 111222333444 --ami ami-0x00000000f`,
	RunE:    runDelete,
}

//var dami *string

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

	w, err := a.DeregisterImage(context.TODO(), i)
	if err != nil {
		fmt.Println("Couldn't describe AMIs.")
	}

	fmt.Println(w)
	//fmt.Println(o.ResultMetadata)

	//for _, v := range o.ResultMetadata {
	//
	//}

	return err
}

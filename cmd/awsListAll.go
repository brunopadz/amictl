package cmd

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"

	cfg "github.com/brunopadz/amictl/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	awsCmd.AddCommand(listAll)
}

var listAll = &cobra.Command{
	Use:     "list-all",
	Short:   "List all AMIs",
	Long:    `List all AMIs for a given region and account.`,
	Example: `  amictl aws list-all --account 123456789012 --region us-east-1`,
	RunE:    runAll,
}

func runAll(cmd *cobra.Command, _ []string) error {

	var c cfg.Config

	err := viper.Unmarshal(&c)
	if err != nil {
		fmt.Println(err)
	}

	sess, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(viper.GetString("aws.region")))

	client := ec2.NewFromConfig(sess)

	input := &ec2.DescribeImagesInput{
		Owners: []string{
			viper.GetString("aws.account"),
		},
	}

	output, err := client.DescribeImages(context.TODO(), input)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range output.Images {
		fmt.Println(aws.ToString(v.ImageId))
	}

	return nil
}

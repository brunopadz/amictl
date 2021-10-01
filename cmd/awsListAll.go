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

	r := viper.GetStringSlice("aws.regions")

	for _, v := range r {
		sess, err := config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(v))

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

		fmt.Println("Listing AMIs in", v)

		for _, i := range output.Images {
			fmt.Println(aws.ToString(i.ImageId))
		}
		fmt.Println("Total of", len(output.Images), "AMIs in", v)
	}

	return nil
}

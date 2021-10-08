package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/pterm/pterm"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	cfg "github.com/brunopadz/amictl/config"
	aaws "github.com/brunopadz/amictl/pkg/providers/aws"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	awsCmd.AddCommand(listAll)
}

var listAll = &cobra.Command{
	Use:     "list-all",
	Short:   "List all AMIs",
	Long:    `List all AMIs for multiple regions.`,
	Example: `  amictl aws list-all`,
	RunE:    runAll,
}

func runAll(cmd *cobra.Command, _ []string) error {

	var c cfg.Config

	err := viper.Unmarshal(&c)
	if err != nil {
		fmt.Println(err)
	}

	r := viper.GetStringSlice("aws.regions")

	td := pterm.TableData{
		{"AMI ID", "REGION"},
	}

	for _, v := range r {
		s, err := aaws.New(v)
		if err != nil {
			log.Fatalln("Couldn't create a session to AWS.")
		}

		client := ec2.NewFromConfig(s)

		input := &ec2.DescribeImagesInput{
			Owners: []string{
				viper.GetString("aws.account"),
			},
		}

		output, err := client.DescribeImages(context.TODO(), input)
		if err != nil {
			fmt.Println(err)
		}

		for _, i := range output.Images {
			td = append(td, []string{aws.ToString(i.ImageId), v})
		}

	}

	err = pterm.DefaultTable.WithHasHeader().WithData(td).Render()
	if err != nil {
		log.Fatalln("Couldn't render the results.")
	}

	return nil
}

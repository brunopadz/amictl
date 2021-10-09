package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	cfg "github.com/brunopadz/amictl/config"
	"github.com/brunopadz/amictl/pkg/commons"
	aaws "github.com/brunopadz/amictl/pkg/providers/aws"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	awsCmd.AddCommand(listUnused)
}

var listUnused = &cobra.Command{
	Use:     "list-unused",
	Short:   "List not used AMIs",
	Long:    `List all unused AMIs for multiple regions.`,
	Example: `  amictl aws list-unused`,
	RunE:    runUnused,
}

func runUnused(cmd *cobra.Command, _ []string) error {

	var c cfg.Config

	err := viper.Unmarshal(&c)
	if err != nil {
		fmt.Println(err)
	}

	r := viper.GetStringSlice("aws.regions")

	d := pterm.TableData{
		{"AMI ID", "REGION"},
	}

	var x []string

	for _, v := range r {

		var AllImages []string
		var UsedImages []string

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
			AllImages = append(AllImages, *i.ImageId)
		}

		var criteria = &ec2.DescribeInstancesInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("image-id"),
					Values: AllImages,
				},
			},
		}

		o, err := client.DescribeInstances(context.TODO(), criteria)
		if err != nil {
			log.Fatalln("Couldn't load instances data.")
		}

		for _, a := range o.Reservations {
			for _, i := range a.Instances {
				UsedImages = append(UsedImages, aws.ToString(i.ImageId))
			}
		}

		x = commons.Compare(AllImages, UsedImages)

		for _, id := range x {
			d = append(d, []string{id, v})
		}

	}

	l := len(d) - 1

	err = pterm.DefaultTable.WithHasHeader().WithData(d).Render()

	if (l > 20) && (l < 50) {
		pterm.Println(pterm.Yellow(fmt.Sprintf("%d AMIs are not being utilized.", l)))
	} else if ( l >= 50) { 
		pterm.Println(pterm.Red(fmt.Sprintf("%d AMIs are not being utilized.", l)))
	} else if (l <= 20){
		pterm.Println(pterm.Green(fmt.Sprintf("%d AMIs are not being utilized.", l)))
	}

	return nil
}

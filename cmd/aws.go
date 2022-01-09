package cmd

import (
	"github.com/spf13/cobra"
)

// awsCmd represents the aws command
var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Manages AWS AMIs",
	Long:  "Manages AWS AMIs by account and region. It's possible to list used, not used and also inspect AMIs.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(cmd.Long)
	},
}

func init() {
	rootCmd.AddCommand(awsCmd)

	delete.Flags().StringVarP(&ami, "ami", "a", "", "AMI ID")

	inspect.Flags().StringVarP(&ami, "ami", "a", "", "AMI ID")
	inspect.Flags().StringVarP(&region, "region", "r", "", "Region where the AMI was created")

	delete.MarkFlagRequired("ami")

	inspect.MarkFlagRequired("ami")
	inspect.MarkFlagRequired("region")

}

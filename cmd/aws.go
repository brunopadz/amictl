package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	account string
	region  string
	cost    bool
)

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Manage AWS AMIs",
	Long:  `Manages AWS AMIs by account and region`,
	Example: `  amictl aws list-all --account 123456789012 --region us-east-1
  amictl aws list-unused --account 123456789012 --region us-east-1`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Long)
	},
	TraverseChildren: true,
}

func init() {
	rootCmd.AddCommand(awsCmd)

	awsCmd.PersistentFlags().StringVarP(&account, "account", "a", "", "AWS account ID")
	awsCmd.PersistentFlags().StringVarP(&region, "region", "r", "", "AWS region ID")
	awsCmd.PersistentFlags().BoolVarP(&cost, "cost", "c", false, "Estimated Cost")

	awsCmd.MarkFlagRequired("account")
	awsCmd.MarkFlagRequired("region")
}

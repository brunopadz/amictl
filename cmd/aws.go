package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	account string
	region  string
)

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Manage AWS AMIs",
	Long:  `Manages AWS AMIs by account and region`,
	Example: `  amictl aws list-all --account 123456789012 --region us-east-1
amictl aws list-unused --account 123456789012 --region us-east-1`,
  TraverseChildren: true,
	Run: AWSCommand(),
}

func init() {
	rootCmd.AddCommand(awsCmd)

	awsCmd.PersistentFlags().StringVarP(&account, "account", "a", "", "AWS account ID")
	awsCmd.PersistentFlags().StringVarP(&region, "region", "r", "", "AWS region ID")

	_ = awsCmd.MarkFlagRequired("account")
	_ = awsCmd.MarkFlagRequired("region")
}

// AWSCommand print the available features from amictl
func AWSCommand() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Long)
	}
}

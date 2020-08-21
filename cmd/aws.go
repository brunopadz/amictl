package cmd

import (
	"github.com/spf13/cobra"
)

var (
	cost    bool
)

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Manage AWS AMIs",
	Long:  `Manages AWS AMIs by account and region`,
	Example: `  amictl aws list-all --account 123456789012 --region us-east-1
  amictl aws list-unused --account 123456789012 --region us-east-1`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(cmd.Long)
	},
}

func init() {
	rootCmd.AddCommand(awsCmd)

	awsCmd.PersistentFlags().StringP("account", "a", "", "AWS account ID")
	awsCmd.PersistentFlags().StringP("region", "r", "", "AWS region ID")
	awsCmd.PersistentFlags().BoolVarP(&cost, "cost", "c", false, "AWS ami estimated cost")

	_ = awsCmd.MarkPersistentFlagRequired("account")
	_ = awsCmd.MarkPersistentFlagRequired("region")
}

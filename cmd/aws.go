package cmd

import (
	"github.com/spf13/cobra"
)

// awsCmd represents the aws command
var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Manage AWS AMIs",
	Long:  "Manages AWS AMIs by account and region",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(cmd.Long)
	},
}

func init() {
	rootCmd.AddCommand(awsCmd)
}

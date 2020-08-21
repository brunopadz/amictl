package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "amictl",
	Short: "amictl is a CLI to control your AMIs and Images.",
	Long: `amictl is a super simple CLI to control your AMIs and Images.
	
⚠️   AWS is the only cloud provider supported.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

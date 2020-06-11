package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "amictl",
	Short: "amictl is a CLI to control your AMIs and cloud images.",
	Long: `amictl is a super simple CLI to control your AMIs and Images.
	
⚠️   AWS is the only cloud provider supported.`,
}

// Execute this function run ListAllCommand
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

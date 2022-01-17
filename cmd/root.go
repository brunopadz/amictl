package cmd

import (
	"fmt"
	"os"

	cfg "github.com/brunopadz/amictl/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "amictl",
	Short: "amictl is a tool to manage AMIs and Cloud Images",
	Long: `amictl is a super simple tool to manage AMIs and Cloud Images.

You can estimate how much money you are spending with unused AMIs, list how many
AMIs are being used and unused per region.

AWS is the only cloud provider supported at this moment.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.amictl.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {

	var c cfg.Config

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".amictl")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	if err := viper.Unmarshal(&c); err != nil {
		fmt.Fprintln(os.Stderr, "Unable to decode config file, %v", err)
	}
}

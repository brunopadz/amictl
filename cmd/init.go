package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/brunopadz/amictl/config"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize amictl",
	Long:  "Creates a config file to configure amitcl",
	RunE:  runInit,
}

func runInit(cmd *cobra.Command, _ []string) error {

	fmt.Println("")
	a := bufio.NewScanner(os.Stdin).Text()

	fmt.Println()
	r := bufio.NewScanner(os.Stdin).Bytes()

	d := config.Config{
		config.AwsConfig{
			Account: a,
			Regions: r,
		},
	}

	y, err := yaml.Marshal()

	return nil
}

func getInput() string {

	i := bufio.NewScanner(os.Stdin)
	i.Text()
}

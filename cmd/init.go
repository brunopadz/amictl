package cmd

import (
	"bufio"
	"fmt"
	"github.com/brunopadz/amictl/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strings"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize amictl.",
	Long:  "Creates a config file to configure amitcl.",
	RunE:  runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, _ []string) error {

	rs := []string{}

	h, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Couldn't determine home directory.")
	}
	p := h + "/.amictl.yaml"

	fmt.Println("Enter your AWS account ID: ")
	i := bufio.NewReader(os.Stdin)

	a, err := i.ReadString('\n')
	if err != nil {
		fmt.Println("Couldn't read AWS account ID.")
		os.Exit(1)
	}

	a = strings.TrimSuffix(a, "\n")

	r := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter the region: ")
		r.Scan()
		t := r.Text()
		if len(t) != 0 {
			rs = append(rs, t)
			fmt.Println("Enter one more region (y/n)?")

			c, err := i.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			}

			c = strings.ToLower(strings.TrimSpace(c))

			if c == "y" || c == "yes" {
				continue
			} else if c == "n" || c == "no" {
				break
			}

		} else {
			break
		}
	}

	if r.Err() != nil {
		fmt.Println("Error: ", r.Err())
	}

	c := config.Config{
		Aws: config.AwsConfig{
			Account: a,
			Regions: rs,
		},
	}

	d, err := yaml.Marshal(c)
	if err != nil {
		fmt.Println("Sei lá")
	}

	err = ioutil.WriteFile(p, d, 0600)

	return nil
}

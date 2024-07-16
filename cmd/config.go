package cmd

import (
	"fmt"
	"github.com/snail2sky/kofm/config"
	"os"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print current configuration",
	Long:  `Print current configuration, if config file not exist, print the program default config`,
	Run: func(cmd *cobra.Command, args []string) {
		configFile, _ := cmd.Flags().GetString("config")
		data, err := os.ReadFile(configFile)
		if err != nil {
			fmt.Print(string(config.DefaultConfig))
		}
		fmt.Print(string(data))
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")
	configCmd.Flags().String("config", "./config.yaml", "The os config file")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

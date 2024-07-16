/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/snail2sky/kofm/app"

	"github.com/spf13/cobra"
)

// fileServerCmd represents the fileServer command
var fileServerCmd = &cobra.Command{
	Use:   "file-server",
	Short: "Supply HTTP file server, to make a repo",
	Long:  `Supply HTTP file server, to make a repo`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		root, _ := cmd.Flags().GetString("root")

		app.RunFileServer(host, port, root)
	},
}

func init() {
	rootCmd.AddCommand(fileServerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fileServerCmd.PersistentFlags().String("foo", "", "A help for foo")
	fileServerCmd.Flags().String("host", "0.0.0.0", "Listen host")
	fileServerCmd.Flags().String("port", "8080", "Listen port")
	fileServerCmd.Flags().String("root", ".", "Web root dir")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fileServerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

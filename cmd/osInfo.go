package cmd

import (
	"fmt"
	"github.com/snail2sky/kofm/app"
	"github.com/spf13/cobra"
	"runtime"
)

// osInfoCmd represents the osInfo command
var osInfoCmd = &cobra.Command{
	Use:   "os-info",
	Short: "Print OS Information",
	Long:  `Print OS Information`,
	Run: func(cmd *cobra.Command, args []string) {
		osInfo := app.LoadOSInfo()
		fmt.Println("hostname:", osInfo["HOSTNAME"])
		fmt.Println("arch:", runtime.GOARCH)
		fmt.Println("type:", runtime.GOOS)
		fmt.Println("OS ID:", osInfo["ID"])
		fmt.Println("OS version:", osInfo["VERSION_ID"])
		fmt.Println("OS like:", osInfo["ID_LIKE"])
		fmt.Println("OS platform id:", osInfo["PLATFORM_ID"])
	},
}

func init() {
	rootCmd.AddCommand(osInfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// osInfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// osInfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

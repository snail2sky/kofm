package cmd

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
	"github.com/spf13/cobra"
)

// osInfoCmd represents the osInfo command
var osInfoCmd = &cobra.Command{
	Use:   "os-info",
	Short: "Print OS Information",
	Long:  `Print OS Information`,
	Run: func(cmd *cobra.Command, args []string) {
		info, _ := host.Info()
		fmt.Println("hostname:", info.Hostname)
		fmt.Println("arch:", info.KernelArch)
		fmt.Println("type:", info.OS)
		fmt.Println("OS ID:", info.Platform)
		fmt.Println("OS version:", info.PlatformVersion)
		fmt.Println("OS platform family:", info.PlatformFamily)
		fmt.Println("OS kernel version:", info.KernelVersion)
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


package cmd

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
	"github.com/spf13/cobra"
)

// osInfoCmd represents the osInfo command
var osInfoCmd = &cobra.Command{
	Use:   "os-info",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		info, _ := host.Info()
		//fmt.Println(info.OS)              // 操作系统类型
		//fmt.Println(info.Platform)        // 操作系统平台
		//fmt.Println(info.PlatformFamily)  // 操作系统平台家族
		//fmt.Println(info.PlatformVersion) // 操作系统平台版本
		//fmt.Println(info.HostID)
		//fmt.Println(info.KernelArch)
		fmt.Printf("%#v\n", *info)
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

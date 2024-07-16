package cmd

import (
	"github.com/snail2sky/kofm/app"
	"github.com/snail2sky/kofm/config"
	"github.com/spf13/cobra"
	"log"
	"path"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "mk-iso",
	Short: "Make target os iso",
	Long: `Initialize the environment and create an iso image file. 
The image file contains the software packages required for subsequent offline installation of k8s.`,
	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		kkZone, _ := cmd.Parent().PersistentFlags().GetString("kk-zone")
		configFile, _ := cmd.Flags().GetString("config")
		osList := config.LoadConfig(configFile)
		if len(osList) == 0 {
			log.Fatal("No Kofm config found")
		}
		initializer := app.NewInitializer(args[0], osList)
		initializer.SetInitializer(initializer)
		pkgsDir := path.Join(initializer.WorkerDir, initializer.GetPkgDir())

		initializer.Mkdir(initializer.WorkerDir, pkgsDir)
		initializer.Install("yum-utils", "createrepo", "genisoimage")
		initializer.DownloadPkg(
			initializer.WorkerDir)
		_ = initializer.MakeRepo(pkgsDir)
		initializer.GetKK(kkZone)
		initializer.MkISO()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")
	initCmd.Flags().String("config", "./config.yaml", "The os config file path")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

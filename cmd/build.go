package cmd

import (
	"github.com/snail2sky/kofm/app"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a offline kubesphere package.",
	Long:  `Build a offline kubesphere package.`,
	Run: func(cmd *cobra.Command, args []string) {
		manifest, _ := cmd.Flags().GetString("manifest")
		output, _ := cmd.Flags().GetString("output")
		kkZone, _ := cmd.Parent().PersistentFlags().GetString("kk-zone")
		kk, _ := cmd.Parent().PersistentFlags().GetString("kk-path")
		builder := app.NewBuilder(manifest, output, kkZone, kk)
		builder.Build()
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")
	buildCmd.Flags().StringP("manifest", "m", "./manifest-sample.yaml", "The manifest file")
	buildCmd.Flags().StringP("output", "o", "./kubesphere.tar.gz", "The kubesphere file")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

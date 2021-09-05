package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// 配置文件路径
var cfgPath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Short: "Parse Log Write And Transform To Other",
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVarP(&cfgPath, "config", "f", "./conf/config.yaml", "config file path")
}

package cmd

import (
	"fmt"
	"os"

	"github.com/0x0BSoD/glci-linter/pkg/gitlab"
	"github.com/0x0BSoD/glci-linter/pkg/helpers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	home := os.Getenv("HOME")
	configFile = fmt.Sprintf("%s/.config/glci-linter/config.yaml", home)
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)

	viper.AutomaticEnv()
	viper.SetEnvPrefix("GLCI")
	helpers.HandleError(viper.BindEnv("ACCESS_TOKEN"))
	helpers.HandleError(viper.BindEnv("LINTER_PATH"))
	helpers.HandleError(viper.BindEnv("SHOW_MERGED"))

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "glci-linter",
	Short: "Generate and lint GitLab-CI yaml with GitLab API",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		client := gitlab.NewClient(".")
		err := client.Lint()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"
	"os"

	"github.com/0x0BSoD/glci-linter/pkg/gitlab"
	"github.com/spf13/cobra"
)

var (
	personalAccessToken string
	repoPath            string
	showMerged          bool
)

func init() {
	rootCmd.PersistentFlags().StringVar(&personalAccessToken, "token", "", "GitLab access token with api access")
	rootCmd.PersistentFlags().StringVar(&repoPath, "path", ".", "Path to Gitr repo, default .")
	rootCmd.PersistentFlags().BoolVar(&showMerged, "merged", false, "Show merged YAML")
}

var rootCmd = &cobra.Command{
	Use:   "glci-linter",
	Short: "Generate and lint GitLab-CI yaml with GitLab API",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		client := gitlab.NewGitLabClient(personalAccessToken, repoPath, showMerged)
		err := client.Lint()
		if err != nil {
			panic(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

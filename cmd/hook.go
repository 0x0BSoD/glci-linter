package cmd

import (
	"fmt"

	"github.com/0x0BSoD/glci-linter/pkg/git"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "With 'hook' subcommand you can add or delete pre-commit hook in repo",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		repo := git.GitRepo{
			Path: ".",
		}
		exist, fullHookPath := repo.CheckHook()
		fmt.Println("Exist: ", exist)
		fmt.Println("Path: ", fullHookPath)
	},
}

var hookAddCmd = &cobra.Command{
	Use:   "add",
	Short: "'hook add' command create symlink from current binary as hook",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		repo := git.GitRepo{
			Path: ".",
		}
		repo.AddHook()
	},
}

var hookDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "'hook delete' command remove symlink",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		repo := git.GitRepo{
			Path: ".",
		}
		repo.DeleteHook()
	},
}

func init() {
	rootCmd.AddCommand(hookCmd)

	hookCmd.AddCommand(hookAddCmd)
	hookCmd.AddCommand(hookDeleteCmd)
}

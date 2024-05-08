package cmd

import (
	"fmt"

	"github.com/0x0BSoD/glci-linter/pkg/gitlab"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "The 'lint' subcommand will try to find and lint gitlab-ci.yaml file in direcory settetd as first arg.",
	Long: `The 'lint' subcommand will try to find and lint gitlab-ci.yaml file in direcory settetd as first arg. For example:

'<cmd> lint 'PATH TO directory with gitlab-ci.yaml'.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := gitlab.NewClient(".")
		err := client.Lint()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(lintCmd)
}

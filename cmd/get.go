package cmd

import (
	"fmt"
	"os"

	"github.com/btoll/github-release/release"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get release details for a repository and tag name",
	Run: func(cmd *cobra.Command, args []string) {
		repository := release.NewRelease(owner, repo, secret, false)
		release, err := repository.GetReleasesByTag(tag)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		repository.PrintRelease(*release)
	},
}

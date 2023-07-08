package cmd

import (
	"fmt"
	"os"

	"github.com/btoll/github-release/release"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all releases for a repository",
	//	Args: func(cmd *cobra.Command, args []string) error {
	//		if len(args) < 1 {
	//			return errors.New("list requires at least one argument")
	//		}
	//		return nil
	//	},
	Run: func(cmd *cobra.Command, args []string) {
		repository := release.NewRelease(owner, repo, secret, false)
		err := repository.ListReleases()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

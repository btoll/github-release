package cmd

import (
	"fmt"
	"os"

	"github.com/btoll/github-release/release"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a release, including its assets",
	Run: func(cmd *cobra.Command, args []string) {
		repository := release.NewRelease(owner, repo, secret, false)
		err := repository.DeleteRelease(tag)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("[SUCCESS] Deleted release %s and its assets in repository `%s`.\n", tag, repo)
	},
}

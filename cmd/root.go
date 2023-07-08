package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "release",
	Short: "github-release is a CLI task manager",
}

//var repository = &release.Release{
//	Owner:   "btoll",
//	Repo:    "",
//	Secret:  "",
//	Private: false,
//}

var (
	assetDir          = ""
	defaultSecretFile = "github-release.secret"
	owner             = "btoll"
	repo              = "repo"
	secret            = ""
	tag               = "tag"
)

func getHomeDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir, nil
}

func init() {
	RootCmd.AddCommand(createCmd)
	RootCmd.AddCommand(deleteCmd)
	RootCmd.AddCommand(getCmd)
	RootCmd.AddCommand(listCmd)

	createCmd.Flags().Bool("draft", false, "Create release as a draft")
	createCmd.Flags().Bool("prerelease", false, "Create release as a prerelease")
	createCmd.Flags().StringVarP(&assetDir, "dir", "d", "", "The location of the directory that contains the release assets")

	RootCmd.PersistentFlags().StringVarP(&repo, "repo", "r", "", "The repository in which to create the release (required)")
	RootCmd.MarkFlagRequired("repo")

	RootCmd.PersistentFlags().StringVarP(&tag, "tag", "t", "", "The name of the tag to create the release (required)")
	RootCmd.MarkFlagRequired("tag")

	homeDir, err := getHomeDir()
	if err != nil {
		secret = ""
	} else {
		secret = fmt.Sprintf("%s/%s", homeDir, defaultSecretFile)
	}

	RootCmd.PersistentFlags().StringVarP(&secret, "secret", "s", "", "The name of the local file that contains the GitHub API token")
}

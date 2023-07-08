package cmd

import (
	"fmt"
	"os"

	"github.com/btoll/github-release/release"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new release",
	Run: func(cmd *cobra.Command, args []string) {
		// This returns the UploadURL, but we're not using it, instead opting
		// to get it from the call to `GetLatestRelease` below.
		repository := release.NewRelease(owner, repo, secret, false)
		_, err := repository.CreateRelease(&tag)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		releaseID, err := repository.GetLatestRelease()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		entries, err := os.ReadDir(assetDir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("[INFO] Uploading release assets in repository", repo)
		for _, entry := range entries {
			// There **should* never be a directory in the package dir.
			if !entry.IsDir() {
				// For now, I'm not checking if there's an error because it will be
				// a PathError, and all the files in `entries` **should** exist.
				file, _ := os.Open(fmt.Sprintf("%s/%s", assetDir, entry.Name()))
				asset, err := repository.UploadReleaseAsset(releaseID, file)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("\tUploaded", *asset.Name)
			}
		}
	},
}

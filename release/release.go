package release

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
)

type Release struct {
	Owner   string
	Repo    string
	Secret  string
	Private bool
	True    bool
	False   bool
}

func NewRelease(owner, repo, secret string, isPrivate bool) Release {
	return Release{
		Owner:   owner,
		Repo:    repo,
		Secret:  secret,
		Private: isPrivate,
	}
}

var client *github.Client
var singleClient *bool

func getClient() *github.Client {
	if singleClient == nil {
		client = github.NewClient(getOAuthClient())
		b := true
		singleClient = &b
	}
	return client
}

var ctx context.Context
var singleContext *bool

func getContext() context.Context {
	if singleContext == nil {
		ctx = context.Background()
		b := true
		singleContext = &b
	}
	return ctx
}

func getOAuthClient() *http.Client {
	return oauth2.NewClient(getContext(), oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: func() string {
				token, err := ioutil.ReadFile("/home/btoll/github-release.secret")
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				// Remove the trailing newline.
				return string(token[0 : len(token)-1])
			}(),
		},
	))
}

// func (r *Release) CreateRelease(tagName *string, draft, preRelease *bool, assets *string) (string, error) {
func (r *Release) CreateRelease(tagName *string) (string, error) {
	body := fmt.Sprintf("%s release", *tagName)
	release, _, err := getClient().Repositories.CreateRelease(getContext(), r.Owner, r.Repo, &github.RepositoryRelease{
		TagName: tagName,
		//		Name:       &name,
		Body:  &body,
		Draft: &r.True,
		//		Prerelease: preRelease,
	})
	if err != nil {
		return "", err
	}

	fmt.Printf("[SUCCESS] Release %s created successfully in repository `%s`.\n", *tagName, r.Repo)
	// TODO return the whole release
	return *release.UploadURL, nil
}

func (r *Release) DeleteRelease(tagName string) error {
	release, err := r.GetReleasesByTag(tagName)
	if err != nil {
		return err
	}
	_, err = getClient().Repositories.DeleteRelease(getContext(), r.Owner, r.Repo, *release.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Release) GetRelease(releaseID int64) (*github.RepositoryRelease, error) {
	release, _, err := getClient().Repositories.GetRelease(getContext(), r.Owner, r.Repo, releaseID)
	if err != nil {
		return nil, err
	}
	return release, nil
}

func (r *Release) GetReleasesByTag(tagName string) (*github.RepositoryRelease, error) {
	release, _, err := getClient().Repositories.GetReleaseByTag(getContext(), r.Owner, r.Repo, tagName)
	if err != nil {
		return nil, err
	}
	return release, nil
}

// TODO return list of releases
func (r *Release) ListReleases() error {
	releases, _, err := getClient().Repositories.ListReleases(getContext(), r.Owner, r.Repo, &github.ListOptions{})
	if err != nil {
		return err
	}

	if len(releases) == 0 {
		fmt.Printf("[INFO] The repository `%s` has no releases.\n", r.Repo)
	} else {
		fmt.Printf("[INFO] The repository `%s` has the following releases:\n", r.Repo)
		for _, release := range releases {
			r.PrintRelease(*release)
		}
	}

	return nil
}

// TODO return entire release
func (r *Release) GetLatestRelease() (int64, error) {
	release, _, err := getClient().Repositories.GetLatestRelease(getContext(), r.Owner, r.Repo)
	if err != nil {
		return 0, err
	}
	return *release.ID, nil
}

func (r *Release) PrintRelease(release github.RepositoryRelease) {
	fmt.Printf("[%s]\n", *release.TagName)
	for _, asset := range release.Assets {
		fmt.Printf("\t%s\n", *asset.Name)
	}
	fmt.Printf("\t%s\n", *release.ZipballURL)
	fmt.Printf("\t%s\n", *release.TarballURL)
}

func (r *Release) UploadReleaseAsset(release_id int64, file *os.File) (*github.ReleaseAsset, error) {
	asset, _, err := getClient().Repositories.UploadReleaseAsset(getContext(), r.Owner, r.Repo, release_id, &github.UploadOptions{
		Name: filepath.Base(file.Name()),
	}, file)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

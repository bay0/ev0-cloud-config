package gist

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/google/go-github/github"
	"github.com/spf13/viper"
)

func Create(client *github.Client, configFilepath string, filename string) (*github.Gist, error) {
	ctx := context.Background()
	f := make(map[github.GistFilename]github.GistFile)
	bytes, err := ioutil.ReadFile(configFilepath)
	if err != nil {
		os.Exit(1)
	}
	f[github.GistFilename(filename)] = github.GistFile{Content: github.String(string(bytes))}
	gist := &github.Gist{
		Description: github.String(viper.GetString("defaultDescription")),
		Public:      github.Bool(viper.GetBool("public")),
		Files:       f,
	}
	gistResponse, _, err := client.Gists.Create(ctx, gist)
	return gistResponse, err
}

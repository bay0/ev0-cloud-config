package gist

import (
	"context"
	"io/ioutil"
	"os"
	"time"

	"github.com/google/go-github/github"
	"github.com/spf13/viper"
)

// Create pushes file to github gist
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

func CreateBackup(client *github.Client, csgopath string, configs []string) (*github.Gist, error) {
	currentTime := time.Now()
	ctx := context.Background()
	f := make(map[github.GistFilename]github.GistFile)
	for _, config := range configs {
		bytes, err := ioutil.ReadFile(csgopath + `\ev0lve\` + config)
		if err != nil {
			os.Exit(1)
		}
		f[github.GistFilename(config)] = github.GistFile{Content: github.String(string(bytes))}
	}
	gist := &github.Gist{
		Description: github.String("Backup: " + currentTime.Format("2006-01-02 15:04:05 Monday")),
		Public:      github.Bool(viper.GetBool("public")),
		Files:       f,
	}
	gistResponse, _, err := client.Gists.Create(ctx, gist)
	return gistResponse, err
}

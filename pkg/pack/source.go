package pack

import (
	"fmt"
	"regexp"

	"github.com/go-git/go-git/v5"
	"github.com/joomcode/errorx"
	"github.com/rs/zerolog/log"
)

var (
	repoNameMatcherSSH   = regexp.MustCompile(`git@.*:([\w-]+)\/([\w._-]+)\.git`)
	repoNameMatcherHTTPS = regexp.MustCompile(`https:\/\/.*\/([\w-]+)\/([\w._-]+)\.git`)
	errNoRemote          = errorx.NewType(errNS, "no_remote")
)

type names struct {
	user    string
	project string
}

func (n names) full() string {
	return fmt.Sprintf("%s/%s", n.user, n.project)
}

func (n names) isEmpty() bool {
	return n.user == "" || n.project == ""
}

var emptyNames = names{"", ""}

func namesFromRemote(remote string) names {
	m := repoNameMatcherSSH.FindStringSubmatch(remote)
	if len(m) > 2 {
		return names{m[1], m[2]}
	}
	m = repoNameMatcherHTTPS.FindStringSubmatch(remote)
	if len(m) > 2 {
		return names{m[1], m[2]}
	}
	return emptyNames
}

func getPackageRemote(packagePath string) (string, error) {
	r, err := git.PlainOpen(packagePath)
	if err != nil {
		return "", errorx.Decorate(err, "failed to open local package git")
	}
	cfg, err := r.Config()
	if err != nil {
		return "", errorx.Decorate(err, "failed to read local git repo config")
	}

	if len(cfg.Remotes) == 0 {
		return "", errNoRemote.New("no git remote at %s", packagePath)
	}

	firstRemoteName := ""
	for name, remote := range cfg.Remotes {
		if firstRemoteName == "" {
			firstRemoteName = name
		}
		if name == "origin" {
			return remote.URLs[0], nil
		}
	}

	log.Info().Msgf("no origin remote found for package at [%s]; using %s", packagePath, firstRemoteName)
	return cfg.Remotes[firstRemoteName].URLs[0], nil
}

func clone(url, location string) error {
	_, err := git.PlainClone(location, false, &git.CloneOptions{
		URL: url,
	})
	return err
}

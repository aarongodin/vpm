package pack

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/go-git/go-git/v5"
	"github.com/rs/zerolog/log"
)

var (
  repoNameMatcherSSH = regexp.MustCompile(`:(\w+\/[\w._-]+)\.git`)
  repoNameMatcherHTTPS = regexp.MustCompile(`https://.*/(\w+\/[\w._-]+)`)
)

var (
  ErrPackage = errors.New("package error")
  ErrPackageNoRepo = fmt.Errorf("%w: no git repository found", ErrPackage)
  ErrPackageNoRemote = fmt.Errorf("%w: no git remotes found", ErrPackage)
)

func newErr(err error) error {
  return fmt.Errorf("%w: %w", ErrPackage, err)
}

func RepoNameFromRemote(remote string) string {
  matchesSSH := repoNameMatcherSSH.FindStringSubmatch(remote)
  if len(matchesSSH) > 1 {
    return matchesSSH[1]
  }
  matchesHTTPS := repoNameMatcherHTTPS.FindStringSubmatch(remote)
  if len(matchesHTTPS) > 1 {
    return matchesHTTPS[1]
  }
  return ""
}

func GetPackageRemote(packagePath string) (string, error) {
  r, err := git.PlainOpen(packagePath)
  if err != nil {
    return "", ErrPackageNoRepo
  }
  cfg, err := r.Config()
  if err != nil {
		return "", newErr(err)
  }

  if len(cfg.Remotes) == 0 {
    return "", ErrPackageNoRemote
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

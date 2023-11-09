package pack

import (
	"fmt"
	"os"
	"path"

	"github.com/rs/zerolog/log"
)

const (
  VIM_LOAD_START = "start"
  VIM_LOAD_OPT = "opt"
)

type Pack struct {
  Name string
  Dirname string
  Location string
  RemoteURL string
  Group string
  Load string
}

func ListGroups(packDir string) ([]string, error) {
  entries, err := os.ReadDir(packDir)
  if err != nil {
    return nil, fmt.Errorf("error reading ~/.vim/pack: %w", err)
  }
  groups := make([]string, 0, len(entries))
  for _, e := range entries {
    groups = append(groups, e.Name())
  }
  return groups, nil
}

func ListPacksForGroup(packDir string, group string) ([]Pack, error) {
  packs := make([]Pack, 0)
  loadTypes, err := os.ReadDir(path.Join(packDir, group))
  if err != nil {
    return nil, err
  }

  for _, loadType := range loadTypes {
    switch loadType.Name() {
      case VIM_LOAD_START, VIM_LOAD_OPT:
        entries, err := os.ReadDir(path.Join(packDir, group, loadType.Name()))
        if err != nil {
          return nil, fmt.Errorf("err reading loadType dir: %w", err)
        }
        for _, e := range entries {
          location := path.Join(packDir, group, loadType.Name(), e.Name())
          remote, _ := GetPackageRemote(location)
          packs = append(packs, Pack{
            Name: RepoNameFromRemote(remote),
            RemoteURL: remote,
            Dirname: e.Name(),
            Location: location,
            Group: group,
            Load: loadType.Name(),
          })
        }
      default:
        log.Warn().Msgf("unknown load directory in %s: %s", path.Join(packDir, group), loadType.Name())
    }
  }

  return packs, nil
}

func ListPacks(packDir string) ([]Pack, error) {
  packs := make([]Pack, 0)

  groups, err := ListGroups(packDir)
  if err != nil {
    return nil, err
  }

  for _, group := range groups {
    groupPacks, err := ListPacksForGroup(packDir, group)
    if err != nil {
      return nil, err
    }
    packs = append(packs, groupPacks...)
  }

  return packs, nil
}

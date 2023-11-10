package pack

import (
  "errors"
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
          remote, _ := getPackageRemote(location)
          names := namesFromRemote(remote)
          packs = append(packs, Pack{
            Name: names.full(),
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

// Location finds the file path for a given pacakge name, if it is already added.
func GetByName(packDir string, name string) (Pack, error) {
  packs, err := ListPacks(packDir)
  if err != nil {
    return Pack{}, err
  }
  for _, p := range packs {
    if p.Name == name {
      return p, nil
    }
  }
  return Pack{}, ErrPackageNotAdded
}

func AddPack(packDir, url, group, load string) (Pack, error) {
  names := namesFromRemote(url)
  if names.isEmpty() {
    return Pack{}, ErrPackageInvalidRemote
  }

  existing, err := GetByName(packDir, names.full())
  if err != nil {
    if errors.Is(err, ErrPackageNotAdded) {
      pack, err := install(packDir, url, group, load, names)
      if err != nil {
        return Pack{}, err
      }
      return pack, nil
    } else {
      return Pack{}, err
    }
  }

  log.Info().Any("pack", existing).Msgf("package already installed")
  return existing, nil
}

func RemovePack(packDir, name string) error {
  existing, err := GetByName(packDir, name)
  if err != nil {
    return err
  }

  if err := os.RemoveAll(existing.Location); err != nil {
    return ErrPackageFileOperation
  }

  return nil
}

func install(packDir, url, group, load string, n names) (Pack, error) {
  location := path.Join(packDir, group, load, n.project)
  if err := clone(url, location); err != nil {
    log.Err(err).
      Str("url", url).
      Str("location", location).
      Msg("error cloning repository")
    return Pack{}, ErrPackageClone
  }
  return Pack{
    Name: n.full(),
    RemoteURL: url,
    Dirname: n.project,
    Location: location,
    Group: group,
    Load: load,
  }, nil
}

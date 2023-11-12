package pack

import (
	"os"
	"path"

	"github.com/joomcode/errorx"
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

var (
  namespace = errorx.NewNamespace("pack")
  packNotFound = errorx.NewType(namespace, "not_found", errorx.NotFound())
  packAlreadyExists = errorx.NewType(namespace, "already_exists", errorx.Duplicate())
)

func IsLoadType(load string) bool {
  return load == VIM_LOAD_START || load == VIM_LOAD_OPT
}

func ListGroups(packDir string) ([]string, error) {
  entries, err := os.ReadDir(packDir)
  if err != nil {
    return nil, errorx.Decorate(err, "failed reading %s", packDir)
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
          return nil, errorx.Decorate(err, "failed reading group %s load type %s", group, loadType.Name())
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
  return Pack{}, packNotFound.New("pack %s not found", name)
}

func AddPack(packDir, url, group, load string) (Pack, error) {
  if len(group) == 0 {
    return Pack{}, errorx.IllegalArgument.New("group must be present")
  }
  if len(load) == 0 || !IsLoadType(load) {
    return Pack{}, errorx.IllegalArgument.New("load must be a valid load type (either start or opt)")
  }

  names := namesFromRemote(url)
  if names.isEmpty() {
    return Pack{}, errorx.IllegalArgument.New("could not determine package name from provided remote URL")
  }

  existing, err := GetByName(packDir, names.full())
  if err != nil && !errorx.IsNotFound(err) { 
    return Pack{}, errorx.Decorate(err, "failed to find existing pack")
  }
  if existing.Name == names.full() {
    return existing, packAlreadyExists.New("pack %s already exists", existing.Name)
  }

  pack, err := install(packDir, url, group, load, names)
  if err != nil {
    return Pack{}, errorx.Decorate(err, "failed to install pack %s", url)
  }
  return pack, nil
}

func RemovePack(packDir, name string) error {
  existing, err := GetByName(packDir, name)
  if err != nil {
    return err
  }

  if err := os.RemoveAll(existing.Location); err != nil {
    return errorx.Decorate(err, "failed to remove pack directory")
  }

  return nil
}

func ChangePack(packDir, name, group, load string) (Pack, error) {
  pack, err := GetByName(packDir, name)
  if err != nil {
    return Pack{}, err
  }
  if pack.Group == group && pack.Load == load {
    log.Info().Msgf("notice: pack %s not changed", pack.Name)
    return pack, nil
  }
  newGroup := pack.Group
  if group != "" {
    newGroup = group
  }
  newLoad := pack.Load
  if load != "" {
    newLoad = load
  }
  newLocation := path.Join(packDir, newGroup, newLoad, pack.Dirname)
  if err := os.MkdirAll(path.Join(packDir, newGroup, newLoad), os.FileMode(int(0766))); err != nil {
    return Pack{}, errorx.Decorate(err, "failed to create group and load directory")
  }
  if err := os.Rename(pack.Location, newLocation); err != nil {
    return Pack{}, errorx.Decorate(err, "failed to change pack %s to group %s and load type %s", pack.Name, group, load)
  }
  pack.Location = newLocation
  pack.Group = newGroup 
  pack.Load = newLoad
  return pack, nil
}

func install(packDir, url, group, load string, n names) (Pack, error) {
  location := path.Join(packDir, group, load, n.project)
  if err := clone(url, location); err != nil {
    return Pack{}, errorx.Decorate(err, "failed to clone pack")
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

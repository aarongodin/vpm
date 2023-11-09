package pack

import (
	"os/user"
	"path"
)

func SetPackDir(packDir *string) error {
  if packDir != nil && *packDir != "" {
    return nil
  }
  current, err := user.Current()
  if err != nil {
    return err
  }
  *packDir = path.Join(current.HomeDir, ".vim/pack")
  return nil
}

package pack

import (
  "errors"
  "fmt"
)

var (
  ErrPackage = errors.New("package error")
  ErrPackageInvalidRemote = errors.New("%w: invalid remote")
  ErrPackageNoRepo = fmt.Errorf("%w: no git repository found", ErrPackage)
  ErrPackageNoRemote = fmt.Errorf("%w: no git remotes found", ErrPackage)
  ErrPackageNotAdded = fmt.Errorf("%w: not yet added", ErrPackage)
  ErrPackageClone = fmt.Errorf("%w: unable to clone", ErrPackage)
  ErrPackageFileOperation = fmt.Errorf("%w: fatal file operation", ErrPackage)
)

func newErr(err error) error {
  return fmt.Errorf("%w: %w", ErrPackage, err)
}


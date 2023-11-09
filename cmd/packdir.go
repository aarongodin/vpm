package cmd

import (
  "github.com/aarongodin/vpm/pkg/pack"
  "github.com/rs/zerolog/log"
)

// packDir is resolved at runtime (ex: `/home/user/.vim/pack`)
var packDir string

func initPackDir() {
  if err := pack.SetPackDir(&packDir); err != nil {
    log.Fatal().Err(err).Msg("fatal error finding user")
  }
}


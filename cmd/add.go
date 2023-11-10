package cmd

import (
  "fmt"
  "github.com/aarongodin/vpm/pkg/pack"
  "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a package",
	Run: func(cmd *cobra.Command, args []string) {
    url := "https://github.com/tpope/vim-commentary.git"
    group := "default"
    load := "opt"
    pack, err := pack.AddPack(packDir, url, group, load)
    if err != nil {
      log.Fatal().Err(err).Msg("fatal error adding package")
    }
    fmt.Printf("pack added: %s\n", pack.Name)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

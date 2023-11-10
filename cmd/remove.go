package cmd

import (
	"fmt"

	"github.com/aarongodin/vpm/pkg/format"
	"github.com/aarongodin/vpm/pkg/pack"
  "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a package",
	Run: func(cmd *cobra.Command, args []string) {
    name := "tpope/vim-commentary"
    if err := pack.RemovePack(packDir, name); err != nil {
      log.Err(err).Msg("error removing package")
    }
    fmt.Println(format.SuccessStyle.Render(
      fmt.Sprintf("package %s removed", name),
    ))
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

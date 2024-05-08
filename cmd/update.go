package cmd

import (
	"fmt"

	"github.com/aarongodin/vpm/pkg/format"
	"github.com/aarongodin/vpm/pkg/pack"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	var updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update packs",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for _, name := range args {
				newSha, err := pack.UpdatePack(packDir, name)
				if err != nil {
					log.Fatal().Err(err).Msg("error updating package")
				}
				log.Info().Str("name", name).Str("sha", newSha).Msg("package updated")
			}
			fmt.Println(format.InfoStyle.Render("done updating packages"))
		},
	}
	rootCmd.AddCommand(updateCmd)
}

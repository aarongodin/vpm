package cmd

import (
	"fmt"
	"os"

	"github.com/aarongodin/vpm/pkg/format"
	"github.com/aarongodin/vpm/pkg/pack"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func init() {
	var sourceCmd = &cobra.Command{
		Use:   "source",
		Short: "Source packs",
		Long:  "Source packs from a YAML file, typically derived from listing the packs on another machine",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			sourceFile, err := os.ReadFile(args[0])
			if err != nil {
				log.Fatal().Err(err).Msg("error reading source file")
			}
			var packs []pack.Pack
			if err := yaml.Unmarshal(sourceFile, &packs); err != nil {
				log.Fatal().Err(err).Msg("error unmarshaling source YAML")
			}
			for _, p := range packs {
				added, err := pack.AddPack(packDir, p.RemoteURL, p.Group, p.Load)
				if err != nil {
					log.Fatal().Err(err).Msg("error adding package")
				}
				log.Info().
					Str("name", added.Name).
					Str("group", added.Group).
					Str("load", added.Load).
					Msg("package added")
			}
			fmt.Println(format.InfoStyle.Render("packs sourced"))
		},
	}
	rootCmd.AddCommand(sourceCmd)
}

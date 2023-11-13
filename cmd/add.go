package cmd

import (
	"github.com/aarongodin/vpm/pkg/pack"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	var group, load string
	addCmd := &cobra.Command{
		Use:   "add <url>",
		Short: "Add a package",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			pack, err := pack.AddPack(packDir, args[0], group, load)
			if err != nil {
				log.Fatal().Err(err).Msg("error adding package")
			}
			log.Info().
				Str("name", pack.Name).
				Str("group", pack.Group).
				Str("load", pack.Load).
				Msg("package added")
		},
	}
	addCmd.Flags().StringVarP(&group, "group", "g", "default", "specify package group")
	addCmd.Flags().StringVarP(&load, "load", "l", "start", "specify package loading")
	rootCmd.AddCommand(addCmd)
}

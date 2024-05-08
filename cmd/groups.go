package cmd

import (
	"fmt"

	"github.com/aarongodin/vpm/pkg/format"
	"github.com/aarongodin/vpm/pkg/pack"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var groupsCmd = &cobra.Command{
	Use:     "groups",
	Aliases: []string{"group", "g"},
	Short:   "List groups",
	Long:    "List groups of packages managed by the subdirectories of the pack folder",
	Run: func(cmd *cobra.Command, args []string) {
		groups, err := pack.ListGroups(packDir)
		if err != nil {
			log.Fatal().Err(err).Msg("unexpected error")
		}
		if len(groups) == 0 {
			fmt.Println(format.InfoStyle.Render("no groups found"))
		}
		for _, g := range groups {
			fmt.Println(g)
		}
	},
}

func init() {
	rootCmd.AddCommand(groupsCmd)
}

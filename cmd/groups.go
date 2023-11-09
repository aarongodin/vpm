package cmd

import (
	"fmt"

	"github.com/aarongodin/vpm/pkg/pack"
  "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// groupsCmd represents the groups command
var groupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "List groups",
  Long: "List groups of packages managed by the subdirectories of the pack folder",
	Run: func(cmd *cobra.Command, args []string) {
    groups, err := pack.ListGroups(packDir)
    if err != nil {
      log.Fatal().Err(err).Msg("fatal error listing groups")
    }
    for _, g := range groups {
      fmt.Println(g)
    }
	},
}

func init() {
	rootCmd.AddCommand(groupsCmd)
}

package cmd

import (
	"fmt"

	"github.com/aarongodin/vpm/pkg/format"
	"github.com/aarongodin/vpm/pkg/pack"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	var removeCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove a package",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for _, name := range args {
				if err := pack.RemovePack(packDir, name); err != nil {
					log.Err(err).Msg("error removing package")
				}
				fmt.Println(format.InfoStyle.Render(
					fmt.Sprintf("package %s removed", name),
				))
			}
		},
	}
	rootCmd.AddCommand(removeCmd)
}

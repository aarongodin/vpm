package cmd

import (
	"fmt"

	"github.com/aarongodin/vpm/pkg/format"
	"github.com/aarongodin/vpm/pkg/pack"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// packsCmd represents the packs command
var packsCmd = &cobra.Command{
	Use:   "packs",
	Short: "List packages",
	Long:  "List packages managed through your .vim/pack directory",
	Run: func(cmd *cobra.Command, args []string) {
		packs, err := pack.ListPacks(packDir)
		if err != nil {
			log.Fatal().Err(err).Msg("unexpected error")
		}
		if len(packs) == 0 {
			fmt.Println(format.InfoStyle.Render("no packages found"))
			return
		}
		fmt.Println(format.ShowPackageList(packs))
	},
}

func init() {
	rootCmd.AddCommand(packsCmd)
}

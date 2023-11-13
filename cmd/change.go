package cmd

import (
	"fmt"

	"github.com/aarongodin/vpm/pkg/format"
	"github.com/aarongodin/vpm/pkg/pack"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	var group, load string
	var changeCmd = &cobra.Command{
		Use:   "change",
		Short: "Change properties of an added package",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for _, name := range args {
				if _, err := pack.ChangePack(packDir, name, group, load); err != nil {
					log.Fatal().Err(err).Msg("error changing package")
				}
			}
			fmt.Println(format.InfoStyle.Render("packs changed"))
		},
	}
	changeCmd.Flags().StringVarP(&group, "group", "g", "", "specify package group")
	changeCmd.Flags().StringVarP(&load, "load", "l", "", "specify package loading")
	rootCmd.AddCommand(changeCmd)
}

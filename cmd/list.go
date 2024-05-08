package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/aarongodin/vpm/pkg/format"
	"github.com/aarongodin/vpm/pkg/pack"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// listCmd represents the packs command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List packages",
	Long:    "List packages managed through your .vim/pack directory",
	Run: func(cmd *cobra.Command, args []string) {
		packs, err := pack.ListPacks(packDir)
		if err != nil {
			log.Fatal().Err(err).Msg("unexpected error")
		}

		switch outputFormat {
		case "yaml":
			output, err := yaml.Marshal(packs)
			if err != nil {
				log.Fatal().Err(err).Msg("error writing YAML")
			}
			fmt.Println(string(output))
		case "json":
			output, err := json.Marshal(packs)
			if err != nil {
				log.Fatal().Err(err).Msg("error writing JSON")
			}
			fmt.Println(string(output))
		default:
			if len(packs) == 0 {
				fmt.Println(format.InfoStyle.Render("no packages found"))
				return
			}
			fmt.Println(format.ShowPackageList(packs))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

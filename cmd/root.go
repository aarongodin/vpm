package cmd

import (
	"fmt"
	"os"

	"github.com/aarongodin/vpm/pkg/pack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var outputFormat string

// packDir is resolved at runtime (ex: `/home/user/.vim/pack`)
var packDir string

func initPackDir() {
	if err := pack.SetPackDir(&packDir); err != nil {
		log.Fatal().Err(err).Msg("fatal error finding user")
	}

	// TODO(aarongodin): what are the right permissions for this directory
	if err := os.MkdirAll(packDir, os.FileMode(int(0766))); err != nil {
		log.Fatal().Err(err).Msg("error creating vim pack directory")
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vpm",
	Short: "Vim Package Manager",
	Long:  "vpm: package management CLI using the default package manager in Vim",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	initPackDir()
	cobra.OnInitialize(initConfig)

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.vpm.yaml)")

	// // Cobra also supports local flags, which will only run
	// // when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "format", "f", "", `Output format (if specified, one of either "yaml" or "json")`)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".vpm" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".vpm")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

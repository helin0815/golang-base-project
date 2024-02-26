package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlabee.chehejia.com/gopkg/lsego/pkg/log"

	"gitlabee.chehejia.com/k8s/liks-gitops/internal/app"
)

var cfgFile string
var version = "devel"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "serve",
	Short:   "A brief description of your application",
	Version: version,
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Infof("Current version: %s", version)
		s, err := app.InitializeServer(cmd.Context())
		if err != nil {
			return err
		}

		return s.Run()
	},
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
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config.yml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

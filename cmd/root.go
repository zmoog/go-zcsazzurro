package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "zcs",
	Short: "A CLI for fetching data from the ZCS Azzurro API",
	Long:  "A CLI for fetching data from the ZCS Azzurro API. It allows you to fetch realtime data for a thing.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if !viper.IsSet("client_id") {
			return fmt.Errorf("client-id is required")
		}
		if !viper.IsSet("api_auth") {
			return fmt.Errorf("api-auth is required")
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		// fmt.Println("Ooops, something went wrong:", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	rootCmd.PersistentFlags().String("client-id", "", "The client ID to use for authentication")
	rootCmd.PersistentFlags().String("api-auth", "", "The API key to use for authentication")
	rootCmd.PersistentFlags().String("api-endpoint", "https://third.zcsazzurroportal.com:19003", "The API endpoint to use for authentication")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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

		// Search config in home directory with name ".zcs" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".zcs")
	}

	viper.SetEnvPrefix("ZCS")
	viper.AutomaticEnv() // read in environment variables that match

	viper.BindPFlag("client_id", rootCmd.PersistentFlags().Lookup("client-id"))
	viper.BindPFlag("api_auth", rootCmd.PersistentFlags().Lookup("api-auth"))
	viper.BindPFlag("api_endpoint", rootCmd.PersistentFlags().Lookup("api-endpoint"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

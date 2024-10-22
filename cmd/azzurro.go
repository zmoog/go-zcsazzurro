package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var azzurroCmd = &cobra.Command{
	Use:   "azzurro",
	Short: "A CLI for fetching data from the ZCS Azzurro API",
	Long:  "A CLI for fetching data from the ZCS Azzurro API. It allows you to fetch realtime data for a thing.",
	Run: func(cmd *cobra.Command, args []string) {
		clientID := viper.GetString("client-id")
		apiAuth := viper.GetString("api-auth")
		apiEndpoint := viper.GetString("api-endpoint")

		fmt.Printf("clientID: %s\n", clientID)
		fmt.Printf("apiAuth: %s\n", apiAuth)
		fmt.Printf("apiEndpoint: %s\n", apiEndpoint)
	},
}

func init() {
	rootCmd.AddCommand(azzurroCmd)
}

package cmd

import (
	"github.com/spf13/cobra"
)

var azzurroCmd = &cobra.Command{
	Use:   "azzurro",
	Short: "A CLI for fetching data from the ZCS Azzurro API",
	Long:  "A CLI for fetching data from the ZCS Azzurro API. It allows you to fetch realtime data for a thing.",
}

func init() {
	rootCmd.AddCommand(azzurroCmd)
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zmoog/zcs/azzurro"
)

var fetchRealtimeCmd = &cobra.Command{
	Use:   "fetch-realtime",
	Short: "Fetch realtime data for a thing",
	Long:  "Fetch realtime data for a thing with the given ID",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		thingID := viper.GetString("thing_id")
		if thingID == "" {
			return fmt.Errorf("thing-id is required")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		client := azzurro.NewClientWithBaseURL(
			viper.GetString("api_auth"),
			viper.GetString("client_id"),
			viper.GetString("api_endpoint"),
		)

		data, err := client.FetchRealtimeData(viper.GetString("thing_id"))
		if err != nil {
			return fmt.Errorf("failed to fetch realtime data: %v", err)
		}

		// Do something with the data
		if !data.RealtimeData.Success {
			return fmt.Errorf("failed to fetch realtime data: %t", data.RealtimeData.Success)
		}

		for _, v := range data.RealtimeData.Params.Value {
			for _, v := range v {
				fmt.Printf("Power importing: %.2f\n", v.PowerImporting)
				fmt.Printf("Power exporting: %.2f\n", v.PowerExporting)
				fmt.Printf("Power generating: %.2f\n", v.PowerGenerating)
				fmt.Printf("Power consuming: %.2f\n", v.PowerConsuming)
				fmt.Printf("Battery level: %d%%\n", v.BatterySoC)
				fmt.Printf("Last update: %s\n", v.LastUpdate)
			}
		}

		return nil
	},
}

func init() {
	fetchRealtimeCmd.Flags().String("thing-id", "", "The ID of the thing to fetch realtime data for")
	viper.BindPFlag("thing_id", fetchRealtimeCmd.Flags().Lookup("thing-id"))

	azzurroCmd.AddCommand(fetchRealtimeCmd)
}

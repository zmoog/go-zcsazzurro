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
		thingKey := viper.GetString("thing_key")
		if thingKey == "" {
			return fmt.Errorf("thing-key is required")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		client := azzurro.NewClientWithBaseURL(
			viper.GetString("api_auth"),
			viper.GetString("client_id"),
			viper.GetString("api_endpoint"),
		)

		data, err := client.FetchRealtimeData(viper.GetString("thing_key"))
		if err != nil {
			return fmt.Errorf("failed to fetch realtime data: %v", err)
		}

		// Do something with the data
		if !data.RealtimeData.Success {
			return fmt.Errorf("failed to fetch realtime data: %t", data.RealtimeData.Success)
		}

		for _, v := range data.RealtimeData.Params.Value {
			for _, v := range v {
				fmt.Printf("Power generating: %.2f\n", v.PowerGenerating)
				fmt.Printf("Power consuming: %.2f\n", v.PowerConsuming)
				fmt.Printf("Power importing: %.2f\n", v.PowerImporting)
				fmt.Printf("Power exporting: %.2f\n", v.PowerExporting)
				fmt.Printf("Power charging: %.2f\n", v.PowerCharging)
				fmt.Printf("Power discharging: %.2f\n", v.PowerDischarging)
				fmt.Printf("Battery level: %d%%\n", v.BatterySoC)
				fmt.Printf("Last update: %s\n", v.LastUpdate)
			}
		}

		return nil
	},
}

func init() {
	fetchRealtimeCmd.Flags().String("thing-key", "", "The key of the thing to fetch realtime data for")
	viper.BindPFlag("thing_key", fetchRealtimeCmd.Flags().Lookup("thing-key"))

	azzurroCmd.AddCommand(fetchRealtimeCmd)
}

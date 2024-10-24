package cmd

import (
	"fmt"

	"github.com/pterm/pterm"
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

		table := pterm.TableData{}

		// Add the header row
		table = append(table, []string{
			"Power generating",
			"Power consuming",
			"Power importing",
			"Power exporting",
			"Power charging",
			"Power discharging",
			"Battery level",
			"Last update",
		})

		// Add the data rows
		for _, v := range data.RealtimeData.Params.Value {
			for _, v := range v {
				table = append(table, []string{
					fmt.Sprintf("%.2f", v.PowerGenerating),
					fmt.Sprintf("%.2f", v.PowerConsuming),
					fmt.Sprintf("%.2f", v.PowerImporting),
					fmt.Sprintf("%.2f", v.PowerExporting),
					fmt.Sprintf("%.2f", v.PowerCharging),
					fmt.Sprintf("%.2f", v.PowerDischarging),
					fmt.Sprintf("%d%%", v.BatterySoC),
					v.LastUpdate.String(),
				})
			}
		}

		// Render the table
		pterm.DefaultTable.WithHasHeader().WithData(table).Render()

		return nil
	},
}

func init() {
	fetchRealtimeCmd.Flags().String("thing-key", "", "The key of the thing to fetch realtime data for")
	viper.BindPFlag("thing_key", fetchRealtimeCmd.Flags().Lookup("thing-key"))

	azzurroCmd.AddCommand(fetchRealtimeCmd)
}

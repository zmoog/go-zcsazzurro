package azzurro

import "time"

// ----------------------------------------------------------------------------
// Requests
// ----------------------------------------------------------------------------

type RealtimeDataRequest struct {
	RealtimeData RealtimeData `json:"realtimeData"`
}

type RealtimeData struct {
	Command string `json:"command"`
	Params  Params `json:"params"`
}

type Params struct {
	ThingKey       string `json:"thingKey"`
	RequiredValues string `json:"requiredValues"`
}

// ----------------------------------------------------------------------------
// Responses
// ----------------------------------------------------------------------------

type RealtimeDataResponse struct {
	RealtimeData struct {
		Params struct {
			Value []map[string]struct {
				EnergyDischargingTotal   float64   `json:"energyDischargingTotal"`
				PowerExporting           float64   `json:"powerExporting"`
				EnergyExportingTotal     float64   `json:"energyExportingTotal"`
				EnergyDischarging        float64   `json:"energyDischarging"`
				BatteryCycletime         int       `json:"batteryCycletime"`
				LastUpdate               time.Time `json:"lastUpdate"`
				EnergyGenerating         float64   `json:"energyGenerating"`
				EnergyAutoconsumingTotal float64   `json:"energyAutoconsumingTotal"`
				EnergyImporting          float64   `json:"energyImporting"`
				EnergyCharging           float64   `json:"energyCharging"`
				PowerImporting           float64   `json:"powerImporting"`
				EnergyChargingTotal      float64   `json:"energyChargingTotal"`
				EnergyConsumingTotal     float64   `json:"energyConsumingTotal"`
				EnergyAutoconsuming      float64   `json:"energyAutoconsuming"`
				PowerConsuming           float64   `json:"powerConsuming"`
				EnergyConsuming          float64   `json:"energyConsuming"`
				PowerGenerating          float64   `json:"powerGenerating"`
				EnergyImportingTotal     float64   `json:"energyImportingTotal"`
				EnergyExporting          float64   `json:"energyExporting"`
				BatterySoC               int       `json:"batterySoC"`
				ThingFind                string    `json:"thingFind"`
				PowerAutoconsuming       float64   `json:"powerAutoconsuming"`
				PowerCharging            float64   `json:"powerCharging"`
				EnergyGeneratingTotal    float64   `json:"energyGeneratingTotal"`
				PowerDischarging         float64   `json:"powerDischarging"`
			} `json:"value"`
		} `json:"params"`
		Success bool `json:"success"`
	} `json:"realtimeData"`
}

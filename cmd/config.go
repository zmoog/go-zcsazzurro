package cmd

type Config struct {
	// The base URL for the API
	BaseURL string `mapstructure:"base_url"`
	// The API key to use for authentication
	APIAuth  string `mapstructure:"api_auth"`
	ClientID string `mapstructure:"client_id"`
}

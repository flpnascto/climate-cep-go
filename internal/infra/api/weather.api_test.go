package api

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestFetchWeatherApi(t *testing.T) {
	if err := godotenv.Load("../../../.env"); err != nil {
		log.Fatalf("No .env file found")
	}
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		log.Fatalf("No WEATHER_API_KEY found")
	}

	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{
			name:    "Test FetchWeatherApi with success",
			value:   "Belo Horizonte",
			wantErr: false,
		},
		{
			name:    "Test FetchWeatherApi with success",
			value:   "São Paulo",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := FetchWeatherApi(tt.value, apiKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchCepApi() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}

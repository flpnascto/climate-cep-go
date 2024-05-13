package api

import (
	"testing"

	"github.com/flpnascto/climate-cep-go/internal/entity"
)

func TestFetchCepApi(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{
			name:    "Test FetchCepAPI with success",
			value:   "01001000",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cep, _ := entity.NewCep(tt.value)
			_, err := FetchCepApi(cep)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchCepApi() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}

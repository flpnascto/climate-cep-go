package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTempCelsius(t *testing.T) {
	tests := []struct {
		name    string
		value   float32
		wantErr bool
	}{
		{
			name:    "Valid Temp (positive celsius)",
			value:   22.4,
			wantErr: false,
		},
		{
			name:    "Valid Temp (negative celsius)",
			value:   -12.3,
			wantErr: false,
		},
		{
			name:    "Valid Temp (equal to absolute zero)",
			value:   -273.15,
			wantErr: false,
		},
		{
			name:    "Valid Temp (less then absolute zero)",
			value:   -273.16,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewTempCelsius(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTempCelsius() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTemperature_GetTemp(t *testing.T) {
	temp, err := NewTempCelsius(22.4)

	assert.Nil(t, err)
	assert.Equal(t, float32(22.4), temp.TempC)
	assert.Equal(t, float32(72.32), temp.TempF)
	assert.Equal(t, float32(295.4), temp.TempK)

}

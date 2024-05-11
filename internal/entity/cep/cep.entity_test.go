package entity

import (
	"testing"
)

func TestNewCep(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{
			name:    "Valid CEP (formatted)",
			value:   "123456-78",
			wantErr: false,
		},
		{
			name:    "Valid CEP (full formatted)",
			value:   "12.3456-78",
			wantErr: false,
		},
		{
			name:    "Valid CEP (with strange character)",
			value:   "123456@78",
			wantErr: false,
		},
		{
			name:    "Valid CEP (only numbers)",
			value:   "12345678",
			wantErr: false,
		},
		{
			name:    "Invalid CEP (too short)",
			value:   "1234567",
			wantErr: true,
		},
		{
			name:    "Invalid CEP (too short with character)",
			value:   "123456a",
			wantErr: true,
		},
		{
			name:    "Invalid CEP (too long)",
			value:   "123456789",
			wantErr: true,
		},
		{
			name:    "Invalid CEP (too long with character)",
			value:   "123456@789",
			wantErr: true,
		},
		{
			name:    "Invalid CEP (empty)",
			value:   "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewCep(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCep() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCep_GetCep(t *testing.T) {
	cep := &Cep{Value: "12345678"}
	if got := cep.GetCep(); got != "12345678" {
		t.Errorf("Cep.GetCep() = %v, want %v", got, "12345678")
	}
}

func TestCep_GetCepFormatted(t *testing.T) {
	cep := &Cep{Value: "12345678"}
	if got := cep.GetCepFormatted(); got != "12.345-678" {
		t.Errorf("Cep.GetCepFormatted() = %v, want %v", got, "12.345-678")
	}
}

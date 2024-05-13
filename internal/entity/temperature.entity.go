package entity

import "errors"

type Temperature struct {
	TempC float32 `json:"temp_c"`
	TempF float32 `json:"temp_f"`
	TempK float32 `json:"temp_k"`
}

func NewTempCelsius(value float32) (*Temperature, error) {
	if value < -273.15 {
		return nil, errors.New("invalid celsius value")
	}
	temp := &Temperature{
		TempC: value,
	}
	temp.TempF = temp.cToF()
	temp.TempK = temp.cToK()

	return temp, nil
}

func (t *Temperature) cToF() float32 {
	return t.TempC*1.8 + 32.0
}

func (t *Temperature) cToK() float32 {
	return t.TempC + 273.0
}

func (t *Temperature) GetTemp() *Temperature {
	return t
}

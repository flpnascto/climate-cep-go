package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/flpnascto/climate-cep-go/internal/entity"
)

type Current struct {
	TempC float64 `json:"temp_c"`
}

type WeatherDataResponse struct {
	Current Current `json:"current"`
}

func temperatureMapper(t WeatherDataResponse) entity.Temperature {
	temp, err := entity.NewTempCelsius(float32(t.Current.TempC))
	if err != nil {
		panic(err)
	}
	return *temp
}

func FetchWeatherApi(city string, apiKey string) (entity.Temperature, error) {
	// configs, err := configs.LoadConfig(".")
	// if err != nil {
	// 	panic(err)
	// }

	cityFormatted := url.QueryEscape(strings.ToLower(city))

	baseUrl := "https://api.weatherapi.com/v1"
	endpoint := "/current.json?"
	// key := "key=" + configs.WeatherApiKey
	key := "key=" + apiKey
	city = "&q=" + cityFormatted
	url := baseUrl + endpoint + key + city

	res, err := http.Get(url)
	if err != nil {
		return entity.Temperature{}, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return entity.Temperature{}, err
	}
	defer res.Body.Close()

	var result WeatherDataResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}

	temp := temperatureMapper(result)
	return temp, nil
}

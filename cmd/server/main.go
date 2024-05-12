package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/flpnascto/climate-cep-go/configs"
	"github.com/flpnascto/climate-cep-go/internal/entity"
)

type CepApiResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type Current struct {
	TempC float64 `json:"temp_c"`
}

type WeatherDataResponse struct {
	Current Current `json:"current"`
}

func TemperatureMapper(t WeatherDataResponse) entity.Temperature {
	temp, err := entity.NewTempCelsius(float32(t.Current.TempC))
	if err != nil {
		panic(err)
	}
	return *temp
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")

		cepQuery := parts[len(parts)-1]
		cep, err := entity.NewCep(cepQuery)
		if err != nil {
			http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
			return
		}

		city, err := fetchCepApi(cep)
		if err != nil {
			http.Error(w, "can not find zipcode", http.StatusNotFound)
			return
		}

		cityFormatted := url.QueryEscape(strings.ToLower(*city))

		temp, err := fetchWeatherApi(cityFormatted)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tempBytes, err := json.Marshal(temp.GetTemp())
		if err != nil {
			panic(err)
		}

		w.Write(tempBytes)
	})

	http.ListenAndServe(":8080", mux)
}

func fetchCepApi(c *entity.Cep) (*string, error) {

	res, err := http.Get("https://viacep.com.br/ws/" + c.GetCep() + "/json/")
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result CepApiResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}

	return &result.Localidade, nil
}

func fetchWeatherApi(city string) (entity.Temperature, error) {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	baseUrl := "https://api.weatherapi.com/v1"
	endpoint := "/current.json?"
	key := "key=" + configs.WeatherApiKey
	city = "&q=" + city
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

	temp := TemperatureMapper(result)
	return temp, nil
}

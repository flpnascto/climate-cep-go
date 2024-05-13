package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/flpnascto/climate-cep-go/configs"
	"github.com/flpnascto/climate-cep-go/internal/entity"
	"github.com/flpnascto/climate-cep-go/internal/infra/api"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")

		cepQuery := parts[len(parts)-1]
		cep, err := entity.NewCep(cepQuery)
		if err != nil {
			http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
			return
		}

		city, err := api.FetchCepApi(cep)
		if err != nil {
			http.Error(w, "can not find zipcode", http.StatusNotFound)
			return
		}

		temp, err := api.FetchWeatherApi(*city, configs.WeatherApiKey)
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

package api

import (
	"encoding/json"
	"io"
	"net/http"

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

func FetchCepApi(c *entity.Cep) (*string, error) {

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

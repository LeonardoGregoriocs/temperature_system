package zipcode

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/leonardogregoriocs/temperature_system/temperature_system/internal/utils"
)

type Data struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"Ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func GetAddressViaCep(zipcode string) (*Data, error) {
	ok, err := utils.Isvalid(zipcode)
	if err != nil || !ok {
		return nil, errors.New("invalid zipcode")
	}

	resp, err := http.Get("http://viacep.com.br/ws/" + zipcode + "/json")
	if err != nil {
		return &Data{}, err
	}

	if resp.Status == "NotFound" {
		return nil, errors.New("can not find zipcode")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != err {
		return &Data{}, err
	}

	var data Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		return &Data{}, err
	}

	return &data, nil
}

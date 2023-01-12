package gcp

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	jsonContentType          = "Content-Type: application/json"
	NeuralMachineTranslation = "nmt"
)

type GetLanguagesResponse struct {
	Data GetSupportedLanguagesResponseList `json:"data"`
}

type GetSupportedLanguagesResponseList struct {
	Languages []GetSupportedLanguagesResponseLanguage `json:"languages"`
}

type GetSupportedLanguagesResponseLanguage struct {
	Language string `json:"language"`
	Name     string `json:"name,omitempty"`
}

func Languages(hc *http.Client, target, model, key string) ([]GetSupportedLanguagesResponseLanguage, error) {
	lu := LanguagesUrl(target, model, key)

	resp, err := hc.Get(lu.String())
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New(resp.Status)
	}

	var lr *GetLanguagesResponse
	err = json.NewDecoder(resp.Body).Decode(&lr)

	return lr.Data.Languages, err
}

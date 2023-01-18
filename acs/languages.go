package acs

import (
	"encoding/json"
	"errors"
	"net/http"
)

type TranslationLanguage struct {
	Name       string `json:"name"`
	NativeName string `json:"nativeName"`
	Dir        string `json:"dir"`
}

type TranslationLanguagesResponse struct {
	Translation map[string]TranslationLanguage `json:"translation"`
}

func TranslationLanguages(hc *http.Client, acceptLanguage string) (*TranslationLanguagesResponse, error) {
	lu := LanguagesUrl(TranslationScope)

	req, err := http.NewRequest(http.MethodGet, lu.String(), nil)
	if err != nil {
		return nil, err
	}

	if acceptLanguage != "" {
		req.Header.Add("Accept-Language", acceptLanguage)
	}

	resp, err := hc.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New(resp.Status)
	}

	var tlr *TranslationLanguagesResponse
	err = json.NewDecoder(resp.Body).Decode(&tlr)

	return tlr, err
}

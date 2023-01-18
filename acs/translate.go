package acs

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/beauxarts/polyglot"
	"net/http"
)

type TranslationsResponse struct {
	Translations []Translation `json:"translations"`
}

type Translation struct {
	Text string `json:"text"`
	To   string `json:"to"`
}

func Translate(hc *http.Client, content []string, from, to string, format polyglot.TranslateFormat, key string) ([]*TranslationsResponse, error) {
	textData := NewTextData(content...)

	data, err := json.Marshal(textData)
	if err != nil {
		return nil, err
	}

	tu := TranslateUrl(from, to, format)

	translateReq, err := http.NewRequest(http.MethodPost, tu.String(), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	translateReq.Header.Add("Content-Type", applicationJsonContentType)
	translateReq.Header.Add("Ocp-Apim-Subscription-Key", key)

	resp, err := hc.Do(translateReq)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New(resp.Status)
	}

	var tr []*TranslationsResponse
	err = json.NewDecoder(resp.Body).Decode(&tr)

	return tr, err
}

package acs

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type TextData struct {
	Text string `json:"text"`
}

func NewTextData(content ...string) []TextData {
	dd := make([]TextData, len(content))
	for i, c := range content {
		dd[i].Text = c
	}
	return dd
}

const (
	applicationJsonContentType = "application/json; charset=UTF-8"
)

type DetectResponse struct {
	Language                   string  `json:"language"`
	Score                      float64 `json:"score"`
	IsTranslationSupported     bool    `json:"isTranslationSupported"`
	IsTransliterationSupported bool    `json:"isTransliterationSupported"`
}

func Detect(hc *http.Client, content, key string) (*DetectResponse, error) {
	dr, err := Detects(hc, key, content)
	if err != nil {
		return nil, err
	}

	if len(dr) > 0 {
		return dr[0], nil
	}

	return nil, nil
}

func Detects(hc *http.Client, key string, content ...string) ([]*DetectResponse, error) {

	textData := NewTextData(content...)

	data, err := json.Marshal(textData)
	if err != nil {
		return nil, err
	}

	du := DetectUrl()

	detectReq, err := http.NewRequest(http.MethodPost, du.String(), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	detectReq.Header.Add("Content-Type", applicationJsonContentType)
	detectReq.Header.Add("Ocp-Apim-Subscription-Key", key)

	resp, err := hc.Do(detectReq)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New(resp.Status)
	}

	var dr []*DetectResponse
	err = json.NewDecoder(resp.Body).Decode(&dr)

	return dr, err
}

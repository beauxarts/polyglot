package gcp

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type DetectRequest struct {
	Query string `json:"q"`
}

type DetectResponse struct {
	Data DetectLanguageResponseList `json:"data"`
}

type DetectLanguageResponseList struct {
	Detections [][]DetectionsListValue `json:"detections"`
}

type DetectionsListValue struct {
	Confidence int    `json:"confidence"`
	IsReliable bool   `json:"isReliable"`
	Language   string `json:"language"`
}

func Detect(hc *http.Client, query string, key string) ([]DetectionsListValue, error) {
	dreq := &DetectRequest{Query: query}

	data, err := json.Marshal(dreq)
	if err != nil {
		return nil, err
	}

	du := DetectUrl(key)

	resp, err := hc.Post(du.String(), applicationJsonContentType, bytes.NewReader(data))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New(resp.Status)
	}

	var dr *DetectResponse
	err = json.NewDecoder(resp.Body).Decode(&dr)

	return dr.Data.Detections[0], err
}

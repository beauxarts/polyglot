package gcp

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type DetectResp struct {
	Data struct {
		Detections [][]struct {
			Confidence int    `json:"confidence"`
			IsReliable bool   `json:"isReliable"`
			Language   string `json:"language"`
		} `json:"detections"`
	} `json:"data"`
}

func PostDetect(hc *http.Client, q string, key string) (*DetectResp, error) {
	tr := &TranslateReq{Query: q}

	req, err := json.Marshal(tr)
	if err != nil {
		return nil, err
	}

	du := DetectUrl(key)

	resp, err := hc.Post(du.String(), jsonContentType, bytes.NewReader(req))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New(resp.Status)
	}

	var dr *DetectResp
	err = json.NewDecoder(resp.Body).Decode(&dr)

	return dr, err
}

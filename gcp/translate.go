package gcp

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type TranslateReq struct {
	Query  string `json:"q"`
	Target string `json:"target,omitempty"`
}

type TranslateResp struct {
	Data struct {
		Translations []struct {
			TranslatedText         string `json:"translatedText"`
			DetectedSourceLanguage string `json:"detectedSourceLanguage"`
		} `json:"translations"`
	} `json:"data"`
}

func PostTranslate(hc *http.Client, q, targetLang, key string) (*TranslateResp, error) {
	treq := &TranslateReq{Query: q, Target: targetLang}

	req, err := json.Marshal(treq)
	if err != nil {
		return nil, err
	}

	tu := TranslateUrl(key)

	resp, err := hc.Post(tu.String(), jsonContentType, bytes.NewReader(req))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New(resp.Status)
	}

	var tresp *TranslateResp
	err = json.NewDecoder(resp.Body).Decode(&tresp)

	return tresp, err
}

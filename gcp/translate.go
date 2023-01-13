package gcp

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/beauxarts/polyglot"
	"net/http"
)

type TranslateRequest struct {
	Query  []string `json:"q"`
	Target string   `json:"target,omitempty"`
	Format string   `json:"format,omitempty"`
	Source string   `json:"source,omitempty"`
	Model  string   `json:"model,omitempty"`
}

type TranslateResponse struct {
	Data TranslateTextResponseList `json:"data"`
}

type TranslateTextResponseList struct {
	Translations []TranslateTextResponseTranslation `json:"translations"`
}

type TranslateTextResponseTranslation struct {
	DetectedSourceLanguage string `json:"detectedSourceLanguage"`
	Model                  string `json:"model"`
	TranslatedText         string `json:"translatedText"`
}

func Translate(hc *http.Client, query []string, target string, format polyglot.TranslateFormat, source, model, key string) ([]TranslateTextResponseTranslation, error) {
	if len(query) > 128 {
		return nil, errors.New("the maximum number of query strings is 128")
	}

	treq := &TranslateRequest{
		Query:  query,
		Target: target,
		Format: string(format),
		Source: source,
		Model:  model,
	}

	data, err := json.Marshal(treq)
	if err != nil {
		return nil, err
	}

	tu := TranslateUrl(key)

	resp, err := hc.Post(tu.String(), jsonContentType, bytes.NewReader(data))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New(resp.Status)
	}

	var tresp *TranslateResponse
	err = json.NewDecoder(resp.Body).Decode(&tresp)

	return tresp.Data.Translations, err
}

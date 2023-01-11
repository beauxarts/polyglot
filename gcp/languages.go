package gcp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	jsonContentType = "Content-Type: application/json"
)

type LanguagesResp struct {
	Data struct {
		Languages []LanguageResp `json:"languages"`
	} `json:"data"`
}

type LanguageResp struct {
	Language string `json:"language"`
	Name     string `json:"name,omitempty"`
}

func GetLanguages(hc *http.Client, key string) (*LanguagesResp, error) {
	lu := LanguagesUrl(key)

	fmt.Println(lu)

	resp, err := hc.Get(lu.String())
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New(resp.Status)
	}

	var lr *LanguagesResp
	err = json.NewDecoder(resp.Body).Decode(&lr)

	return lr, err
}

type PostLanguageReq struct {
	Target string `json:"target"`
}

func PostLanguages(hc *http.Client, targetLang string, key string) (*LanguagesResp, error) {
	plr := &PostLanguageReq{Target: targetLang}

	req, err := json.Marshal(plr)
	if err != nil {
		return nil, err
	}

	lu := LanguagesUrl(key)

	resp, err := hc.Post(lu.String(), jsonContentType, bytes.NewReader(req))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New(resp.Status)
	}

	var lr *LanguagesResp
	err = json.NewDecoder(resp.Body).Decode(&lr)

	return lr, err
}

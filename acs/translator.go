package acs

import (
	"errors"
	"github.com/beauxarts/polyglot"
	"net/http"
)

type Translator struct {
	httpClient *http.Client
	key        string
}

func NewTranslator(hc *http.Client, key string) (polyglot.Translator, error) {
	if key == "" {
		return nil, errors.New("key is required")
	}

	return &Translator{
		httpClient: hc,
		key:        key,
	}, nil
}

func (t *Translator) Languages(language string) (map[string]string, error) {

	lr, err := TranslationLanguages(t.httpClient, language)
	if err != nil {
		return nil, err
	}

	langName := make(map[string]string)

	for lc, ln := range lr.Translation {
		langName[lc] = ln.Name
	}

	return langName, nil
}

func (t *Translator) Detect(content string) (string, error) {

	dr, err := Detect(t.httpClient, content, t.key)
	if err != nil {
		return "", err
	}

	return dr.Language, nil
}

func (t *Translator) Translate(source, target string, format polyglot.TranslateFormat, query ...string) ([]string, error) {

	tr, err := Translate(t.httpClient, query, source, target, format, t.key)
	if err != nil {
		return nil, err
	}

	translations := make([]string, len(tr))
	for i, translation := range tr {
		if len(translation.Translations) > 0 {
			translations[i] = translation.Translations[0].Text
		}
	}

	return translations, nil
}

func (t *Translator) IsHTMLSupported() bool {
	return true
}

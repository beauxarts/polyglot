package gcp

import (
	"errors"
	"github.com/beauxarts/polyglot"
	"net/http"
)

type Translator struct {
	httpClient *http.Client
	model      string
	key        string
}

func NewTranslator(hc *http.Client, model, key string) (polyglot.Translator, error) {
	if key == "" {
		return nil, errors.New("key is required")
	}

	return &Translator{
		httpClient: hc,
		model:      model,
		key:        key,
	}, nil
}

func (t *Translator) Languages(targetLanguage string) (map[string]string, error) {

	lr, err := Languages(t.httpClient, targetLanguage, t.model, t.key)
	if err != nil {
		return nil, err
	}

	langName := make(map[string]string)

	for _, ln := range lr {
		langName[ln.Language] = ln.Name
	}

	return langName, nil
}

func (t *Translator) Detect(content string) (string, error) {

	dlv, err := Detect(t.httpClient, content, t.key)
	if err != nil {
		return "", err
	}

	if len(dlv) > 0 {
		return dlv[0].Language, nil
	}

	return "", nil
}

func (t *Translator) Translate(source, target string, format polyglot.TranslateFormat, query ...string) ([]string, error) {

	tresp, err := Translate(t.httpClient, query, target, format, source, t.model, t.key)
	if err != nil {
		return nil, err
	}

	translations := make([]string, 0, len(tresp))

	for _, tr := range tresp {
		translations = append(translations, tr.TranslatedText)
	}

	return translations, nil
}

func (t *Translator) IsHTMLSupported() bool {
	return true
}

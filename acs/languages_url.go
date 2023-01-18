package acs

import (
	"net/url"
	"strings"
)

const (
	TranslationScope     = "translation"
	TransliterationScope = "transliteration"
	DictionaryScope      = "dictionary"

	apiVersionParam = "api-version"
)

func LanguagesUrl(scopes ...string) *url.URL {
	u := &url.URL{
		Scheme: httpsScheme,
		Host:   cognitiveServicesAPIHost,
		Path:   languagesPath,
	}

	q := u.Query()
	q.Set(apiVersionParam, currentAPIVersion)
	if len(scopes) > 0 {
		q.Set("scope", strings.Join(scopes, ","))
	}
	u.RawQuery = q.Encode()

	return u
}

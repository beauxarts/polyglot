package acs

import (
	"github.com/beauxarts/polyglot"
	"net/url"
)

const (
	fromParam     = "from"
	toParam       = "to"
	textTypeParam = "textType"
)

var translateFormatsMap = map[polyglot.TranslateFormat]string{
	polyglot.Text: "plain",
	polyglot.HTML: "html",
}

func TranslateUrl(from, to string, format polyglot.TranslateFormat) *url.URL {
	u := &url.URL{
		Scheme: httpsScheme,
		Host:   cognitiveServicesAPIHost,
		Path:   translatePath,
	}

	q := u.Query()
	q.Set(apiVersionParam, currentAPIVersion)
	q.Set(fromParam, from)
	q.Set(toParam, to)
	if tt, ok := translateFormatsMap[format]; ok {
		q.Set(textTypeParam, tt)
	}
	u.RawQuery = q.Encode()

	return u
}

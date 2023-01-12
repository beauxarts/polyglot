package gcp

import "net/url"

const (
	targetParam = "target"
	modelParam  = "model"
	keyParam    = "key"
)

func LanguagesUrl(target, model, key string) *url.URL {
	if key == "" {
		return nil
	}

	u := &url.URL{
		Scheme: httpsScheme,
		Host:   translationAPIHost,
		Path:   languagesPath,
	}

	q := u.Query()
	if target != "" {
		q.Set(targetParam, target)
	}
	if model != "" {
		q.Set(modelParam, model)
	}
	q.Set(keyParam, key)
	u.RawQuery = q.Encode()

	return u
}

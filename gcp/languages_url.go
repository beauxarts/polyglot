package gcp

import "net/url"

const (
	keyParam = "key"
)

func LanguagesUrl(key string) *url.URL {
	if key == "" {
		return nil
	}

	u := &url.URL{
		Scheme: httpsScheme,
		Host:   translationAPIHost,
		Path:   languagesPath,
	}

	q := u.Query()
	q.Set(keyParam, key)
	u.RawQuery = q.Encode()

	return u
}

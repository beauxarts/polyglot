package gcp

import "net/url"

func TranslateUrl(key string) *url.URL {
	if key == "" {
		return nil
	}

	u := &url.URL{
		Scheme: httpsScheme,
		Host:   translationAPIHost,
		Path:   languageTranslateV2Path,
	}

	q := u.Query()
	q.Set(keyParam, key)
	u.RawQuery = q.Encode()

	return u
}

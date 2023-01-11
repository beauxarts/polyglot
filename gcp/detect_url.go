package gcp

import "net/url"

func DetectUrl(key string) *url.URL {
	if key == "" {
		return nil
	}

	u := &url.URL{
		Scheme: httpsScheme,
		Host:   translationAPIHost,
		Path:   detectPath,
	}

	q := u.Query()
	q.Set(keyParam, key)
	u.RawQuery = q.Encode()

	return u
}

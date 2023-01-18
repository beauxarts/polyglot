package acs

import (
	"net/url"
)

func DetectUrl() *url.URL {
	u := &url.URL{
		Scheme: httpsScheme,
		Host:   cognitiveServicesAPIHost,
		Path:   detectPath,
	}

	q := u.Query()
	q.Set(apiVersionParam, currentAPIVersion)
	u.RawQuery = q.Encode()

	return u
}

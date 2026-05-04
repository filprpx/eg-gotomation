package restapi

import (
	"crypto/tls"
	"net/http"
)

func SkipTLSVerify(httpClient *http.Client) {
	transport, ok := httpClient.Transport.(*http.Transport)
	if !ok || transport == nil {
		transport = &http.Transport{}
	}
	if transport.TLSClientConfig == nil {
		transport.TLSClientConfig = &tls.Config{}
	}
	transport.TLSClientConfig.InsecureSkipVerify = true
	httpClient.Transport = transport
}

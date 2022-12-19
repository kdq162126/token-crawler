package utils

import (
	"net/http"
	"net/url"
	"os"
	"strings"
)

func PassProxy(client *http.Client) error {
	proxyEnv := os.Getenv("PROXY_ENDPOINT")
	if len(proxyEnv) == 0 {
		// fmt.Println("skip proxy")
		return nil
	}
	if strings.HasPrefix(proxyEnv, "http") {
		proxyUrl, err := url.Parse(proxyEnv)
		if err != nil {
			return err
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	}
	return nil
}

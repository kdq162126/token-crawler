package utils

import (
	"fmt"
	"net/http"
	"time"
)

func ProxyHttpClient() *http.Client {
	client := &http.Client{
		Timeout: 600 * time.Second,
	}
	err := PassProxy(client)
	if err != nil {
		fmt.Println(err)
	}
	return client
}

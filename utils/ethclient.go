package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func DialProxy(url string) (*ethclient.Client, error) {
	client := http.Client{Timeout: 60 * time.Second}
	err := PassProxy(&client)
	if err != nil {
		fmt.Println(err)
	}
	c, err := rpc.DialHTTPWithClient(url, &client)
	if err != nil {
		return nil, err
	}
	return ethclient.NewClient(c), nil
}

func RpcDialProxy(url string) (*rpc.Client, error) {
	client := http.Client{Timeout: 60 * time.Second}
	err := PassProxy(&client)
	if err != nil {
		fmt.Println(err)
	}
	c, err := rpc.DialHTTPWithClient(url, &client)
	if err != nil {
		return nil, err
	}
	return c, nil
}

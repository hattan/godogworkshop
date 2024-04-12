package sdk

import (
	"fmt"
	"io"
	"log"
	"net/http"

	animals "github.com/hattan/godog"
)

type GoDogOptions struct {
	Host string
	Port string
}

type GoDogClient struct {
	options GoDogOptions
}

func NewGoDogClient(options GoDogOptions) *GoDogClient {
	return &GoDogClient{
		options: options,
	}
}

func (client *GoDogClient) GetDogs() []animals.Dog {
	host := client.getUri()
	uri := fmt.Sprintf("%s/%s", host, "dog")
	dogs := make([]animals.Dog, 0)
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	_ = sb
	return dogs
}

func (client *GoDogClient) getUri() string {
	host := client.options.Host
	if client.options.Port != "" {
		host = host + ":" + client.options.Port
	}
	return host
}

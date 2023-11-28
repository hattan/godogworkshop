package httpclient

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	animals "github.com/hattan/godog"
)

func GetDogs() []*animals.Dog {
	url := "http://localhost:8080/dog"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var dogs []*animals.Dog
	json.Unmarshal(data, &dogs)
	return dogs
}

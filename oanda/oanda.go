package oanda

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Api struct {
	ApiKey    string
	AccountID string
	Endpoint  string
}

var Client http.Client

func init() {
	Client = http.Client{
		Timeout: time.Duration(10 * time.Second),
	}
}

func (a *Api) NewRequest(method string, url string, data interface{}) (*http.Request, error) {
	fullURL := fmt.Sprintf("%s/%s", a.Endpoint, url)
	body, _ := json.Marshal(data)
	req, _ := http.NewRequest(method, fullURL, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.ApiKey))

	return req, nil
}

// req = http.Request
// target = response struct
func (a *Api) MakeRequest(req *http.Request, target interface{}) error {
	res, err := Client.Do(req)
	if err != nil {
		log.Fatal(err)
		defer res.Body.Close()
	}

	if res.StatusCode >= 400 {
		var message map[string]string
		_ = json.NewDecoder(res.Body).Decode(&message)
		return errors.New(message["errorMessage"])
	}

	return json.NewDecoder(res.Body).Decode(target)
}

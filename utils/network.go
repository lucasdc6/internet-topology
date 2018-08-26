package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetGeneric - Get request to url with unmarshall in res
func GetGeneric(url string, res interface{}) (error error) {
	resp, err := http.Get(url)
	if err != nil {
		error = err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		error = err
	}

	err = json.Unmarshal([]byte(body), &res)

	if err != nil {
		error = err
	}

	return error
}

package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/lucasdc6/internet-topology/environment"
)

var debug = environment.GetDebugFor("bgpview") || environment.GetDebug()

// GetGeneric - Get request to url with unmarshall in res
func GetGeneric(url string, res interface{}) (error error) {
	if debug {
		fmt.Printf("GetGeneric from: %s\n", url)
	}
	time.Sleep(500 * time.Millisecond)

	resp, err := http.Get(url)
	if err != nil {
		return error
	}

	if debug {
		fmt.Printf("Read data\n")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return error
	}

	if debug {
		fmt.Printf("Unmarshal data\n")
	}
	err = json.Unmarshal([]byte(body), &res)

	if err != nil {
		return error
	}

	if debug {
		fmt.Printf("Finish GetGeneric\n")
	}
	return nil
}

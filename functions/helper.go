// Package functions provides the operations needed to perform POST requests on the
// desired object looking to pull data on.
package functions	

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mmendez3800/startgg-go/objects"
)

var (
	AuthToken string
	client = &http.Client{
		Timeout: time.Second * 10,
	}
)

// tokenSet verifies if the Start GG authentication token has been set by the user.
func tokenSet() bool {
	if AuthToken == "" {
		return false
	} else {
		return true
	}
}

// createQuery constructs the GraphQL query that is used for the POST request.
func createQuery(queryString string, searchKey string, searchValue interface{}) map[string]interface{} {
	return map[string]interface{} {
		"query": queryString,
		"variables": map[string]interface{} {
			searchKey: searchValue,
		},
	}
}

// validateData checks if the data received in the POST request is valid.
func validateData(data []byte) string {
	results := &objects.FailedCall{}
	err := json.Unmarshal(data, results)
	if err != nil {
		return "Unable To Validate Data"
	}

	return results.Message
}

// runQuery performs the POST request to the Start GG API.
func runQuery(query []byte) ([]byte, error) {
	// Creates the POST request and checks for errors.
	req, err := http.NewRequest("POST", "https://api.start.gg/gql/alpha", bytes.NewBuffer(query))
	if err != nil {
		return nil, fmt.Errorf("HTTP Request - %w", err)
	}

	// Sets the headers within the request.
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AuthToken))

	// Sends the request and receives the response of it.
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP Response - %w", err)
	}
	defer res.Body.Close()

	// Reads the data of the response.
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Read Data - %w", err)
	}

	// Validates the data received by the response of the request.
	validation := validateData(data)
	if validation != "" {
		return nil, fmt.Errorf("Data Validation - %s", validation)
	}

	return data, nil
}

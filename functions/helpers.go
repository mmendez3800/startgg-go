package functions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mmendez3800/startgg-go/objects"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	contentType = "application/json"
	url = "https://api.start.gg/gql/alpha"
	queryKey = "query"
	variablesKey = "variables"
)

var (
	AuthToken string
	client = &http.Client{
		Timeout: time.Second * 10,
	}
)

type queryData map[string]interface{}

func tokenSet() bool {
	if AuthToken == "" {
		return false
	} else {
		return true
	}
}

func createQuery(queryString string, searchKey string, searchValue interface{}) queryData {
	return queryData {
		queryKey: queryString,
		variablesKey: queryData {
			searchKey: searchValue,
		},
	}
}

func validateData(data []byte) string {
	results := &objects.FailedCall{}
	err := json.Unmarshal(data, results)
	if err != nil {
		return "Unable To Validate Data"
	}

	return results.Message
}

func runQuery(query []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(query))
	if err != nil {
		return nil, fmt.Errorf("HTTP Request - %v", err)
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AuthToken))

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP Response - %v", err)
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Read Data - %v", err)
	}

	validation := validateData(data)
	if validation != "" {
		return nil, fmt.Errorf("Data Validation - %s", validation)
	}

	return data, nil
}

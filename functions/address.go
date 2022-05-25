package functions

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"github.com/mmendez3800/startgg-go/objects"
	"github.com/mmendez3800/startgg-go/queries"
	"net/http"
)

func GetAddress(key string, value interface{}) (*objects.SmashGGAddress, error) {
	query, err := json.Marshal(createQuery(queries.AddressQuery, key, value))
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, contentType, bytes.NewBuffer(query))
	if err != nil {
		return nil, err
	}

	defer closeBody(resp.Body)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	results := new(objects.SmashGGAddress)
	err = json.Unmarshal(data, results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

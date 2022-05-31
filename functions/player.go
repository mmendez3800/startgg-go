// Package functions provides the operations needed to perform POST requests on the
// desired object looking to pull data on.
package functions

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mmendez3800/startgg-go/objects"
	"github.com/mmendez3800/startgg-go/queries"
)

// GetPlayer provides either the data of the Player object received based on the parameters
// given or the error that resulted in the attempt.
func GetPlayer(key string, value interface{}) (*objects.StartGGPlayer, error) {
	// Checks that the authentication token has been set.
	if (tokenSet() == false) {
		return nil, errors.New("Token Verification - Authentication Token Not Set")
	}

	// Creates a JSON encoding of the query being made.
	query, err := json.Marshal(createQuery(queries.PlayerQuery, key, value))
	if err != nil {
		return nil, fmt.Errorf("JSON Marshal - %w", err)
	}

	// Runs the query on the encoding and returns the results.
	data, err := runQuery(query)
	if err != nil {
		return nil, err
	}

	// Parses the JSON encoding and interprets the results based on the object's definition.
	results := &objects.StartGGPlayer{}
	err = json.Unmarshal(data, results)
	if err != nil {
		return nil, fmt.Errorf("JSON Unmarshal - %w", err)
	}

	return results, nil
}

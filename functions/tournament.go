package functions

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mmendez3800/startgg-go/objects"
	"github.com/mmendez3800/startgg-go/queries"
)

func GetTournament(key string, value interface{}) (*objects.StartGGTournament, error) {
	if (tokenSet() == false) {
		return nil, errors.New("Token Verification - Authorization Token Not Set")
	}

	query, err := json.Marshal(createQuery(queries.TournamentQuery, key, value))
	if err != nil {
		return nil, fmt.Errorf("JSON Marshal - %v", err)
	}

	data, err := runQuery(query)
	if err != nil {
		return nil, err
	}

	results := &objects.StartGGTournament{}
	err = json.Unmarshal(data, results)
	if err != nil {
		return nil, fmt.Errorf("JSON Unmarshal - %v", err)
	}

	return results, nil
}

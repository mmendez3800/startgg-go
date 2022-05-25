package functions

import (
	"io"
)

const (
	contentType = "application/json"
	url = "https://api.start.gg/gql/alpha"
	queryKey = "query"
	variablesKey = "variables"
)

type queryData map[string]interface{}

func createQuery(queryString string, searchKey string, searchValue interface{}) queryData {
	return queryData {
		queryKey: queryString,
		variablesKey: queryData {
			searchKey: searchValue,
		},
	}
}

func closeBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		return
	}
}

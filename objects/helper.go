// Package objects defines the structures of the StartGG objects that can be called upon
// when performing the API calls.
package objects

type Errors struct {
	Message string `json:"message"`
	Path []string `json:"path"`
	Locations []Locations `json:"locations"`
	Extensions Extensions `json:"extensions"`
}

type Locations struct {
	Line int `json:"line"`
	Column int `json:"column"`
}

type Extensions struct {
	Category string `json:"category"`
}

type FailedCall struct {
	Succes bool `json:"success"`
	Message string `json:"message"`
}

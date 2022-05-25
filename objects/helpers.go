package objects

type Errors struct {
	Message string `json:"message"`
	Locations []Locations `json:"locations"`
	Path []string `json:"path"`
	Extensions Extensions `json:"extensions"`
}

type Locations struct {
	Line int `json:"line"`
	Column int `json:"column"`
}

type Extensions struct {
	Code string `json:"code"`
	Timestamp string `json:"timestamp"`
}

package model

type header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type parameter struct {
	T     string `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type request struct {
	Method         string      `json:"method"`
	Url            string      `json:"url"`
	RequestLine    string      `json:"requestline"`
	HeaderNames    []string    `json:"headernames"`
	Headers        []header    `json:"headers"`
	ParameterNames []string    `json:"parameternames"`
	Parameters     []parameter `json:"parameters"`
	Body           string      `json:"body"`
}

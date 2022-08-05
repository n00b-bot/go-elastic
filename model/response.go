package model

type cookie struct {
	Domain     string `json:"domain"`
	Expiration string `json:"expiration"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Value      string `json:"value"`
}

type response struct {
	Status                int      `json:"status"`
	ResponseLine          string   `json:"responseline"`
	Content_type          string   `json:"content_type"`
	Inferred_content_type string   `json:"inferred_content_type"`
	HeaderNames           []string `json:"header_names"`
	Headers               []header `json:"headers"`
	Cookies               []cookie `json:"cookies"`
	Body                  string   `json:"body"`
}

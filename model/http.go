package model

type Http struct {
	Timestamp string   `json:"timestamp"`
	Protocol  string   `json:"protocol"`
	Host      string   `json:"host"`
	Port      int      `json:"port"`
	Request   request  `json:"request"`
	Response  response `json:"response"`
}

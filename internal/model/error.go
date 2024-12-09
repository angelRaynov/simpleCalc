package model

type Error struct {
	Expression string `json:"expression"`
	Endpoint   string `json:"endpoint"`
	Frequency  int    `json:"frequency"`
	Type       string `json:"type"`
}

type ErrorCLI struct {
	Error string `json:"error"`
}

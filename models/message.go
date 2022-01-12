package models

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

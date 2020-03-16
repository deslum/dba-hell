package types

type Message struct {
	Id        uint64 `json:"id"`
	Name      string `json:"name"`
	Number    int    `json:"number"`
	Body      string `json:"body"`
	Timestamp int64  `json:"ts"`
}

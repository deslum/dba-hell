package types

type Message struct {
	Id        uint64 `json:"id"`
	Body      string `json:"b"`
	Timestamp int64  `json:"ts"`
}

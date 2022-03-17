package server

type Data struct {
	Client    string  `json:"client"`
	Port      string  `json:"port"`
	Address   string  `json:"address"`
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}

package serial

type Translator struct {
	Translator string `json:"TRANSLATOR"`
	Timestamp  string `json:"TIMESTAMP"`
	// map[seasonNumber]map[episodeNumber]token
	Tokens map[int64]map[int64]string `json:"TOKENS"`
}

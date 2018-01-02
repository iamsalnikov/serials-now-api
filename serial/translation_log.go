package serial

type TranslationLog struct {
	E          int64  `json:"E,string"`
	S          int64  `json:"S,string"`
	Timestamp  string `json:"TIMESTAMP"`
	Translator string `json:"TRANSLATOR"`
}

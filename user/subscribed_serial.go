package user

type SubscribedSerial struct {
	ID           int64   `json:"ID,string"`
	IMDB         float32 `json:"IMDB,string"`
	Kinopoisk    float32 `json:"KINOPOISK,string"`
	Year         int32   `json:"YEAR,string"`
	TranslatorID int64   `json:"T,string"`
	Translator   string  `json:"TRANSLATOR"`
	TitleRu      string  `json:"TITLE_RU"`
}

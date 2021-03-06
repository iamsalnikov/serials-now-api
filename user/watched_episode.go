package user

type WatchedEpisode struct {
	SerialID     int64   `json:"ID,string"`
	TranslatorID int64   `json:"T,string"`
	Translator   string  `json:"TRANSLATOR"`
	Season       int64   `json:"S,string"`
	Episode      int64   `json:"E,string"`
	IMDB         float32 `json:"IMDB,string"`
	Kinopoisk    float32 `json:"KINOPOISK,string"`
	TitleRu      string  `json:"TITLE_RU"`
	Year         int32   `json:"YEAR,string"`
}

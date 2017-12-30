package index

type Serial struct {
	ID        int64   `json:"ID,string"`
	IMDB      float32 `json:"IMDB,string"`
	Kinopoisk float32 `json:"KINOPOISK,string"`
	Year      int32   `json:"YEAR,string"`
	TitleRu   string  `json:"TITLE_RU"`
}

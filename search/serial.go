package search

type Serial struct {
	Comments int64 `json:"COMMENTS"`
	Country string `json:"COUNTRY"`
	Dislike int64 `json:"DISLIKE"`
	Favorites int64 `json:"FAVORITES"`
	Genre string `json:"GENRE"`
	ID int64 `json:"ID"`
	IMDB float32 `json:"IMDB"`
	Kinopoisk float32 `json:"KINOPOISK"`
	Like int64 `json:"LIKE"`
	Stat int64 `json:"STAT"`
	Subscriptions int64 `json:"SUBSCRIPTIONS"`
	TitleEn string `json:"TITLE_EN"`
	TitleRu string `json:"TITLE_RU"`
	Watched int64 `json:"WATCHED"`
	Year int64 `json:"YEAR"`
}

package serial

type Serial struct {
	Comments      int64   `json:"COMMENTS,string"`
	Country       string  `json:"COUNTRY"`
	Description   string  `json:"DESCRIPTION"`
	Dislike       int64   `json:"DISLIKE,string"`
	Favorites     int64   `json:"FAVORITES,string"`
	Genre         string  `json:"GENRE"`
	ID            int64   `json:"ID,string"`
	IMDB          float32 `json:"IMDB,string"`
	Kinopoisk     float32 `json:"KINOPOISK,string"`
	Like          int64   `json:"LIKE,string"`
	Stat          int64   `json:"STAT,string"`
	Subscriptions int64   `json:"SUBSCRIPTIONS,string"`
	TitleEn       string  `json:"TITLE_EN"`
	TitleRu       string  `json:"TITLE_RU"`
	Watched       int64   `json:"WATCHED,string"`
	Year          int32   `json:"YEAR,string"`
	Update        int64   `json:"UPDATE,string"`
	Vote          int64   `json:"VOTE,string"`
}

package serial

type EndpointData struct {
	Serial                   Serial                   `json:"0"`
	Comments                 []Comment                `json:"1,omitempty"`
	Translators              map[int64]Translator     `json:"2,omitempty"`
	Share                    Share                    `json:"3,omitempty"`
	WatchHistory             WatchHistory             `json:"5,omitempty"`
	TranslationSubscriptions TranslationSubscriptions `json:"6,omitempty"`
	TranslationLog           []TranslationLog         `json:"7,omitempty"`
}

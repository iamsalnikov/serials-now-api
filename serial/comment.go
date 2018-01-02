package serial

type Comment struct {
	Text      string `json:"TEXT"`
	Timestamp string `json:"TIMESTAMP"`
	User      int64  `json:"USER,string"`
}

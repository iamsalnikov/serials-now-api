package user

type Comment struct {
	SerialID  int64  `json:"ID,string"`
	Text      string `json:"TEXT"`
	Timestamp string `json:"TIMESTAMP"`
}

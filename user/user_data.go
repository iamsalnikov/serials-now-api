package user

type UserData struct {
	FavoriteEpisodes  []FavoriteEpisode
	WatchedEpisodes   []WatchedEpisode
	NewEpisodes       []NewEpisode
	SubscribedSerials []SubscribedSerial
	VotedSerials      []VotedSerial
	Comments          []Comment
}

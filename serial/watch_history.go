package serial

// map[translationID]map[seasonNumber]map[episodeNumber]flag (true or false)
type WatchHistory map[int64]map[int64]map[int64]bool

package tracker

type Tracker struct {
	Name    string `json:"name"`
	Source  string `json:"source"`
	Created int64  `json:"created"`
	Updated int64  `json:"updated"`
}

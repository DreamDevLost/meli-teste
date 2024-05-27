package responses

type GetItemDescription struct {
	Text        string   `json:"text"`
	PlainText   string   `json:"plain_text"`
	LastUpdated string   `json:"last_updated"`
	DateCreated string   `json:"date_created"`
	Snapshot    Snapshot `json:"snapshot"`
}

type Snapshot struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
	Status string `json:"status"`
}

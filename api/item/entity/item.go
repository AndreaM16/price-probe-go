package itementity

type Item struct {
	ID          string `json:"id"`
	Pid         string `json:"pid"`
	URL         string `json:"url"`
	Category    string `json:"category"`
	Img         string `json:"img"`
	Description string `json:"description"`
	Title       string `json:"title"`
}

type Items struct {
	Items []*Item
}

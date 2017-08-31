package priceentity

type Price struct {
	Item      string  `json:"item"`
	Date      []int   `json:"date"`
	Price     float64 `json:"price"`
	Estimated float64 `json:"estimated"`
}

type Prices struct {
	Prices []Price `json:"prices"`
}

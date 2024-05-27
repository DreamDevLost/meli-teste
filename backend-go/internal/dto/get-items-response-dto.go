package dto

type GetItemsResponseDto struct {
	Author     Author   `json:"author"`
	Categories []string `json:"categories"`
	Items      []*Item  `json:"items"`
}

type Author struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}

type Item struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Price        Price  `json:"price"`
	Picture      string `json:"picture"`
	Condition    string `json:"condition"`
	FreeShipping bool   `json:"free_shipping"`
}

type Price struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Decimals int     `json:"decimals"`
}

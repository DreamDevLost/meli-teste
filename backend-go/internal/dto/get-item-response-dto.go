package dto

type GetItemResponseDto struct {
	Author Author     `json:"author"`
	Item   SingleItem `json:"item"`
}

type SingleItem struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Price        Price  `json:"price"`
	Picture      string `json:"picture"`
	Condition    string `json:"condition"`
	FreeShipping bool   `json:"free_shipping"`
	SoldQuantity int64  `json:"sold_quantity"`
	Description  string `json:"description"`
}

package responses

import "time"

type GetItemsResponse struct {
	SiteID                 string            `json:"site_id"`
	CountryDefaultTimeZone string            `json:"country_default_time_zone"`
	Query                  string            `json:"query"`
	Paging                 Paging            `json:"paging"`
	Results                []Result          `json:"results"`
	Sort                   Sort              `json:"sort"`
	AvailableSorts         []Sort            `json:"available_sorts"`
	Filters                []Filter          `json:"filters"`
	AvailableFilters       []AvailableFilter `json:"available_filters"`
	PDPTracking            PDPTracking       `json:"pdp_tracking"`
}

type AvailableFilter struct {
	ID     string                 `json:"id"`
	Name   string                 `json:"name"`
	Type   string                 `json:"type"`
	Values []AvailableFilterValue `json:"values"`
}

type AvailableFilterValue struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Results int64  `json:"results"`
}

type Sort struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Filter struct {
	ID     string        `json:"id"`
	Name   string        `json:"name"`
	Type   string        `json:"type"`
	Values []FilterValue `json:"values"`
}

type FilterValue struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	PathFromRoot []Sort `json:"path_from_root,omitempty"`
}

type PDPTracking struct {
	Group       bool          `json:"group"`
	ProductInfo []ProductInfo `json:"product_info"`
}

type ProductInfo struct {
	ID     string `json:"id"`
	Score  int64  `json:"score"`
	Status string `json:"status"`
}

type Paging struct {
	Total          int64 `json:"total"`
	PrimaryResults int64 `json:"primary_results"`
	Offset         int64 `json:"offset"`
	Limit          int64 `json:"limit"`
}

type Result struct {
	ID                  string                     `json:"id"`
	Title               string                     `json:"title"`
	Condition           string                     `json:"condition"`
	ThumbnailID         string                     `json:"thumbnail_id"`
	CatalogProductID    *string                    `json:"catalog_product_id"`
	ListingTypeID       string                     `json:"listing_type_id"`
	Permalink           string                     `json:"permalink"`
	BuyingMode          string                     `json:"buying_mode"`
	SiteID              string                     `json:"site_id"`
	CategoryID          string                     `json:"category_id"`
	DomainID            string                     `json:"domain_id"`
	Thumbnail           string                     `json:"thumbnail"`
	CurrencyID          string                     `json:"currency_id"`
	OrderBackend        int64                      `json:"order_backend"`
	Price               float64                    `json:"price"`
	OriginalPrice       float64                    `json:"original_price"`
	SalePrice           interface{}                `json:"sale_price"`
	AvailableQuantity   int64                      `json:"available_quantity"`
	OfficialStoreID     *int64                     `json:"official_store_id"`
	UseThumbnailID      bool                       `json:"use_thumbnail_id"`
	AcceptsMercadopago  bool                       `json:"accepts_mercadopago"`
	Shipping            Shipping                   `json:"shipping"`
	StopTime            time.Time                  `json:"stop_time"`
	Seller              Seller                     `json:"seller"`
	Attributes          []ResultAttribute          `json:"attributes"`
	Installments        Installments               `json:"installments"`
	WinnerItemID        interface{}                `json:"winner_item_id"`
	CatalogListing      bool                       `json:"catalog_listing"`
	Discounts           interface{}                `json:"discounts"`
	Promotions          []interface{}              `json:"promotions"`
	InventoryID         interface{}                `json:"inventory_id"`
	VariationID         *string                    `json:"variation_id,omitempty"`
	VariationFilters    []string                   `json:"variation_filters,omitempty"`
	VariationsData      map[string]VariationsDatum `json:"variations_data,omitempty"`
	DifferentialPricing *DifferentialPricing       `json:"differential_pricing,omitempty"`
	OfficialStoreName   *string                    `json:"official_store_name,omitempty"`
}

type ResultAttribute struct {
	ID                 string           `json:"id"`
	Name               string           `json:"name"`
	ValueID            *string          `json:"value_id"`
	ValueName          *string          `json:"value_name"`
	AttributeGroupID   string           `json:"attribute_group_id"`
	AttributeGroupName string           `json:"attribute_group_name"`
	ValueStruct        *Struct          `json:"value_struct"`
	Values             []AttributeValue `json:"values"`
	Source             int64            `json:"source"`
	ValueType          string           `json:"value_type"`
}

type Struct struct {
	Number interface{} `json:"number"`
	Unit   string      `json:"unit"`
}

type AttributeValue struct {
	ID     *string `json:"id"`
	Name   *string `json:"name"`
	Struct *Struct `json:"struct"`
	Source int64   `json:"source"`
}

type DifferentialPricing struct {
	ID int64 `json:"id"`
}

type Installments struct {
	Quantity   int64   `json:"quantity"`
	Amount     float64 `json:"amount"`
	Rate       float64 `json:"rate"`
	CurrencyID string  `json:"currency_id"`
}

type Seller struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
}

type Shipping struct {
	StorePickUp   bool        `json:"store_pick_up"`
	FreeShipping  bool        `json:"free_shipping"`
	LogisticType  string      `json:"logistic_type"`
	Mode          string      `json:"mode"`
	Tags          []string    `json:"tags"`
	Benefits      interface{} `json:"benefits"`
	Promise       interface{} `json:"promise"`
	ShippingScore int64       `json:"shipping_score"`
}

type VariationsDatum struct {
	Thumbnail     string                     `json:"thumbnail"`
	Ratio         string                     `json:"ratio"`
	Name          string                     `json:"name"`
	PicturesQty   int64                      `json:"pictures_qty"`
	Price         float64                    `json:"price"`
	UserProductID string                     `json:"user_product_id"`
	Attributes    []VariationsDatumAttribute `json:"attributes"`
}

type VariationsDatumAttribute struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ValueName string `json:"value_name"`
	ValueType string `json:"value_type"`
}

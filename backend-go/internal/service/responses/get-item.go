package responses

import "time"

type GetItemResponse struct {
	ID              string      `json:"id"`
	SiteID          string      `json:"site_id"`
	Title           string      `json:"title"`
	SellerID        int64       `json:"seller_id"`
	CategoryID      string      `json:"category_id"`
	OfficialStoreID interface{} `json:"official_store_id"`
	Price           float64     `json:"price"`
	BasePrice       float64     `json:"base_price"`
	OriginalPrice   float64     `json:"original_price"`
	CurrencyID      string      `json:"currency_id"`
	InitialQuantity int64       `json:"initial_quantity"`
	SaleTerms       []Attribute `json:"sale_terms"`
	BuyingMode      string      `json:"buying_mode"`
	ListingTypeID   string      `json:"listing_type_id"`
	Condition       string      `json:"condition"`
	Permalink       string      `json:"permalink"`
	// ThumbnailID                  string        `json:"thumbnail_id"`
	Thumbnail string    `json:"thumbnail"`
	Pictures  []Picture `json:"pictures"`
	// VideoID                      interface{}   `json:"video_id"`
	// Descriptions                 []interface{} `json:"descriptions"`
	// AcceptsMercadopago           bool          `json:"accepts_mercadopago"`
	// NonMercadoPagoPaymentMethods []interface{} `json:"non_mercado_pago_payment_methods"`
	Shipping Shipping `json:"shipping"`
	// InternationalDeliveryMode    string        `json:"international_delivery_mode"`
	// SellerAddress                SellerAddress `json:"seller_address"`
	// SellerContact                interface{}   `json:"seller_contact"`
	Location Location `json:"location"`
	// CoverageAreas    []interface{} `json:"coverage_areas"`
	Attributes       []Attribute   `json:"attributes"`
	ListingSource    string        `json:"listing_source"`
	Variations       []interface{} `json:"variations"`
	Status           string        `json:"status"`
	SubStatus        []interface{} `json:"sub_status"`
	Tags             []string      `json:"tags"`
	Warranty         string        `json:"warranty"`
	CatalogProductID string        `json:"catalog_product_id"`
	DomainID         string        `json:"domain_id"`
	ParentItemID     interface{}   `json:"parent_item_id"`
	DealIDS          []interface{} `json:"deal_ids"`
	AutomaticRelist  bool          `json:"automatic_relist"`
	DateCreated      time.Time     `json:"date_created"`
	LastUpdated      time.Time     `json:"last_updated"`
	Health           interface{}   `json:"health"`
	CatalogListing   bool          `json:"catalog_listing"`
}

type Attribute struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	ValueID     *string `json:"value_id"`
	ValueName   string  `json:"value_name"`
	Values      []Value `json:"values"`
	ValueType   string  `json:"value_type"`
	ValueStruct *Struct `json:"value_struct"`
}

type Value struct {
	ID     *string `json:"id"`
	Name   string  `json:"name"`
	Struct *Struct `json:"struct"`
}

type Location struct {
}

type Picture struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	SecureURL string `json:"secure_url"`
	Size      string `json:"size"`
	MaxSize   string `json:"max_size"`
	Quality   string `json:"quality"`
}

type SellerAddress struct {
	City           City           `json:"city"`
	State          Country        `json:"state"`
	Country        Country        `json:"country"`
	SearchLocation SearchLocation `json:"search_location"`
	ID             int64          `json:"id"`
}

type City struct {
	Name string `json:"name"`
}

type Country struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SearchLocation struct {
	Neighborhood Country `json:"neighborhood"`
	City         Country `json:"city"`
	State        Country `json:"state"`
}

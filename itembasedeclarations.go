package itembase

type Profile struct {
	Active    bool   `json:"active"`
	AvatarURL string `json:"avatar_url"`
	Contact   struct {
		Contact []Contact `json:"contact"`
	} `json:"contact"`
	CreatedAt         string `json:"created_at"`
	Currency          string `json:"currency"`
	DisplayName       string `json:"display_name"`
	ID                string `json:"id"`
	Language          string `json:"language"`
	Locale            string `json:"locale"`
	OriginalReference string `json:"original_reference"`
	PlatformID        string `json:"platform_id"`
	PlatformName      string `json:"platform_name"`
	SourceID          string `json:"source_id"`
	Status            string `json:"status"`
	Type              string `json:"type"`
	UpdatedAt         string `json:"updated_at"`
	URL               string `json:"url"`
}

type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
	Line1   string `json:"line_1"`
	Name    string `json:"name"`
	Zip     string `json:"zip"`
}

type Contact struct {
	Addresses []Address `json:"addresses"`
	Emails    []struct {
		Value string `json:"value"`
	} `json:"emails"`
	Phones []interface{} `json:"phones"`
}

type Buyer struct {
	Active            bool    `json:"active"`
	Contact           Contact `json:"contact"`
	CreatedAt         string  `json:"created_at"`
	Currency          string  `json:"currency"`
	DateOfBirth       string  `json:"date_of_birth"`
	FirstName         string  `json:"first_name"`
	ID                string  `json:"id"`
	Language          string  `json:"language"`
	LastName          string  `json:"last_name"`
	Locale            string  `json:"locale"`
	Note              string  `json:"note"`
	OptOut            bool    `json:"opt_out"`
	OriginalReference string  `json:"original_reference"`
	SourceID          string  `json:"source_id"`
	Status            string  `json:"status"`
	Type              string  `json:"type"`
	UpdatedAt         string  `json:"updated_at"`
	URL               string  `json:"url"`
}

type Category struct {
	CategoryID string `json:"category_id"`
	Language   string `json:"language"`
	Value      string `json:"value"`
}

type ProductDescription struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}

type Brand struct {
	Name struct {
		Language string `json:"language"`
		Value    string `json:"value"`
	} `json:"name"`
}

type Product struct {
	Active     bool       `json:"active"`
	Brand      Brand      `json:"brand"`
	Categories []Category `json:"categories"`
	Condition  string     `json:"condition"`
	CreatedAt  string     `json:"created_at"`
	Currency   struct {
		Currency string `json:"currency"`
	} `json:"currency"`
	Description []ProductDescription `json:"description"`
	ID          string               `json:"id"`
	Identifier  struct {
		ID string `json:"id"`
	} `json:"identifier"`
	Name []struct {
		Language string `json:"language"`
		Value    string `json:"value"`
	} `json:"name"`
	OriginalReference string `json:"original_reference"`
	PictureUrls       []struct {
		URLOriginal string `json:"url_original"`
	} `json:"picture_urls"`
	PricePerUnit float64 `json:"price_per_unit"`
	Shipping     []struct {
		Price           float64 `json:"price"`
		ShippingService string  `json:"shipping_service"`
	} `json:"shipping"`
	SourceID         string `json:"source_id"`
	StockInformation struct {
		InStock        bool    `json:"in_stock"`
		InventoryLevel float64 `json:"inventory_level"`
		InventoryUnit  string  `json:"inventory_unit"`
	} `json:"stock_information"`
	Tax       float64       `json:"tax"`
	TaxRate   float64       `json:"tax_rate"`
	UpdatedAt string        `json:"updated_at"`
	URL       string        `json:"url"`
	Variants  []interface{} `json:"variants"`
}

type Billing struct {
	Address Address `json:"address"`
}

type Transaction struct {
	Billing           Billing   `json:"billing"`
	Buyer             Buyer     `json:"buyer"`
	CreatedAt         string    `json:"created_at"`
	Currency          string    `json:"currency"`
	ID                string    `json:"id"`
	OriginalReference string    `json:"original_reference"`
	Products          []Product `json:"products"`
	Shipping          struct {
		Address Address `json:"address"`
	} `json:"shipping"`
	SourceID string `json:"source_id"`
	Status   struct {
		Global   string `json:"global"`
		Payment  string `json:"payment"`
		Shipping string `json:"shipping"`
	} `json:"status"`
	TotalPrice    float64 `json:"total_price"`
	TotalPriceNet float64 `json:"total_price_net"`
	TotalTax      float64 `json:"total_tax"`
	UpdatedAt     string  `json:"updated_at"`
}

type Transactions struct {
	Transactions         []Transaction `json:"documents"`
	NumDocumentsFound    float64       `json:"num_documents_found"`
	NumDocumentsReturned float64       `json:"num_documents_returned"`
}

type Profiles struct {
	Profiles             []Profile `json:"documents"`
	NumDocumentsFound    float64   `json:"num_documents_found"`
	NumDocumentsReturned float64   `json:"num_documents_returned"`
}

type Products struct {
	Products             []Product `json:"documents"`
	NumDocumentsFound    float64   `json:"num_documents_found"`
	NumDocumentsReturned float64   `json:"num_documents_returned"`
}

type Buyers struct {
	Buyers               []Buyer `json:"documents"`
	NumDocumentsFound    float64 `json:"num_documents_found"`
	NumDocumentsReturned float64 `json:"num_documents_returned"`
}

type User struct {
	UUID              string `json:"uuid"`
	Username          string `json:"username"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	MiddleName        string `json:"middle_name"`
	NameFormat        string `json:"name_format"`
	Locale            string `json:"locale"`
	Email             string `json:"email"`
	PreferredCurrency string `json:"preffered_currency"`
}

package itembase

// TODO: Some entities/models don't have the full set of fields from the API.
// Some of the implementation detail structs (Contacts, Billing, pagination
// containers, etc.) could perhaps be unexported.

// A Profile represents a user profile entity from the itembase API.
//
// See http://sandbox.api.itembase.io/swagger-ui/
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

// An Address represents a mailing address model from the itembase API.
type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
	Line1   string `json:"line_1"`
	Name    string `json:"name"`
	Zip     string `json:"zip"`
}

// A Contact represents a container of contact information from itembase API
// models.
type Contact struct {
	Addresses []Address `json:"addresses"`
	Emails    []struct {
		Value string `json:"value"`
	} `json:"emails"`
	Phones []interface{} `json:"phones"`
}

// A Buyer represents a buyer entity from the itembase API.
//
// See http://sandbox.api.itembase.io/swagger-ui/
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

// A Category represents a product category model from the itembase API.
type Category struct {
	CategoryID string `json:"category_id"`
	Language   string `json:"language"`
	Value      string `json:"value"`
}

// A ProductDescription represents a product description model from the itembase
// API, which may be in a specified language.
type ProductDescription struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}

// A Brand represents a product brand model from the itembase API.
type Brand struct {
	Name struct {
		Language string `json:"language"`
		Value    string `json:"value"`
	} `json:"name"`
}

// A Product represents a product entity from the itembase API.
//
// See http://sandbox.api.itembase.io/swagger-ui/
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

// Billing represents a model from the itembase API containing the billing
// address of a Transaction.
type Billing struct {
	Address Address `json:"address"`
}

// A Transaction represents a transaction entity from the itembase API.
//
// See http://sandbox.api.itembase.io/swagger-ui/
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

// Transactions is a container for pagination of Transaction entities.
type Transactions struct {
	Transactions         []Transaction `json:"documents"`
	NumDocumentsFound    float64       `json:"num_documents_found"`
	NumDocumentsReturned float64       `json:"num_documents_returned"`
}

// Profiles is a container for pagination of Profile entities.
type Profiles struct {
	Profiles             []Profile `json:"documents"`
	NumDocumentsFound    float64   `json:"num_documents_found"`
	NumDocumentsReturned float64   `json:"num_documents_returned"`
}

// Products is a container for pagination of Product entities.
type Products struct {
	Products             []Product `json:"documents"`
	NumDocumentsFound    float64   `json:"num_documents_found"`
	NumDocumentsReturned float64   `json:"num_documents_returned"`
}

// Buyers is a container for pagination of Buyer entities.
type Buyers struct {
	Buyers               []Buyer `json:"documents"`
	NumDocumentsFound    float64 `json:"num_documents_found"`
	NumDocumentsReturned float64 `json:"num_documents_returned"`
}

// A User represents a user entity from the itembase API, such as returned from
// the "me" endpoint.
type User struct {
	UUID              string `json:"uuid"`
	Username          string `json:"username"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	MiddleName        string `json:"middle_name"`
	NameFormat        string `json:"name_format"`
	Locale            string `json:"locale"`
	Email             string `json:"email"`
	PreferredCurrency string `json:"preferred_currency"`
}

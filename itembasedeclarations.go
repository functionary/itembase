package itembase

import (
	"encoding/json"
	"log"
)

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
	Active      bool                 `json:"active"`
	Brand       Brand                `json:"brand"`
	Categories  []Category           `json:"categories"`
	Condition   string               `json:"condition"`
	CreatedAt   string               `json:"created_at"`
	Currency    string               `json:"currency"`
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

func (p *Product) InStock() bool {
	if p.StockInformation.InStock {
		return true
	}
	return false
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

func (t *Transaction) Completed() bool {
	if t.Status.Global == "completed" {
		return true
	}
	return false
}

// ItembaseResponse is a container for any Itembase response.
// It returns the resultset, Number of found documents and Number of documents returned
type ItembaseResponse struct {
	Documents            []interface{} `json:"documents"`
	NumDocumentsFound    int           `json:"num_documents_found"`
	NumDocumentsReturned int           `json:"num_documents_returned"`
}

// Transactions is a container for pagination of Transaction entities.
type Transactions struct {
	Transactions []Transaction `json:"documents"`
}

func (t *Transactions) Add(transaction interface{}) {

	var newTransaction Transaction
	convertTo(transaction, &newTransaction)
	t.Transactions = append(t.Transactions, newTransaction)

}

func (t *Transactions) Completed() (transactions []Transaction) {
	for _, transaction := range t.Transactions {
		if transaction.Completed() {
			transactions = append(transactions, transaction)
		}
	}

	return
}

// Profiles is a container for pagination of Profile entities.
type Profiles struct {
	Profiles []Profile `json:"documents"`
}

func (p *Profiles) Add(profile interface{}) {

	var newProfile Profile
	convertTo(profile, &newProfile)
	p.Profiles = append(p.Profiles, newProfile)

}

// Products is a container for pagination of Product entities.
type Products struct {
	Products []Product `json:"documents"`
}

func (p *Products) Add(product interface{}) {

	var newProduct Product
	convertTo(product, &newProduct)
	p.Products = append(p.Products, newProduct)

}

func (p *Products) InStock() (products []Product) {
	for _, product := range p.Products {
		if product.InStock() {
			products = append(products, product)
		}
	}

	return
}

// Buyers is a container for pagination of Buyer entities.
type Buyers struct {
	Buyers []Buyer `json:"documents"`
}

func (b *Buyers) Add(buyer interface{}) {

	var newBuyer Buyer
	convertTo(buyer, &newBuyer)
	b.Buyers = append(b.Buyers, newBuyer)

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

func convertTo(inputInterface, outputType interface{}) {
	jsonBLOB, err := json.Marshal(inputInterface)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(jsonBLOB, &outputType)
	if err != nil {
		log.Fatal(err)
	}
}

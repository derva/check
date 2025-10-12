package config

var EnvVars EnvVariables

type EnvVariables struct {
	VeryfiAPIKEY      string
	VeryfiURL         string
	VeryfiAPIUsername string
	VeryfiClientID    string
	DBHost            string
	DBUsername        string
	DBPassword        string
	DBPort            string
	DBName            string
}

type InnerData struct {
	AccountNumber           string      `json:"account_number"`
	BillTo                  BillTo      `json:"bill_to"`
	Cashback                interface{} `json:"cashback"`
	Category                string      `json:"category"`
	CountryCode             string      `json:"country_code"`
	CreatedDate             string      `json:"created_date"`
	CurrencyCode            string      `json:"currency_code"`
	Date                    string      `json:"date"`
	DeliveryDate            interface{} `json:"delivery_date"`
	Discount                interface{} `json:"discount"`
	DocumentReferenceNumber interface{} `json:"document_reference_number"`
	DocumentTitle           string      `json:"document_title"`
	DocumentType            string      `json:"document_type"`
	DueDate                 interface{} `json:"due_date"`
	DuplicateOf             int         `json:"duplicate_of"`
	ExternalID              interface{} `json:"external_id"`
	ID                      int         `json:"id"`
	ImgFileName             string      `json:"img_file_name"`
	ImgThumbnailURL         string      `json:"img_thumbnail_url"`
	ImgURL                  string      `json:"img_url"`
	Insurance               interface{} `json:"insurance"`
	InvoiceNumber           string      `json:"invoice_number"`
	IsDuplicate             bool        `json:"is_duplicate"`
	IsMoneyIn               bool        `json:"is_money_in"`
	LineItems               []LineItem  `json:"line_items"`
	Meta                    MetaData    `json:"meta"`
	Notes                   interface{} `json:"notes"`
	OcrText                 string      `json:"ocr_text"`
	Payment                 Payment     `json:"payment"`
	PDFURL                  string      `json:"pdf_url"`
	ReferenceNumber         string      `json:"reference_number"`
	Subtotal                float64     `json:"subtotal"`
	Tax                     float64     `json:"tax"`
	Total                   float64     `json:"total"`
	UpdatedDate             string      `json:"updated_date"`
	Vendor                  Vendor      `json:"vendor"`
}

type BillTo struct {
	Address       string  `json:"address"`
	Name          *string `json:"name"`
	ParsedAddress *string `json:"parsed_address"`
	VatNumber     *string `json:"vat_number"`
}

type LineItem struct {
	Date        interface{} `json:"date"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Quantity    float64     `json:"quantity"`
	ID          int         `json:"id"`
	Total       float64     `json:"total"`
}

type MetaData struct {
	Language       []string `json:"language"`
	OcrScore       float64  `json:"ocr_score"`
	Owner          string   `json:"owner"`
	ProcessedPages int      `json:"processed_pages"`
	Source         string   `json:"source"`
}

type Payment struct {
	CardNumber  *string `json:"card_number"`
	DisplayName string  `json:"display_name"`
	Type        string  `json:"type"`
}

type Vendor struct {
	Name      string `json:"name"`
	Category  string `json:"category"`
	VATNumber string `json:"vat_number"`
}

type OuterData struct {
	Message string    `json:"message"`
	Text    InnerData `json:"text"`
}

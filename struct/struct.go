package structs

type BillerCategory struct {
	Id          int      `json:"Id"`
	Name        string   `json:"Name"`
	Description string   `json:"Description"`
	Billers     []string `json:"Billers,omitempty"` // Optional since it might be empty
}

type BillerCategoriesResponse struct {
	BillerCategories     []BillerCategory `json:"BillerCategories"`
	ResponseCode         string           `json:"ResponseCode"`
	ResponseCodeGrouping string           `json:"ResponseCodeGrouping"`
}

type Biller struct {
	Type                   string `json:"Type,omitempty"`
	Id                     int    `json:"Id"`
	PayDirectProductId     int    `json:"PayDirectProductId"`
	PayDirectInstitutionId int    `json:"PayDirectInstitutionId"`
	Name                   string `json:"Name"`
	ShortName              string `json:"ShortName"`
	Narration              string `json:"Narration,omitempty"`
	CustomerField1         string `json:"CustomerField1,omitempty"`
	CustomerField2         string `json:"CustomerField2,omitempty"`
	LogoUrl                string `json:"LogoUrl,omitempty"`
	Url                    string `json:"Url,omitempty"`
	Surcharge              string `json:"Surcharge,omitempty"`
	CustomSectionUrl       string `json:"CustomSectionUrl,omitempty"`
	CurrencyCode           string `json:"CurrencyCode"`
	CurrencySymbol         string `json:"CurrencySymbol"`
	QuickTellerSiteUrlName string `json:"QuickTellerSiteUrlName"`
	SupportEmail           string `json:"SupportEmail,omitempty"`
	RiskCategoryId         string `json:"RiskCategoryId,omitempty"`
	PageFlowInfo           struct {
		Elements         []interface{} `json:"Elements"`
		FinishButtonName string        `json:"FinishButtonName"`
		StartPage        string        `json:"StartPage"`
		UsesPaymentItems bool          `json:"UsesPaymentItems"`
		PerformInquiry   bool          `json:"PerformInquiry"`
		AllowRetry       bool          `json:"AllowRetry"`
	} `json:"PageFlowInfo"`
	CategoryId    int    `json:"CategoryId"`
	CategoryName  string `json:"CategoryName"`
	MediumImageId string `json:"MediumImageId,omitempty"`
	SmallImageId  string `json:"SmallImageId,omitempty"`
	AmountType    int    `json:"AmountType"`
}

type Category struct {
	Id          int      `json:"Id"`
	Name        string   `json:"Name"`
	Description string   `json:"Description,omitempty"`
	Billers     []Biller `json:"Billers"`
}

type BillersIdResponse struct {
	BillerList struct {
		Count    int        `json:"Count"`
		Category []Category `json:"Category"`
	} `json:"BillerList"`
	ResponseCode         string `json:"ResponseCode"`
	ResponseCodeGrouping string `json:"ResponseCodeGrouping"`
}

type PaymentTransaction struct {
	Customers []struct {
		PaymentCode string `json:"PaymentCode"`
		CustomerID  string `json:"CustomerId"`
	} `json:"customers"`
	TerminalId string `json:"TerminalId"`
}

// ElectricityPaymentOption represents a single payment option for electricity bills.
type ElectricityPaymentOption struct {
	ID                                      int     `json:"id"`
	Name                                    string  `json:"name"`
	CountryCode                             string  `json:"countryCode"`
	CountryName                             string  `json:"countryName"`
	Type                                    string  `json:"type"`
	ServiceType                             string  `json:"serviceType"`
	LocalAmountSupported                    bool    `json:"localAmountSupported"`
	LocalTransactionCurrencyCode            string  `json:"localTransactionCurrencyCode"`
	MinLocalTransactionAmount               float64 `json:"minLocalTransactionAmount"`
	MaxLocalTransactionAmount               float64 `json:"maxLocalTransactionAmount"`
	LocalTransactionFee                     float64 `json:"localTransactionFee"`
	LocalTransactionFeeCurrencyCode         string  `json:"localTransactionFeeCurrencyCode"`
	LocalDiscountPercentage                 float64 `json:"localDiscountPercentage"`
	InternationalAmountSupported            bool    `json:"internationalAmountSupported"`
	InternationalTransactionCurrencyCode    string  `json:"internationalTransactionCurrencyCode"`
	MinInternationalTransactionAmount       float64 `json:"minInternationalTransactionAmount"`
	MaxInternationalTransactionAmount       float64 `json:"maxInternationalTransactionAmount"`
	InternationalTransactionFee             float64 `json:"internationalTransactionFee"`
	InternationalTransactionFeeCurrencyCode string  `json:"internationalTransactionFeeCurrencyCode"`
	LocalTransactionFeePercentage           float64 `json:"localTransactionFeePercentage"`
	InternationalTransactionFeePercentage   float64 `json:"internationalTransactionFeePercentage"`
	InternationalDiscountPercentage         float64 `json:"internationalDiscountPercentage"`
	Fx                                      struct {
		Rate         float64 `json:"rate"`
		CurrencyCode string  `json:"currencyCode"`
	} `json:"fx"`
	DenominationType          string      `json:"denominationType"`
	LocalFixedAmounts         interface{} `json:"localFixedAmounts"`
	InternationalFixedAmounts interface{} `json:"localFixedAmounts"`
}

type UtilBill struct {
	SubscriberAccountNumber string      `json:"subscriberAccountNumber"`
	Amount                  float64     `json:"amount"`
	AmountID                interface{} `json:"amountId"` // Can be null
	BillerID                int         `json:"billerId"`
	UseLocalAmount          bool        `json:"useLocalAmount"`
	ReferenceID             string      `json:"referenceId"`
	AdditionalInfo          struct {
		InvoiceID interface{} `json:"invoiceId"` // Can be null
	} `json:"additionalInfo"`
}

type Response struct {
	ResponseDescription string  `json:"response_description"`
	Content             Content `json:"content"`
}

type Content struct {
	ServiceName    string      `json:"ServiceName"`
	ServiceID      string      `json:"serviceID"`
	ConvenienceFee string      `json:"convinience_fee"`
	Variations     []Variation `json:"varations"`
}

type Variation struct {
	VariationCode   string `json:"variation_code"`
	Name            string `json:"name"`
	VariationAmount string `json:"variation_amount"`
	FixedPrice      string `json:"fixedPrice"`
}

type UserData struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type UserResponse struct {
	Status   string      `json:"status"`
	Message  string      `json:"message"`
	Error    string      `json:"error"`
	UserData interface{} `json:"userData"`
}

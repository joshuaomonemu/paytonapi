package Structs

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

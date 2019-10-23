package nct

// Define the data for this product here
// ******Start******

// AgreementComponent Agreement component information of NCT
type AgreementComponent struct {
	PolicyID      string  `json:"policyId"`
	Country       string  `json:"country"`
	InsuredPeople string  `json:"issuredPeople"`
	Gender        string  `json:"gender"`
	FaceAmount    float64 `json:"faceAmount"`
	Premium       float64 `json:"premium"`
}

// Policy information of NCT
type Policy struct {
	ObjectType		string	`json:"docType"`
	InsurancePolicyNo string `json:"insurancePolicyNo"`
	Status            string `json:"status"`
	StatusRemark      string `json:"statusRemark"`

	BankRefNo       string  `json:"bankRefNo"`
	ProdCatOverview string  `json:"prodCatOverview"`
	AgentCode       string  `json:"agentCode"`
	Premium         float64 `json:"premium"`
	Currency        string  `json:"currency"`   //HKD,CNY,MOP,USD
	PayMode         float64 `json:"payMode"`    //1,3,6,12
	PlanRegion      string  `json:"planRegion"` //HKG or PRC
}

// PolicyInsurancePrivate information of NCT
type PolicyInsurancePrivate struct {
	ObjectType			string	`json:"docType"`
	InsurancePolicyNo	string	`json:"insurancePolicyNo"`
	BankRefNo			string	`json:"bankRefNo"`
	BankRating			string	`json:"bankRating"`		//A,B,C,D
}

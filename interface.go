package paystack

import "time"

type Paystack struct {
	APIKey *string
}

type PostResponseData struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Reference        string `json:"reference"`
		AuthorizationURL string `json:"authorization_url"`
		AccessCode       string `json:"access_code"`
	} `json:"data"`
}

type GetResponseData struct {
	Status     bool            `json:"status"`
	Message    string          `json:"message"`
	// StatusCode int             `json:"statusCode"`
	Data       TransactionData `json:"data"`
}

type TransactionData struct {
	ID                 int64                  `json:"id"`
	Domain             string                 `json:"domain"`
	Status             string                 `json:"status"`
	Reference          string                 `json:"reference"`
	Amount             int                    `json:"amount"`
	Message            *string                `json:"message"` // Use a pointer to allow for null values
	GatewayResponse    string                 `json:"gateway_response"`
	PaidAt             time.Time              `json:"paid_at"`
	CreatedAt          time.Time              `json:"created_at"`
	Channel            string                 `json:"channel"`
	Currency           string                 `json:"currency"`
	IPAddress          string                 `json:"ip_address"`
	Metadata           string                 `json:"metadata"`
	Log                LogData                `json:"log"`
	Fees               int                    `json:"fees"`
	FeesSplit          *int                   `json:"fees_split"` // Use a pointer to allow for null values
	Authorization      AuthorizationData      `json:"authorization"`
	Customer           CustomerData           `json:"customer"`
	Plan               *string                `json:"plan"` // Use a pointer to allow for null values
	Split              map[string]interface{} `json:"split"`
	OrderID            *string                `json:"order_id"` // Use a pointer to allow for null values
	PaidAtISO          string                 `json:"paidAt"`
	CreatedAtISO       string                 `json:"createdAt"`
	RequestedAmount    int                    `json:"requested_amount"`
	POSTransactionData *string                `json:"pos_transaction_data"` // Use a pointer to allow for null values
	Source             *string                `json:"source"`               // Use a pointer to allow for null values
	FeesBreakdown      *string                `json:"fees_breakdown"`       // Use a pointer to allow for null values
	TransactionDate    string                 `json:"transaction_date"`
	PlanObject         map[string]interface{} `json:"plan_object"`
	Subaccount         map[string]interface{} `json:"subaccount"`
}

type LogData struct {
	StartTime int64     `json:"start_time"`
	TimeSpent int       `json:"time_spent"`
	Attempts  int       `json:"attempts"`
	Errors    int       `json:"errors"`
	Success   bool      `json:"success"`
	Mobile    bool      `json:"mobile"`
	Input     []string  `json:"input"`
	History   []History `json:"history"`
}

type History struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Time    int    `json:"time"`
}

type AuthorizationData struct {
	AuthorizationCode string  `json:"authorization_code"`
	BIN               string  `json:"bin"`
	Last4             string  `json:"last4"`
	ExpMonth          string  `json:"exp_month"`
	ExpYear           string  `json:"exp_year"`
	Channel           string  `json:"channel"`
	CardType          string  `json:"card_type"`
	Bank              string  `json:"bank"`
	CountryCode       string  `json:"country_code"`
	Brand             string  `json:"brand"`
	Reusable          bool    `json:"reusable"`
	Signature         string  `json:"signature"`
	AccountName       *string `json:"account_name"` // Use a pointer to allow for null values
}

type CustomerData struct {
	ID                       int64   `json:"id"`
	FirstName                *string `json:"first_name"` // Use a pointer to allow for null values
	LastName                 *string `json:"last_name"`  // Use a pointer to allow for null values
	Email                    string  `json:"email"`
	CustomerCode             string  `json:"customer_code"`
	Phone                    *string `json:"phone"`    // Use a pointer to allow for null values
	Metadata                 *string `json:"metadata"` // Use a pointer to allow for null values
	RiskAction               string  `json:"risk_action"`
	InternationalFormatPhone *string `json:"international_format_phone"` // Use a pointer to allow for null values
}

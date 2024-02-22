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
	Status  bool   `json:"status"`
	Message string `json:"message"`
	// StatusCode int             `json:"statusCode"`
	Data TransactionData `json:"data"`
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

// listbanks
type FilterBanks struct {
	Country                string `json:"country" schema:"country" validate:"required"`
	UseCursor              string `json:"use_cursor" schema:"use_cursor" validate:"required"`
	PerPage                int    `json:"perPage" schema:"perPage" validate:"required,min=1"`
	PayWithTransfer        *bool  `json:"pay_with_bank_transfer" schema:"pay_with_bank_transfer"`
	PayWithBank            *bool  `json:"pay_with_bank" schema:"pay_with_bank"`
	EnabledForVerification *bool  `json:"enabled_for_verification" schema:"enabled_for_verification"`
	Next                   string `json:"next" schema:"next"`
	Previous               string `json:"previous" schema:"previous"`
	Gateway                string `json:"gateway" schema:"gateway"`
	Type                   string `json:"type" schema:"type"`
	Currency               string `json:"currency" schema:"currency"`
}

type TransactionStatus string

const (
	Failed    TransactionStatus = "failed"
	Success   TransactionStatus = "success"
	Abandoned TransactionStatus = "abandoned"
)

// listbanks
type ListTransactions struct {
	PerPage    int               `json:"perPage" schema:"perPage" validate:"required,min=10"`
	Page       int               `json:"page" schema:"page" validate:"required,min=1"`
	Customer   string            `json:"customer" schema:"customer"`
	TerminalID string            `json:"terminalid" schema:"terminalid"`
	Status     TransactionStatus `json:"status" schema:"status"`
	Amount     int               `json:"amount" schema:"amount"`
	To         *time.Time        `json:"to" schema:"to" validate:"omitempty,timestamp"`
	From       *time.Time        `json:"from" schema:"from" validate:"omitempty,timestamp"`
}

type FullResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []struct {
		ID              int64       `json:"id"`
		Domain          string      `json:"domain"`
		Status          string      `json:"status"`
		Reference       string      `json:"reference"`
		Amount          int         `json:"amount"`
		Message         interface{} `json:"message"`
		GatewayResponse string      `json:"gateway_response"`
		PaidAt          time.Time   `json:"paid_at"`
		CreatedAt       time.Time   `json:"created_at"`
		Channel         string      `json:"channel"`
		Currency        string      `json:"currency"`
		IPAddress       string      `json:"ip_address"`
		Metadata        interface{} `json:"metadata"`
		Log             struct {
			StartTime int           `json:"start_time"`
			TimeSpent int           `json:"time_spent"`
			Attempts  int           `json:"attempts"`
			Errors    int           `json:"errors"`
			Success   bool          `json:"success"`
			Mobile    bool          `json:"mobile"`
			Input     []interface{} `json:"input"`
			History   []struct {
				Type    string `json:"type"`
				Message string `json:"message"`
				Time    int    `json:"time"`
			} `json:"history"`
		} `json:"log"`
		Fees      int         `json:"fees"`
		FeesSplit interface{} `json:"fees_split"`
		Customer  struct {
			ID           int         `json:"id"`
			FirstName    interface{} `json:"first_name"`
			LastName     interface{} `json:"last_name"`
			Email        string      `json:"email"`
			Phone        interface{} `json:"phone"`
			Metadata     interface{} `json:"metadata"`
			CustomerCode string      `json:"customer_code"`
			RiskAction   string      `json:"risk_action"`
		} `json:"customer"`
		Authorization struct {
			AuthorizationCode string      `json:"authorization_code"`
			Bin               string      `json:"bin"`
			Last4             string      `json:"last4"`
			ExpMonth          string      `json:"exp_month"`
			ExpYear           string      `json:"exp_year"`
			Channel           string      `json:"channel"`
			CardType          string      `json:"card_type"`
			Bank              string      `json:"bank"`
			CountryCode       string      `json:"country_code"`
			Brand             string      `json:"brand"`
			Reusable          bool        `json:"reusable"`
			Signature         string      `json:"signature"`
			AccountName       interface{} `json:"account_name"`
		} `json:"authorization"`
		Plan struct {
		} `json:"plan"`
		Split struct {
		} `json:"split"`
		Subaccount struct {
		} `json:"subaccount"`
		OrderID         interface{} `json:"order_id"`
		PaidAtSec          time.Time   `json:"paidAt"`
		CreatedAtSec       time.Time   `json:"createdAt"`
		RequestedAmount int         `json:"requested_amount"`
		Source          struct {
			Source     string      `json:"source"`
			Type       string      `json:"type"`
			Identifier interface{} `json:"identifier"`
			EntryPoint string      `json:"entry_point"`
		} `json:"source"`
		ConnectData struct {
		} `json:"connect_data"`
		PosTransactionData interface{} `json:"pos_transaction_data"`
	} `json:"data"`
	Meta struct {
		Total       int    `json:"total"`
		TotalVolume int    `json:"total_volume"`
		Skipped     int    `json:"skipped"`
		PerPage     string `json:"perPage"`
		Page        int    `json:"page"`
		PageCount   int    `json:"pageCount"`
	} `json:"meta"`
}

type MetaTransaction struct {
	Total     int    `json:"total"`
	Skipped   int    `json:"skipped"`
	PerPage   string `json:"perPage"`
	Page      int    `json:"page"`
	PageCount int    `json:"pageCount"`
}

type BankResponse struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Data    []Bank   `json:"data"`
	Meta    BankMeta `json:"meta"`
}

type Bank struct {
	Name             string `json:"name"`
	Slug             string `json:"slug"`
	Code             string `json:"code"`
	Longcode         string `json:"longcode"`
	Gateway          string `json:"gateway"` // Pointer to allow for null values
	PayWithBank      bool   `json:"pay_with_bank"`
	Active           bool   `json:"active"`
	SupportsTransfer bool   `json:"supports_transfer"`
	IsDeleted        bool   `json:"is_deleted"`
	Country          string `json:"country"`
	Currency         string `json:"currency"`
	Type             string `json:"type"`
	ID               int    `json:"id"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
}

type BankMeta struct {
	Next     string `json:"next"`     // Pointer to allow for null values
	Previous string `json:"previous"` // Pointer to allow for null values
	PerPage  int    `json:"perPage"`
}

type TransferInput struct {
	Amount    float64 `json:"amount" schema:"amount" validate:"required"`
	Recipient string  `json:"recipient" schema:"recipient" validate:"required"`
	Reason    string  `json:"reason" schema:"recipient" validate:"required"`
	Currency  string  `json:"currency"`
	Reference string  `json:"reference"`
}

type InitTransferResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Integration  int       `json:"integration"`
		Domain       string    `json:"domain"`
		Amount       int       `json:"amount"`
		Currency     string    `json:"currency"`
		Source       string    `json:"source"`
		Reason       string    `json:"reason"`
		Recipient    int       `json:"recipient"`
		Status       string    `json:"status"`
		TransferCode string    `json:"transfer_code"`
		ID           int       `json:"id"`
		CreatedAt    time.Time `json:"createdAt"`
		UpdatedAt    time.Time `json:"updatedAt"`
	} `json:"data"`
}

type ConfirmTransferInput struct {
	TransferCode string `json:"transfer_code"`
	OTP          string `json:"otp"`
}

type ConfirmTransferResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Domain        string      `json:"domain"`
		Amount        int         `json:"amount"`
		Currency      string      `json:"currency"`
		Reference     string      `json:"reference"`
		Source        string      `json:"source"`
		SourceDetails interface{} `json:"source_details"`
		Reason        string      `json:"reason"`
		Status        string      `json:"status"`
		Failures      interface{} `json:"failures"`
		TransferCode  string      `json:"transfer_code"`
		TitanCode     interface{} `json:"titan_code"`
		TransferredAt interface{} `json:"transferred_at"`
		ID            int         `json:"id"`
		Integration   int         `json:"integration"`
		Recipient     int         `json:"recipient"`
		CreatedAt     time.Time   `json:"createdAt"`
		UpdatedAt     time.Time   `json:"updatedAt"`
	} `json:"data"`
}

//failed transfer struct type Response struct {
type FailedTrasnferResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Meta    struct {
		NextStep string `json:"nextStep"`
	} `json:"meta"`
	Type string `json:"type"`
	Code string `json:"code"`
}

//create recipient
type AccountDetails struct {
	Type          string `json:"type" schema:"type" validate:"required"`
	Name          string `json:"name" schema:"name" validate:"required"`
	AccountNumber string `json:"account_number" schema:"account_number" validate:"required"`
	BankCode      string `json:"bank_code" schema:"bank_code" validate:"required"`
	Currency      string `json:"currency" schema:"currency" validate:"required"`
	Description   string `json:"description" schema:"description" validate:"required"`
}

type Recipient struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Active        bool      `json:"active"`
		CreatedAt     time.Time `json:"createdAt"`
		Currency      string    `json:"currency"`
		Domain        string    `json:"domain"`
		ID            int       `json:"id"`
		Integration   int       `json:"integration"`
		Name          string    `json:"name"`
		RecipientCode string    `json:"recipient_code"`
		Type          string    `json:"type"`
		UpdatedAt     time.Time `json:"updatedAt"`
		IsDeleted     bool      `json:"is_deleted"`
		Details       struct {
			AuthorizationCode *string `json:"authorization_code"`
			AccountNumber     string  `json:"account_number"`
			AccountName       *string `json:"account_name"`
			BankCode          string  `json:"bank_code"`
			BankName          string  `json:"bank_name"`
		} `json:"details"`
	} `json:"data"`
}

type Error struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

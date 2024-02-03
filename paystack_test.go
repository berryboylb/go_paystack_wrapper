package paystack

import (
	"reflect"
	"testing"
	"time"
)

// Mock response data for testing
var mockInitializeResponse = PostResponseData{
	Status:  true,
	Message: "Authorization URL created",
	Data: struct {
		Reference        string `json:"reference"`
		AuthorizationURL string `json:"authorization_url"`
		AccessCode       string `json:"access_code"`
	}{
		Reference:        "trx_1234567890abcdef",
		AuthorizationURL: "https://example.com/authorize/trx_1234567890abcdef",
		AccessCode:       "ACCESSCODE123456",
	},
}

var mockVerifyResponse = GetResponseData{
	Message: "Verification successful",
	Status:  true,
	Data: TransactionData{
		Status:          "success",
		Message:         nil, // Null value
		ID:              3516052615,
		Domain:          "test",
		Reference:       "0l2qk643pk",
		Amount:          2000000,
		GatewayResponse: "Successful",
		PaidAt:          time.Date(2024, 2, 3, 0, 53, 26, 0, time.UTC),
		CreatedAt:       time.Date(2024, 2, 3, 0, 53, 3, 0, time.UTC),
		Channel:         "card",
		Currency:        "NGN",
		IPAddress:       "105.112.28.160",
		Metadata:        "",
		Log: LogData{
			StartTime: 1706921603,
			TimeSpent: 4,
			Attempts:  1,
			Errors:    0,
			Success:   true,
			Mobile:    false,
			Input:     []string{},
			History:   []History{{Type: "action", Message: "Attempted to pay with card", Time: 3}, {Type: "success", Message: "Successfully paid with card", Time: 4}},
		},
		Fees:      40000,
		FeesSplit: nil, // Null value
		Authorization: AuthorizationData{
			AuthorizationCode: "AUTH_c1u2j76bg5",
			BIN:               "408408",
			Last4:             "4081",
			ExpMonth:          "12",
			ExpYear:           "2030",
			Channel:           "card",
			CardType:          "visa ",
			Bank:              "TEST BANK",
			CountryCode:       "NG",
			Brand:             "visa",
			Reusable:          true,
			Signature:         "SIG_Y4eB6IMiaETU1RztDqaY",
			AccountName:       nil, // Null value
		},
		Customer: CustomerData{
			ID:                       157441900,
			FirstName:                nil, // Null value
			LastName:                 nil, // Null value
			Email:                    "customer@email.com",
			CustomerCode:             "CUS_7plng9e53v77tva",
			Phone:                    nil, // Null value
			Metadata:                 nil, // Null value
			RiskAction:               "default",
			InternationalFormatPhone: nil, // Null value
		},
		Plan:               nil, // Null value
		Split:              map[string]interface{}{},
		OrderID:            nil, // Null value
		PaidAtISO:          "2024-02-03T00:53:26.000Z",
		CreatedAtISO:       "2024-02-03T00:53:03.000Z",
		RequestedAmount:    2000000,
		POSTransactionData: nil, // Null value
		Source:             nil, // Null value
		FeesBreakdown:      nil, // Null value
		TransactionDate:    "2024-02-03T00:53:03.000Z",
		PlanObject:         map[string]interface{}{},
		Subaccount:         map[string]interface{}{},
	},
}

func TestInitialize(t *testing.T) {
	p := NewPaystackClient("test-api-key")
	payload := map[string]interface{}{
		"name":   "john doe",
		"email":  "johndoe@test.com",
		"amount": 1000 * 100,
	}
	resp, err := p.Initialize(payload)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if resp == nil {
		t.Fatal("Expected a response, but got nil")
	}
	// Check that the response has the correct structure
	if resp.Status != true || resp.Message == "" {
		t.Errorf("Expected response to have 'Status' and 'Message', but got: %+v", resp)
	}

	// Check that the 'Data' field exists and has the correct type
	dataValue := reflect.ValueOf(resp).Elem().FieldByName("Data")
	if !dataValue.IsValid() {
		t.Errorf("Expected 'Data' field to exist in the response")
	}

	// Check that the 'Data' field is a struct with the correct fields
	dataStruct := dataValue.Interface().(struct {
		Reference        string `json:"reference"`
		AuthorizationURL string `json:"authorization_url"`
		AccessCode       string `json:"access_code"`
	})
	if dataStruct.Reference == "" || dataStruct.AuthorizationURL == "" || dataStruct.AccessCode == "" {
		t.Errorf("Expected 'Data' to contain 'Reference', 'AuthorizationURL', and 'AccessCode', but got: %+v", dataStruct)
	}
}

func TestVerify(t *testing.T) {
	p := NewPaystackClient("test-api-key")
	resp, err := p.Verify("0l2qk643pk")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if !reflect.DeepEqual(resp, &mockVerifyResponse) {
		t.Errorf("Expected response to be: %v, but got: %v", mockVerifyResponse, resp)
	}
}

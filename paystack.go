package paystack

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/berryboylb/go_paystack_wrapper/requests"
	"github.com/google/uuid"
	"io"
	"net/http"
	"regexp"
)

func NewPaystackClient(apiKey string) *Paystack {
	return &Paystack{APIKey: &apiKey}
}

func isValidEmail(email string) bool {
	// Regular expression pattern for email validation
	pattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
}

func (p *Paystack) Initialize(payload interface{}) (*PostResponseData, error) {
	// Assert that the payload is a map with string keys
	payloadMap, ok := payload.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expected payload to be a map with string keys, got %T", payload)
	}

	// Validate and extract 'email' field
	email, emailExists := payloadMap["email"].(string)
	if !emailExists {
		return nil, errors.New("payload must contain 'email' field")
	}
	if !isValidEmail(email) {
		return nil, errors.New("invalid email address" + email)
	}

	// Validate and extract 'amount' field
	amount, amountExists := payloadMap["amount"].(float64)
	if !amountExists {
		return nil, errors.New("payload must contain 'amount' field of type float64")
	}
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	paystackClient := requests.NewAPIClient(*p.APIKey)
	resp, err := paystackClient.Post("/transaction/initialize", payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//read data from the body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Error reading response body: " + err.Error())
	}

	//check for response
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		var failed Error
		err = json.Unmarshal(body, &failed)
		if err != nil {
			return nil, errors.New("Error decoding failed JSON: " + err.Error())
		}
		return nil, errors.New("received non-200 response" + failed.Message)
	}

	// convert to JSON object and return it
	var responseData PostResponseData
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, errors.New("Error decoding JSON: " + err.Error())
	}
	//check if response.status is not okay
	if !responseData.Status {
		return nil, errors.New("failed to initialize transaction " + responseData.Message)
	}

	//return data
	return &responseData, nil
}

func (p *Paystack) Verify(reference string) (*GetResponseData, error) {
	//initialize new request
	paystackClient := requests.NewAPIClient(*p.APIKey)
	resp, err := paystackClient.Get("/transaction/verify/" + reference)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Error reading response body: " + err.Error())
	}

	//check for response
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		var failed Error
		err = json.Unmarshal(body, &failed)
		if err != nil {
			return nil, errors.New("Error decoding failed JSON: " + err.Error())
		}
		return nil, errors.New("received non-200 response" + failed.Message)
	}

	// convert to JSON object and return it
	var responseData GetResponseData
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, errors.New("Error decoding JSON: " + err.Error())
	}
	//return data
	return &responseData, nil
}

func (p *Paystack) ListTransactions(filter ListTransactions) (*FullResponse, error) {
	//validate arguments
	err := Validate(filter)
	if err != nil {
		return nil, errors.New("Error validating  arguments: " + err.Error())
	}
	// filter fields that are empty
	filtered := filter.FilterFields()

	//encode values as params
	encodedParams, err := encodeFilteredFields(filtered)
	if err != nil {
		return nil, errors.New("Error encoding filtered data: " + err.Error())
	}

	// Construct the full URL with query parameters
	fullURL := "/transaction" + "?" + encodedParams

	//initialize new request
	paystackClient := requests.NewAPIClient(*p.APIKey)
	resp, err := paystackClient.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Error reading response body: " + err.Error())
	}

	//check for response
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		var failed Error
		err = json.Unmarshal(body, &failed)
		if err != nil {
			return nil, errors.New("Error decoding failed JSON: " + err.Error())
		}
		return nil, errors.New("received non-200 response" + failed.Message)
	}

	// convert to JSON object and return it
	var response FullResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.New("Error decoding JSON: " + err.Error())
	}
	//return data
	return &response, nil
}

func (p *Paystack) ListBanks(filter FilterBanks) (*BankResponse, error) {
	//validate arguments
	err := Validate(filter)
	if err != nil {
		return nil, errors.New("Error validating  arguments: " + err.Error())
	}
	// filter fields that are empty
	filtered := filter.FilterEmptyFields()

	//encode values as params
	encodedParams, err := encodeFilteredFields(filtered)
	if err != nil {
		return nil, errors.New("Error encoding filtered data: " + err.Error())
	}

	// Construct the full URL with query parameters
	fullURL := "/bank" + "?" + encodedParams

	//initialize new request
	paystackClient := requests.NewAPIClient(*p.APIKey)
	resp, err := paystackClient.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Error reading response body: " + err.Error())
	}
	//check for response
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		var failed Error
		err = json.Unmarshal(body, &failed)
		if err != nil {
			return nil, errors.New("Error decoding failed JSON: " + err.Error())
		}
		return nil, errors.New("received non-200 response" + failed.Message)
	}

	// convert to JSON object and return it
	var response BankResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.New("Error decoding JSON: " + err.Error())
	}
	//return data
	return &response, nil
}

func (p *Paystack) Transfer(payload TransferInput) (*InitTransferResponse, error) {
	//validate arguments
	err := Validate(payload)
	if err != nil {
		return nil, errors.New("Error validating  arguments: " + err.Error())
	}

	// Set default values for optional fields
	if payload.Currency == "" {
		payload.Currency = "NGN"
	}
	if payload.Reference == "" {
		payload.Reference = uuid.New().String()
	}

	//build payload
	requestBody := map[string]interface{}{
		"source":    "balance", //only balance is allowed for now
		"reason":    payload.Reason,
		"amount":    payload.Amount,
		"recipient": payload.Recipient,
		"currency":  payload.Currency,
		"reference": payload.Reference,
	}

	//initialize new request
	paystackClient := requests.NewAPIClient(*p.APIKey)
	resp, err := paystackClient.Post("/transfer", requestBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Error reading response body: " + err.Error())
	}

	//check for response
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		var failed Error
		err = json.Unmarshal(body, &failed)
		if err != nil {
			return nil, errors.New("Error decoding failed JSON: " + err.Error())
		}
		return nil, errors.New("received non-200 response" + failed.Message)
	}

	// convert to JSON object and return it
	var response InitTransferResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.New("Error decoding JSON: " + err.Error())
	}
	//return data
	return &response, nil
}

func (p *Paystack) ConfirmTransfer(payload ConfirmTransferInput) (*ConfirmTransferResponse, error) {
	//validate arguments
	err := Validate(payload)
	if err != nil {
		return nil, errors.New("Error validating  arguments: " + err.Error())
	}
	requestBody := map[string]interface{}{
		"transfer_code": payload.TransferCode, //only balance is allowed for now
		"otp":           payload.OTP,
	}
	//initialize new request
	paystackClient := requests.NewAPIClient(*p.APIKey)
	resp, err := paystackClient.Post("/transfer/finalize_transfer", requestBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Error reading response body: " + err.Error())
	}

	//check for response
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		var failed Error
		err = json.Unmarshal(body, &failed)
		if err != nil {
			return nil, errors.New("Error decoding failed JSON: " + err.Error())
		}
		return nil, errors.New("received non-200 response " + failed.Message)
	}

	// convert to JSON object and return it
	var response ConfirmTransferResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.New("Error decoding JSON: " + err.Error())
	}
	//return data
	return &response, nil
}

func (p *Paystack) CreateRecipient(payload AccountDetails) (*Recipient, error) {
	//validate arguments
	err := Validate(payload)
	if err != nil {
		return nil, errors.New("Error validating  arguments: " + err.Error())
	}
	//initialize new request
	paystackClient := requests.NewAPIClient(*p.APIKey)
	resp, err := paystackClient.Post("/transferrecipient", payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Error reading response body: " + err.Error())
	}
	//check for response
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		var failed Error
		err = json.Unmarshal(body, &failed)
		if err != nil {
			return nil, errors.New("Error decoding failed JSON: " + err.Error())
		}
		return nil, errors.New("received non-200 response" + failed.Message)
	}

	// convert to JSON object and return it
	var response Recipient
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.New("Error decoding JSON: " + err.Error())
	}
	//return data
	return &response, nil
}

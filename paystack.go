package paystack

import (
	"encoding/json"
	"errors"
	"github.com/berryboylb/go_paystack_wrapper/requests"
	"io"
	"net/http"
)

func NewPaystackClient(apiKey string) *Paystack {
	return &Paystack{APIKey: &apiKey}
}

func (p *Paystack) Initialize(payload interface{}) (*PostResponseData, error) {
	paystackClient := requests.NewAPIClient(*p.APIKey)
	resp, err := paystackClient.Post("/transaction/initialize", payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check the HTTP response status
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("received non-200 response")
	}

	//read data from the body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Error reading response body: " + err.Error())
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

	// Check the HTTP response status
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("received non-200 response")
	}

	// read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Error reading response body: " + err.Error())
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

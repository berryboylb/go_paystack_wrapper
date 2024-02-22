# Paystack wrapper

Paystack wrapper is a module around paystack that allows you to  interact with the Paystack API in an easy and simple way. It provides methods for making payment,

## Installation

```bash
go get -u github.com/berryboylb/go_paystack_wrapper
```

## Usage

### Basic Example

```go

package main

import (
	"fmt"
	"log"

	"github.com/berryboylb/go_paystack_wrapper" 
)

var apikey  = "your-api-key" // replace with your paystack test or live key
var payStackClient := NewPaystackClient(apikey) //initialize new paystack variable
func main() {
	
    //to initialize a transaction

    //create a custom payload
    payload := map[string]interface{}{
		"name":   "john doe",
		"email":  "johndoe@test.com",
		"amount": float64(1000 * 100),
	}

    // call the initialize method
	resp, err := payStackClient.Initialize(payload)
	if err != nil {
		//handle your error
        log.Fatal(err)
	}
    fmt.Println("Transaction initialized: ", resp)


    
    //To verify the payment status of an authorization

    //define your reference
    reference := "0l2qk643pk"

    //call the verify method
    resp, err := payStackClient.Verify(reference)
	if err != nil {
		//handle your error
        log.Fatal(err)
	}

    fmt.Println("Transaction data: ", resp)

	//list transactions
	resp, err := payStackClient.ListTransactions(ListTransactions{
		PerPage: 10,
		Page:    1,
	})
	if err != nil {
		//handle your error
        log.Fatal(err)
	}
	
    fmt.Println("Transaction data: ", resp)

	//list banks
	resp, err := payStackClient.ListBanks(FilterBanks{
		PerPage:   50,
		Country:   "nigeria",
		UseCursor: "true",
	})
	if err != nil {
		//handle your error
        log.Fatal(err)
	}
	
    fmt.Println("banks: ", resp)

	//create recipient

	resp, err := payStackClient.CreateRecipient(AccountDetails{
		Type:          "nuban",
		Name:          "OLURUNFEMI WINNER DARAMOLA",
		AccountNumber: "0087476870",
		BankCode:      "044",
		Currency:      "NGN",
		Description:   "test",
	})
	if err != nil {
		//handle your error
        log.Fatal(err)
	}
	fmt.Println("recipient: ", resp)


	//initialize transfer

	resp, err := payStackClient.Transfer(TransferInput{
		Amount:    50 * 100,
		Recipient: "RCP_c8y67uhuvl2xmws",
		Reason:    "test",
	})
	if err != nil {
		//handle your error
        log.Fatal(err)
	}
	fmt.Println("transfer: ", resp)


	//confirm transfer
	resp, err := payStackClient.ConfirmTransfer(ConfirmTransferInput{
		TransferCode: "TRF_ivi6mjnpzx2ccfbd",
		OTP:          "522839",
	})
	if err != nil {
		//handle your error
        log.Fatal(err)
	}

	fmt.Println("confirm transfer: ", resp)
}
```

## Configuration

- `APIKey`: Your paystack API key.
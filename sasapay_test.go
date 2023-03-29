package sasapay_test

import (
	"testing"

	"github.com/salticon/sasapay-golang"
	"github.com/salticon/sasapay-golang/models"
)

var clientId = "8mgx3sf4QhfZpN7aG9DIVdrrMVyTFxU89gz5gaur"
var clientSecret = "EWbIcQEhd3acV8vcAAyuldKpp2EaWNpda4GfQHuANW5biExHDLcGLuxJ6BV1UgHNODfXUUsQqwHBSlc9KINFofXQjQ7DuqI124aICYjsz5MiGn5KajTA8F1YbOQMhHtM"
var sp = sasapay.NewSasaPay(clientId, clientSecret, "600980", int(sasapay.Sandbox))

func TestSetAccessToken(t *testing.T) {
	// _ := sasapay.NewSasaPay(" ", " ", int(sasapay.Sandbox))
}
func TestC2B(t *testing.T) {
	response, err := sp.Customer2Business(models.C2BRequest{

		MerchantCode:     "600980",
		Currency:         "KES",
		NetworkCode:      "0",
		PhoneNumber:      "254703545191",
		TransactionDesc:  "desc",
		AccountReference: "ref",
		Amount:           2,
		CallBackURL:      "https://posthere.io/67df-4d9c-9386",
	})
	if err != nil {
		t.Error(err)
	}

	respCcbProcess, err := sp.C2BProces(response.CheckoutRequestID, "4345")

	if err != nil {
		t.Error(err)
	}
	t.Log(respCcbProcess.Detail)

}
func TestB2c(t *testing.T) {
	respone, err := sp.Business2Customer(models.B2CRequest{

		MerchantCode:                 sp.MerchantCode,
		MerchantTransactionReference: "dsd",
		Amount:                       1,
		Currency:                     "KES",
		ReceiverNumber:               "254703545191",
		Channel:                      "0",
		Reason:                       "test reason",
		CallBackURL:                  "https://posthere.io/67df-4d9c-9386",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(respone.Detail)
}

func TestB2B(t *testing.T) {
	response, err := sp.Business2Business(models.B2BRequest{
		MerchantCode:                 sp.MerchantCode,
		MerchantTransactionReference: "uoiwp",
		Currency:                     "KES",
		Amount:                       1,
		ReceiverMerchantCode:         "94000",
		CallBackURL:                  "https://posthere.io/67df-4d9c-9386",
		Reason:                       "test",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(response.Detail)
}

func TestCheckTransactioStatus(t *testing.T) {
	response, err := sp.CheckTransactionStatus("6e1e251f-afb0-****-a097-f1ae0e0b2ce6")
	if err != nil {
		t.Error(err)
	}
	t.Log(response.Data)
}

func TestVerifyTransaction(t *testing.T) {
	response, err := sp.VerifyTransaction("CDKQYQHMD")
	if err != nil {
		t.Error(err)
	}
	t.Log(response.Data)
}

func TestMerchantBalance(t *testing.T) {
	response, err := sp.CheckMerchantBalance()
	if err != nil {
		t.Error(err)
	}
	t.Log(response.Data)
}

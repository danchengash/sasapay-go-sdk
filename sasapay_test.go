package sasapay_test

import (
	"testing"

	"github.com/salticon/sasapay-golang"
	"github.com/salticon/sasapay-golang/models"
)

var clientId = "8mgx3sf4QhfZpN7aG9DIVdrrMVyTFxU89gz5gaur"
var clientSecret = "EWbIcQEhd3acV8vcAAyuldKpp2EaWNpda4GfQHuANW5biExHDLcGLuxJ6BV1UgHNODfXUUsQqwHBSlc9KINFofXQjQ7DuqI124aICYjsz5MiGn5KajTA8F1YbOQMhHtM"

func TestSetAccessToken(t *testing.T) {
	// _ := sasapay.NewSasaPay(" ", " ", int(sasapay.Sandbox))
}
func TestC2B(t *testing.T) {
	sasapay := sasapay.NewSasaPay(clientId, clientSecret, int(sasapay.Sandbox))
	sasapay.Customer2Business(models.C2BRequest{

		MerchantCode:     "600980",
		Currency:         "KES",
		NetworkCode:      "0",
		PhoneNumber:      "2547****191",
		TransactionDesc:  "desc",
		AccountReference: "ref",
		Amount:           2,
		CallBackURL:      "https://posthere.io/67df-4d9c-9386",
	})

}

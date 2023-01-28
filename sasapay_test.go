package sasapay_test

import (
	"testing"

	"github.com/salticon/sasapay-golang"
)

func TestSetAccessToken(t *testing.T) {
	sasapay := sasapay.NewSasaPay(" ", " ", int(sasapay.Sandbox))
	sasapay.SetAccessToken()
}

package sasapay

import (
	"fmt"

	"github.com/salticon/sasapay-golang/helpers"
)

type SasaPay struct {
	Environment  int
	ClientId     string
	ClientSecret string
}

func NewSasaPay(clientId string, clientSecret string, environment int) SasaPay {

	return SasaPay{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Environment:  environment,
	}

}

// setAccessToken returns a time bound access token to call allowed APIs.
// This token should be used in all other subsequent responses to the APIs
func (s *SasaPay) SetAccessToken() {
	res, err := helpers.NewReq("https://sasapay.co.ke/", nil, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("____+++++++++++_____")
	fmt.Println(string(res))

}


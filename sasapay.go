package sasapay

import (
	"encoding/base64"
	"fmt"

	"github.com/salticon/sasapay-golang/helpers"
	"github.com/salticon/sasapay-golang/models"
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
func (s *SasaPay) SetAccessToken() (*models.AccessTokenResponse, error) {
	url := s.baseURL() + SetAccessTokenUrl
	b := []byte(s.ClientId + ":" + s.ClientSecret)
	encoded := base64.StdEncoding.EncodeToString(b)
	headers := make(map[string]string)
	headers["Authorization"] = "Basic " + encoded
	res, err := helpers.NewReq(url, nil, &headers)
	if err != nil {
		fmt.Println(err)
		return nil, &models.RequestError{StatusCode: res.StatusCode(), Message: string(res.Body()), Url: res.LocalAddr().String()}
	}
	fmt.Println("____+++++++++++_____")
	fmt.Println(string(res.Body()))
	accToken := models.AccessTokenResponse{}
	errRes := models.RequestError{}
	if res.StatusCode() >= 200 || res.StatusCode() <= 300 {
		resBody :=  models.UnmarshalAccessTokenResponse(res.Body())

		if resBody {
			
		}
		accToken, err =
		if err != nil {
			return nil, err
		}
	}
	return &accToken, nil
}

func (s *SasaPay) baseURL() string {
	if s.Environment == int(Production) {
		return "https://api.sasapay.me/api/v1/"
	}
	return "https://api.sasapay.me/api/v1/"
}

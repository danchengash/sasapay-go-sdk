package sasapay

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/salticon/sasapay-golang/helpers"
	"github.com/salticon/sasapay-golang/models"
)

type SasaPay struct {
	Environment  int
	ClientId     string
	ClientSecret string
	cacheToken   models.AccessTokenResponse
}

func NewSasaPay(clientId string, clientSecret string, environment int) SasaPay {
	var accessToken = models.AccessTokenResponse{}

	return SasaPay{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Environment:  environment,
		cacheToken:   accessToken,
	}

}

// setAccessToken returns a time bound access token to call allowed APIs.
// This token should be used in all other subsequent responses to the APIs
func (s *SasaPay) setAccessToken() (string, error) {
	if time.Until(s.cacheToken.ExpiresAt.UTC()).Seconds() > 0 {
		return s.cacheToken.AccessToken, nil
	}
	url := s.baseURL() + SetAccessTokenUrl
	b := []byte(s.ClientId + ":" + s.ClientSecret)
	encoded := base64.StdEncoding.EncodeToString(b)
	headers := make(map[string]string)
	headers["Authorization"] = "Basic " + encoded
	res, err := helpers.NewReq(url, nil, &headers, false)
	if err != nil {
		fmt.Println(err)
		return "", &models.RequestError{StatusCode: res.StatusCode(), Message: string(res.Body()), Url: res.LocalAddr().String()}
	}
	if res.StatusCode() >= 200 && res.StatusCode() <= 300 {
		s.cacheToken, err = models.UnmarshalAccessTokenResponse(res.Body())
		s.cacheToken.ExpiresAt = time.Now().Add(time.Duration(s.cacheToken.ExpiresIn) * time.Second)
		if err != nil {
			return "", err
		}
	}
	println(s.cacheToken.AccessToken)
	return s.cacheToken.AccessToken, nil
}

func (s *SasaPay) RegisterCallBackUrl(merchantcode string, confirmationUrl string) (*models.RegisterConfirmationURLResponse, error) {
	token, err := s.setAccessToken()
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	params := make(map[string]interface{})
	params["MerchantCode"] = merchantcode
	params["ConfirmationUrl"] = confirmationUrl
	reqEntityBytes, _ := json.Marshal(params)
	headers["Authorization"] = "Bearer " + token
	url := s.baseURL() + registerCallBack
	res, err := helpers.NewReq(url, &reqEntityBytes, &headers)
	if err != nil {
		return nil, &models.RequestError{StatusCode: res.StatusCode(), Message: string(res.Body()), Url: res.LocalAddr().String()}
	}
	resp, err := models.UnmarshalRegisterConfirmationURLResponse(res.Body())
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *SasaPay)

func (s *SasaPay) baseURL() string {
	if s.Environment == int(Production) {
		return "https://api.sasapay.me/api/v1/"
	}
	return "https://api.sasapay.me/api/v1/"
}

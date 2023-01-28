package helpers

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

var contentyTypeHeaderJson = []byte("application/json")
func NewReq(url string, body *[]byte, headers *map[string]string) (string, error) {
	readTimeout, _ := time.ParseDuration("15s")
	writeTimeout, _ := time.ParseDuration("15s")
	maxIdleConnDuration, _ := time.ParseDuration("30m")
	tlsConf := &tls.Config{InsecureSkipVerify: true}
	dial := (&fasthttp.TCPDialer{Concurrency: 100, DNSCacheDuration: time.Hour}).Dial
	client := fasthttp.Client{
		Name:                          "sasapay-sdk",
		ReadTimeout:                   readTimeout,
		WriteTimeout:                  writeTimeout,
		MaxIdleConnDuration:           maxIdleConnDuration,
		NoDefaultUserAgentHeader:      true,
		DisableHeaderNamesNormalizing: true,
		DisablePathNormalizing:        true,
		TLSConfig:                     tlsConf,
		Dial:                          dial,
	}
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	req.SetRequestURI(url)
	req.Header.SetContentTypeBytes(contentyTypeHeaderJson)
	if headers != nil {
		for key, value := range *headers {
			req.Header.Set(key, value)
		}
	}

	//GET REQUEST.
	if body == nil {
		req.Header.SetMethod(fasthttp.MethodGet)

	} else if body != nil {
		//POST REQUEST
		req.Header.SetMethod(fasthttp.MethodPost)
		req.SetBodyRaw(*body)
	}

	err := client.Do(req, resp)
	if err != nil {
		fmt.Printf("<ERROR ->>: %s\n", err)

	}
	// RELEASE RESOURCES.
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(resp)

	return string(resp.Body()), nil
}

// 	if err == nil {
// 		fmt.Printf("DEBUG Response: %s\n", resp.Body())
// 	} else {
// 		fmt.Fprintf(os.Stderr, "ERR Connection error: %v\n", err)
// 	}
// 	fasthttp.ReleaseResponse(resp)

// 	reqEntity := &Entity{
// 		Name: "New entity",
// 	}
// 	reqEntityBytes, _ := json.Marshal(reqEntity)

// 	req := fasthttp.AcquireRequest()
// 	req.SetRequestURI("http://localhost:8080/")
// 	req.Header.SetMethod(fasthttp.MethodPost)
// 	req.SetBodyRaw(reqEntityBytes)
// 	resp := fasthttp.AcquireResponse()
// 	err := client.DoTimeout(req, resp, reqTimeout)
// 	fasthttp.ReleaseRequest(req)

// 	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
// 	if err != nil {
// 		return "", nil
// 	}

// 	for key, value := range headers {
// 		request.Header.Set(key, value)
// 	}

// 	client := &http.Client{Timeout: 60 * time.Second}
// 	res, err := client.Do(request)
// 	if res != nil {
// 		defer res.Body.Close()
// 	}
// 	if err != nil {
// 		return "", err
// 	}

// 	stringBody, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		return "", err
// 	}

// 	return string(stringBody), nil
// }

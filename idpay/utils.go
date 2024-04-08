package idpay

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
)

func request(
	apiKey string,
	sandBox bool,
	method string,
	url string,
	body []byte,
) (*int, []byte, error) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-KEY", apiKey)
	request.Header.Set("X-SANDBOX", strconv.FormatBool(sandBox))
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, nil, err
	}
	defer response.Body.Close()
	reponseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}
	// fmt.Printf("response status code: %d\tresponse body: %s\n", response.StatusCode, reponseBody)
	return &response.StatusCode, reponseBody, nil
}

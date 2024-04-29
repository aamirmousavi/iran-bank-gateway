package zarinpal

import (
	"bytes"
	"io"
	"net/http"
)

type method string

const (
	_PAYMENT = "payment"
	_VERIFY  = "verify"
)

func request(
	sandBox bool,
	method method,
	body []byte,
) (*int, []byte, error) {
	var url string
	switch method {
	case _PAYMENT:
		if sandBox {
			url = CREATE_SANDBOX_TRANSACTION_URL
		} else {
			url = CREATE_TRANSACTION_URL
		}
	case _VERIFY:
		if sandBox {
			url = VERIFY_SANDBOX_TRANSACTION_URL
		} else {
			url = VERIFY_TRANSACTION_URL
		}
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, nil, err
	}
	request.Header.Set("Content-Type", "application/json")
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

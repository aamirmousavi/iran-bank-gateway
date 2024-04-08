package idpay

import (
	"encoding/json"
	"fmt"
)

func (hand *idpay) Payment(params *PaymentRequest) (*int, *PaymentResponse, *ErrorResponse, error) {
	apiKey, sandBox, err := hand.GetIdentityData()
	if err != nil {
		return nil, nil, nil, err
	}
	payload, err := params.raw()
	if err != nil {
		return nil, nil, nil, err
	}
	statusCode, response, err := request(
		apiKey, sandBox,
		"POST",
		CREATE_TRANSACTION_URL,
		payload,
	)
	if err != nil {
		return nil, nil, nil, err
	}
	// if status code is not 2XX we expect an error response
	if *statusCode < 200 || *statusCode > 299 {
		responseError := new(ErrorResponse)
		if err = json.Unmarshal(response, &responseError); err != nil {
			return statusCode, nil, nil, err
		}
		return statusCode, nil, responseError, nil
	}
	result := new(PaymentResponse)
	if err = json.Unmarshal(response, &result); err != nil {
		return statusCode, nil, nil, fmt.Errorf("error in unmarshal response: %w", err)
	}
	return statusCode, result, nil, nil
}

func (hand *idpay) Verify(params *VerifyRequest) (*int, *VerifyResponse, *ErrorResponse, error) {
	apiKey, sandBox, err := hand.GetIdentityData()
	if err != nil {
		return nil, nil, nil, err
	}
	payload, err := params.raw()
	if err != nil {
		return nil, nil, nil, err
	}
	statusCode, response, err := request(
		apiKey, sandBox,
		"POST",
		VERIFY_TRANSACTION_URL,
		payload,
	)
	if err != nil {
		return nil, nil, nil, err
	}
	// if status code is not 2XX we expect an error response
	if *statusCode < 200 || *statusCode > 299 {
		responseError := new(ErrorResponse)
		if err = json.Unmarshal(response, &responseError); err != nil {
			return statusCode, nil, nil, err
		}
		responseError.RawResponse = response
		return statusCode, nil, responseError, nil
	}
	result := new(VerifyResponse)
	result.RawResponse = response
	if err = json.Unmarshal(response, &result); err != nil {
		return statusCode, nil, nil, fmt.Errorf("error in unmarshal response: %w", err)
	}
	return statusCode, result, nil, nil
}

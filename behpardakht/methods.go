package behpardakht

import (
	"encoding/xml"
	"fmt"
)

func (hand *behpardakht) Payment(params *paymentRequest) (
	*int,
	*PaymentResponse,
	error,
) {
	userId, password, err := hand.GetIdentityData()
	if err != nil {
		return nil, nil, err
	}
	payload, err := params.raw(userId, password)
	if err != nil {
		return nil, nil, err
	}

	statusCode, response, err := request("POST", CREATE_TRANSACTION_URL, payload)
	if err != nil {
		return nil, nil, err
	}
	// if status code is not 2XX we expect an error response
	if *statusCode < 200 || *statusCode > 299 {
		responseError := new(PaymentResponse)
		if err = xml.Unmarshal(response, &responseError); err != nil {
			return nil, nil, fmt.Errorf("error raw response: %v", string(response))
		}
		return statusCode, responseError, nil
	}
	result := new(PaymentResponse)
	if err = xml.Unmarshal(response, &result); err != nil {
		return statusCode, nil, fmt.Errorf("error in unmarshal response: %w", err)
	}
	result.modifyResponse()
	return statusCode, result, nil
}

func (hand *behpardakht) Verify(params *verifyRequest) (
	*int,
	*VerifyResponse,
	error,
) {
	userId, password, err := hand.GetIdentityData()
	if err != nil {
		return nil, nil, err
	}
	payload, err := params.raw(userId, password)
	if err != nil {
		return nil, nil, err
	}

	statusCode, response, err := request("POST", VERIFY_TRANSACTION_URL, payload)
	if err != nil {
		return nil, nil, err
	}
	// if status code is not 2XX we expect an error response
	if *statusCode < 200 || *statusCode > 299 {
		responseError := new(VerifyResponse)
		if err = xml.Unmarshal(response, &responseError); err != nil {
			return nil, nil, fmt.Errorf("error raw response: %v", string(response))
		}
		return statusCode, responseError, nil
	}
	result := new(VerifyResponse)
	result.rawResponse = response
	if err = xml.Unmarshal(response, &result); err != nil {
		return statusCode, nil, fmt.Errorf("error in unmarshal response: %w", err)
	}
	result.modifyResponse()
	return statusCode, result, nil
}

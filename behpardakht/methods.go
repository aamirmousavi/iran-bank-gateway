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
	result := new(paymentResponse)
	result.rawResponse = response
	if err = xml.Unmarshal(response, &result); err != nil {
		return statusCode, nil, fmt.Errorf("error in unmarshal response: %w", err)
	}
	jsonResult, err := result.intoJson()
	if err != nil {
		return statusCode, nil, err
	}
	return statusCode, jsonResult, nil
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
	result := new(verifyResponse)
	result.rawResponse = response
	if err = xml.Unmarshal(response, &result); err != nil {
		return statusCode, nil, fmt.Errorf("error in unmarshal response: %w", err)
	}
	jsonResult, err := result.intoJson()
	if err != nil {
		return statusCode, nil, err
	}
	return statusCode, jsonResult, nil
}

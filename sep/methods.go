package sep

import "encoding/json"

func (hand *sep) Payment(params *paymentRequest) (
	*int,
	*PaymentResponse,
	*ErrorResponse,
	error,
) {
	_, terminalCode, err := hand.GetIdentityData()
	if err != nil {
		return nil, nil, nil, err
	}
	payload, err := params.raw(terminalCode)
	if err != nil {
		return nil, nil, nil, err
	}
	statusCode, response, err := request(
		CREATE_TRANSACTION_URL,
		payload,
	)
	if err != nil {
		return nil, nil, nil, err
	}
	// if status code is not 2XX we expect an error response
	if *statusCode < 200 || *statusCode > 299 {
		responseError := new(ErrorResponse)
		if err := json.Unmarshal(response, &responseError); err != nil {
			return statusCode, nil, nil, err
		}
		return statusCode, nil, responseError, nil
	}
	result := new(PaymentResponse)
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, nil, nil, err
	}
	return statusCode, result, nil, nil
}

func (hand *sep) Verify(params *verifyRequest) (
	*int,
	*VerifyResponse,
	error,
) {
	_, terminalCode, err := hand.GetIdentityData()
	if err != nil {
		return nil, nil, err
	}
	payload, err := params.raw(terminalCode)
	if err != nil {
		return nil, nil, err
	}
	statusCode, response, err := request(
		VERIFY_TRANSACTION_URL,
		payload,
	)
	if err != nil {
		return nil, nil, err
	}
	result := new(VerifyResponse)
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, nil, err
	}
	return statusCode, result, nil
}

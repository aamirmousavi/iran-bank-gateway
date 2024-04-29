package zarinpal

import "encoding/json"

func (hand *zarinpal) Payment(params *PaymentRequest) (*int, *PaymentResponse, error) {
	apiKey, sandBox, err := hand.GetIdentityData()
	if err != nil {
		return nil, nil, err
	}
	payload, err := params.raw(apiKey)
	if err != nil {
		return nil, nil, err
	}
	statusCode, response, err := request(
		sandBox,
		_PAYMENT,
		payload,
	)
	if err != nil {
		return nil, nil, err
	}
	result := new(PaymentResponse)
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, nil, err
	}
	return statusCode, result, nil
}

func (hand *zarinpal) Verify(params *VerifyRequest) (*int, *VerifyResponse, error) {
	apiKey, sandBox, err := hand.GetIdentityData()
	if err != nil {
		return nil, nil, err
	}
	payload, err := params.raw(apiKey)
	if err != nil {
		return nil, nil, err
	}
	statusCode, response, err := request(
		sandBox,
		_VERIFY,
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

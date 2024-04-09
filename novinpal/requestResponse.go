package novinpal

import "encoding/json"

type ErrorResponse struct {
	Status           int    `json:"status"`
	ErrorCode        int    `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
}

type PaymentRequest struct {
	ApiKey      string  `json:"api_key"`
	Amount      uint64  `json:"amount"`
	ReturnUrl   string  `json:"return_url"`
	OrderId     string  `json:"order_id"`
	Description *string `json:"description,omitempty"`
	Mobile      *string `json:"mobile,omitempty"`
	CardNumber  *string `json:"card_number,omitempty"`
}

func NewPaymentRequest(
	amount uint64,
	returnUrl, orderId string,
	description, mobile, cardNumber *string,
) *PaymentRequest {
	return &PaymentRequest{
		Amount:      amount,
		ReturnUrl:   returnUrl,
		OrderId:     orderId,
		Description: description,
		Mobile:      mobile,
		CardNumber:  cardNumber,
	}
}

func (pr *PaymentRequest) raw(
	apiKey string,
) ([]byte, error) {
	pr.ApiKey = apiKey
	return json.Marshal(pr)
}

type PaymentResponse struct {
	RefId  string `json:"refId"`
	Status int    `json:"status"`
}

type VerifyRequest struct {
	ApiKey string `json:"api_key"`
	RefId  string `json:"ref_id"`
}

func NewVerifyRequest(
	refId string,
) *VerifyRequest {
	return &VerifyRequest{
		RefId: refId,
	}
}

func (vr *VerifyRequest) raw(
	apiKey string,
) ([]byte, error) {
	vr.ApiKey = apiKey
	return json.Marshal(vr)
}

type VerifyResponse struct {
	PaidAt         string `json:"paidAt"`
	CardNumber     string `json:"cardNumber"`
	Status         int    `json:"status"`
	Amount         uint64 `json:"amount"`
	RefNumber      string `json:"refNumber"`
	RefId          string `json:"refId"`
	Description    string `json:"description"`
	OrderId        string `json:"orderId"`
	VerifiedBefore bool   `json:"verifiedBefore"`
}

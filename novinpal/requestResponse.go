package novinpal

import (
	"fmt"
	"net/url"
)

type ErrorResponse struct {
	Status           int    `json:"status"`
	ErrorCode        string `json:"errorCode"`
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
	formData := url.Values{}
	formData.Add("api_key", apiKey)
	formData.Add("amount", fmt.Sprint(pr.Amount))
	formData.Add("return_url", pr.ReturnUrl)
	formData.Add("order_id", pr.OrderId)
	if pr.Description != nil {
		formData.Add("description", *pr.Description)
	}
	if pr.Mobile != nil {
		formData.Add("mobile", *pr.Mobile)
	}
	if pr.CardNumber != nil {
		formData.Add("card_number", *pr.CardNumber)
	}
	return []byte(formData.Encode()), nil
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
	formData := url.Values{}
	formData.Add("api_key", apiKey)
	formData.Add("ref_id", vr.RefId)
	return []byte(formData.Encode()), nil
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

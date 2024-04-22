package novinpal

import (
	"bytes"
	"fmt"
	"mime/multipart"
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
) (*bytes.Buffer, string, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	writer.WriteField("api_key", apiKey)
	writer.WriteField("amount", fmt.Sprint(pr.Amount))
	writer.WriteField("return_url", pr.ReturnUrl)
	writer.WriteField("order_id", pr.OrderId)
	if pr.Description != nil {
		writer.WriteField("description", *pr.Description)
	}
	if pr.Mobile != nil {
		writer.WriteField("mobile", *pr.Mobile)
	}
	if pr.CardNumber != nil {
		writer.WriteField("card_number", *pr.CardNumber)
	}
	if err := writer.Close(); err != nil {
		return nil, "", err
	}
	return payload, writer.FormDataContentType(), nil
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
) (*bytes.Buffer, string, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	writer.WriteField("api_key", apiKey)
	writer.WriteField("ref_id", vr.RefId)
	if err := writer.Close(); err != nil {
		return nil, "", err
	}
	return payload, writer.FormDataContentType(), nil
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

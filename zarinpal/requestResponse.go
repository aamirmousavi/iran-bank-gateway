package zarinpal

import "encoding/json"

type PaymentRequest struct {
	MerchantID  string  `json:"MerchantID"`
	Amount      uint64  `json:"Amount"`
	CallbackURL string  `json:"CallbackURL"`
	Description string  `json:"Description"`
	Email       *string `json:"Email,omitempty"`
	Mobile      *string `json:"Mobile,omitempty"`
}

func NewPaymentRequest(
	amount uint64,
	callbackUrl string,
	description string,
	email *string,
	mobile *string,
) *PaymentRequest {
	return &PaymentRequest{
		Amount:      amount,
		CallbackURL: callbackUrl,
		Description: description,
		Email:       email,
		Mobile:      mobile,
	}
}

func (pr *PaymentRequest) raw(merchantID string) ([]byte, error) {
	pr.MerchantID = merchantID
	return json.Marshal(pr)
}

type PaymentResponse struct {
	Status    int    `json:"Status"`
	Authority string `json:"Authority"`
}

type VerifyRequest struct {
	MerchantID string `json:"MerchantID"`
	Authority  string `json:"Authority"`
	Amount     uint64 `json:"Amount"`
}

func NewVerifyRequest(
	authority string,
	amount uint64,
) *VerifyRequest {
	return &VerifyRequest{
		Authority: authority,
		Amount:    amount,
	}
}

func (vr *VerifyRequest) raw(merchantID string) ([]byte, error) {
	vr.MerchantID = merchantID
	return json.Marshal(vr)
}

type VerifyResponse struct {
	Status int `json:"Status"`
	RefID  int `json:"RefID"`
}

package zarinpal

import "encoding/json"

type PaymentRequest struct {
	MerchantID  string  `json:"merchant_id"`
	Amount      uint64  `json:"amount"`
	CallbackURL string  `json:"callback_url"`
	Description string  `json:"description"`
	Email       *string `json:"email,omitempty"`
	Mobile      *string `json:"mobile,omitempty"`
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
	Status    int    `json:"status"`
	Authority string `json:"authority"`
}

type VerifyRequest struct {
	MerchantID string `json:"merchant_id"`
	Authority  string `json:"authority"`
	Amount     uint64 `json:"amount"`
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
	Status int `json:"status"`
	RefID  int `json:"ref_id"`
}

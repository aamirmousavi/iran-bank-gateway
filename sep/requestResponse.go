package sep

import "encoding/json"

type ErrorResponse struct {
	Status    int    `json:"status"`
	ErrorCode string `json:"errorCode"`
	ErrorDesc string `json:"errorDesc"`
}

type PaymentResponse struct {
	Status int    `json:"status"`
	Token  string `json:"token"`
}

func (pr *PaymentResponse) Success() bool {
	return pr.Status == 1
}

func NewPaymentRequest(
	amount uint64,
	callbackURL string,
	resNum string,
	callNumber *string,
) *paymentRequest {
	return &paymentRequest{}
}

type paymentRequest struct {
	TerminalId  string  `json:"terminalId"`
	Action      string  `json:"action"`
	Amount      uint64  `json:"amount"`
	CallbackURL string  `json:"callbackUrl"`
	ResNum      string  `json:"resNum"`
	CallNumber  *string `json:"callNumber,omitempty"`
}

func (pr *paymentRequest) raw(
	terminalCode string,
) ([]byte, error) {
	pr.TerminalId = terminalCode
	pr.Action = "Token"
	return json.Marshal(pr)
}

type VerifyResponse struct {
	ResultCode        int    `json:"ResultCode"`
	ResultDescription string `json:"ResultDescription"`
	Success           bool   `json:"Success"`
}

func NewVerifyRequest(
	refNum string,
) *verifyRequest {
	return &verifyRequest{
		RefNum: refNum,
	}
}

type verifyRequest struct {
	RefNum         string `json:"RefNum"`
	TerminalNumber string `json:"TerminalNumber"`
}

func (vr *verifyRequest) raw(
	terminalCode string,
) ([]byte, error) {
	vr.TerminalNumber = terminalCode
	return json.Marshal(vr)
}

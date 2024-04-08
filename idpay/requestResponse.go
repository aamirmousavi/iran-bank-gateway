package idpay

import (
	"encoding/json"
	"strconv"
)

type ErrorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	RawResponse  []byte `json:"-"`
}

type PaymentRequest struct {
	OrderId     string  `json:"order_id"`
	Amount      uint64  `json:"amount"`
	Name        string  `json:"name"`
	Phone       *string `json:"phone"`
	Mail        *string `json:"mail"`
	Description *string `json:"desc"`
	Callback    string  `json:"callback"`
}

func NewPaymentRequest(
	orderId string,
	amount uint64,
	name string,
	phone, mail, description *string,
	callback string,
) *PaymentRequest {
	return &PaymentRequest{
		OrderId:     orderId,
		Amount:      amount,
		Name:        name,
		Phone:       phone,
		Mail:        mail,
		Description: description,
		Callback:    callback,
	}
}

func (ct *PaymentRequest) raw() ([]byte, error) {
	b, err := json.Marshal(ct)
	if err != nil {
		return nil, err
	}
	return b, nil
}

type PaymentResponse struct {
	Id   string `json:"id"`
	Link string `json:"link"`
}

type VerifyRequest struct {
	Id      string `json:"id"`
	OrderId string `json:"order_id"`
}

func NewVerifyRequest(
	id, orderId string,
) *VerifyRequest {
	return &VerifyRequest{
		Id:      id,
		OrderId: orderId,
	}
}

func (vt *VerifyRequest) raw() ([]byte, error) {
	b, err := json.Marshal(vt)
	if err != nil {
		return nil, err
	}
	return b, nil
}

type VerifyResponse struct {
	Status  int    `json:"status"`
	TrackId string `json:"track_id"`
	Id      string `json:"id"`
	OrderId string `json:"order_id"`
	// TODO: check again the amount value when you are calling it
	Amount      string  `json:"amount"`
	Date        string  `json:"date"`
	Payment     Payment `json:"payment"`
	Verify      Verify  `json:"verify"`
	RawResponse []byte  `json:"-"`
}

func (vr *VerifyResponse) GetStatusString() string {
	return strconv.Itoa(vr.Status)
}

func (vr *VerifyResponse) Raw() ([]byte, error) {
	b, err := json.Marshal(vr)
	if err != nil {
		return nil, err
	}
	return b, nil
}

type Payment struct {
	TrackId    string      `json:"track_id"`
	Amount     string      `json:"amount"`
	CardNumber string      `json:"card_no"`
	CardHash   string      `json:"card_hash_no"`
	Date       interface{} `json:"date"`
}

type Verify struct {
	Date interface{} `json:"date"`
}

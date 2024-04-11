package behpardakht

import (
	"encoding/xml"
	"strconv"
	"strings"
	"time"
)

type paymentRequest struct {
	XMLName      xml.Name `xml:"ns1:bpPayRequest"`
	TerminalId   string   `xml:"terminalId"`
	UserName     string   `xml:"userName"`
	UserPassword string   `xml:"userPassword"`
	OrderId      string   `xml:"orderId"`
	Amount       uint64   `xml:"amount"`
	LocalDate    string   `xml:"localDate"`
	LocalTime    string   `xml:"localTime"`
	CallBackUrl  string   `xml:"callBackUrl"`
	PayerId      string   `xml:"payerId"`
}

func NewPaymentRequest(
	orderId string,
	amount uint64,
	callBackUrl string,
	payerId string,
) *paymentRequest {
	return &paymentRequest{
		OrderId:     orderId,
		Amount:      amount,
		LocalDate:   time.Now().Format("20060402"),
		LocalTime:   time.Now().Format("150405"),
		CallBackUrl: callBackUrl,
		PayerId:     payerId,
	}
}

func (pr *paymentRequest) raw(
	userId string,
	password string,
) ([]byte, error) {
	pr.TerminalId = userId
	pr.UserName = userId
	pr.UserPassword = password
	root := newSoapRoot()
	root.Body.Request = pr
	return root.Marshal()
}

type paymentResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		XMLName xml.Name `xml:"Body"`
		BpPay   struct {
			XMLName xml.Name `xml:"bpPayRequestResponse"`
			Return  string   `xml:"return"`
		}
	}
	responseCode int    `xml:"-"`
	refId        string `xml:"-"`
	rawResponse  []byte `xml:"-"`
}

type PaymentResponse struct {
	ResponseCode int    `json:"response_code"`
	RefId        string `json:"ref_id"`
	RawResponse  string `json:"raw_response"`
}

func (pr *paymentResponse) intoJson() (*PaymentResponse, error) {
	if err := pr.modifyResponse(); err != nil {
		return nil, err
	}
	return &PaymentResponse{
		ResponseCode: pr.responseCode,
		RefId:        pr.refId,
		RawResponse:  string(pr.rawResponse),
	}, nil
}

func (pr *paymentResponse) modifyResponse() error {
	params := strings.Split(pr.Body.BpPay.Return, ",")
	if len(params) > 0 {
		if params[0] == "0" {
			pr.responseCode = -1
		} else if params[0] == "" {
			return nil
		} else {
			code, err := strconv.Atoi(params[0])
			if err != nil {
				return err
			}
			pr.responseCode = code
		}
		if len(params) > 1 {
			pr.refId = params[1]
		}
	}
	return nil
}

type verifyRequest struct {
	XMLName         xml.Name `xml:"ns1:bpVerifyRequest"`
	TerminalId      string   `xml:"terminalId"`
	UserName        string   `xml:"userName"`
	Password        string   `xml:"userPassword"`
	OrderId         string   `xml:"orderId"`
	SaleOrderId     string   `xml:"saleOrderId"`
	SaleReferenceId string   `xml:"saleReferenceId"`
}

func NewVerifyRequest(
	orderId string,
	saleOrderId string,
	saleReferenceId string,
) *verifyRequest {
	return &verifyRequest{
		OrderId:         orderId,
		SaleOrderId:     saleOrderId,
		SaleReferenceId: saleReferenceId,
	}
}

func (vr *verifyRequest) raw(
	userId string,
	password string,
) ([]byte, error) {
	vr.TerminalId = userId
	vr.UserName = userId
	vr.Password = password
	root := newSoapRoot()
	root.Body.Request = vr
	return root.Marshal()
}

type verifyResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		XMLName xml.Name `xml:"Body"`
		BpPay   struct {
			XMLName xml.Name `xml:"bpVerifyRequestResponse"`
			Return  string   `xml:"return"`
		}
	}
	responseCode int    `xml:"-"`
	rawResponse  []byte `xml:"-"`
}

func (vr *verifyResponse) intoJson() (*VerifyResponse, error) {
	if err := vr.modifyResponse(); err != nil {
		return nil, err
	}

	return &VerifyResponse{
		ResponseCode: vr.responseCode,
		RawResponse:  string(vr.rawResponse),
	}, nil
}

type VerifyResponse struct {
	ResponseCode int    `json:"response_code"`
	RawResponse  string `json:"raw_response"`
}

func (vr *verifyResponse) modifyResponse() error {
	if vr.Body.BpPay.Return == "0" {
		vr.responseCode = -1
		return nil
	}
	code, err := strconv.Atoi(vr.Body.BpPay.Return)
	if err != nil {
		return err
	}
	vr.responseCode = code
	return nil
}

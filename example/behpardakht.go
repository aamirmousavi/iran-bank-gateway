package example

import (
	"log"

	"github.com/aamirmousavi/iran-bank-gateway/behpardakht"
)

type behpardakhtUserPass struct {
	username, password string
}

func (beh *behpardakhtUserPass) GetIdentityData() (string, string, error) {
	return beh.username, beh.password, nil
}

func Behpardakh() {
	userPass := &behpardakhtUserPass{
		username: "", // behpardakh username
		password: "", // behpardakht password
	}
	behpardakhtInstance := behpardakht.New(userPass)

	// payment (Request for terminal)
	statusCode, paymentResponse, err := behpardakhtInstance.Payment(
		behpardakht.NewPaymentRequest(
			"1",         // order id
			1_000_000,   //amount
			"/callback", // callback
			"1",         // payer id (user id)
		),
	)
	if err != nil {
		log.Fatalf("err = %v\n", err)
	}
	if paymentResponse.ResponseCode() != behpardakht.Success {
		log.Fatalf("response = %#v\n", paymentResponse)
	}
	log.Printf("payment status code is = %v\n", statusCode)

	// verify (verify a transaction)
	statusCode, verifyResponse, err := behpardakhtInstance.Verify(
		behpardakht.NewVerifyRequest(
			"1",               // order id
			"SaleOrderId",     // SaleOrderId PostForm parameter in callback
			"SaleReferenceId", // SaleReferenceId PostForm parameter in callback
		),
	)
	if err != nil {
		log.Fatalf("err = %v\n", err)
	}
	if verifyResponse.ResponseCode() != behpardakht.Success {
		log.Fatalf("response = %#v\n", verifyResponse)
	}
	log.Printf("verify status code is = %v\n", statusCode)

}

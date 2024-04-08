package example

import (
	"fmt"
	"log"

	"github.com/aamirmousavi/iran-bank-gateway/idpay"
)

type idpayIdentity struct {
	apiKey  string
	sandbox bool
}

func (iden *idpayIdentity) GetIdentityData() (string, bool, error) {
	return iden.apiKey, iden.sandbox, nil
}
func IdPay() {
	identity := &idpayIdentity{
		apiKey:  "",    // idpay api-key
		sandbox: false, // idpay sandbox
	}
	idpayInstance := idpay.New(identity)

	// payment (Request for terminal)
	statusCode, paymentResponse, idpayErr, err := idpayInstance.Payment(
		idpay.NewPaymentRequest(
			"1",               // order id
			uint64(1_000_000), // amount
			"amir mousavi",    //name
			nil,               // phone
			nil,               // email
			nil,               //description
			"/callback",       // callback
		),
	)
	if idpayErr != nil || err != nil {
		log.Fatalf("idpay err = %v\t err = %v\n", idpayErr, err)
	}
	fmt.Printf("status code is = %v\npayment response is = %#v\n", statusCode, paymentResponse)

	// verify (verify a transaction)
	statusCode, verifyResponse, idpayErr, err := idpayInstance.Verify(
		idpay.NewVerifyRequest(
			paymentResponse.Id, // id
			"1",                //order id
		),
	)
	if idpayErr != nil || err != nil {
		log.Fatalf("idpay err = %v\t err = %v\n", idpayErr, err)
	}
	fmt.Printf("status code is = %v\tverify response is = %#v\n", statusCode, verifyResponse)

}

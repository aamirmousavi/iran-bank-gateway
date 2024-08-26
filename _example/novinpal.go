package example

import (
	"fmt"
	"log"

	"github.com/aamirmousavi/iran-bank-gateway/novinpal"
)

type novinpalIdentity struct {
	apiKey string
}

func (iden *novinpalIdentity) GetIdentityData() (string, error) {
	return iden.apiKey, nil
}

func Novinpal() {
	identity := &novinpalIdentity{
		apiKey: "", // novinpal api-key
	}
	novinpalInstance := novinpal.New(identity)

	// payment (Request for terminal)
	statusCode, paymentResponse, novipalErr, err := novinpalInstance.Payment(
		novinpal.NewPaymentRequest(
			uint64(1_000_000), // amount
			"/callback",       // callback
			"1",               // order id
			nil,               // description
			nil,               // mobile
			nil,               // card number
		),
	)
	if novipalErr != nil || err != nil {
		log.Fatalf("novinpal err = %#v\t err = %v\n", novipalErr, err)
	}
	fmt.Printf("status code is = %v\n", statusCode)
	fmt.Printf("payment response is = %#v\n", paymentResponse)

	// verify (verify a transaction)
	statusCode, verifyResponse, novipalErr, err := novinpalInstance.Verify(
		novinpal.NewVerifyRequest(
			fmt.Sprint(paymentResponse.RefId), // ref id
		),
	)
	if novipalErr != nil || err != nil {
		log.Fatalf("novinpal err = %v\t err = %v\n", novipalErr, err)
	}
	fmt.Printf("status code is = %v\n", statusCode)
	fmt.Printf("verify response is = %#v\n", verifyResponse)
}

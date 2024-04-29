package example

import (
	"fmt"
	"log"

	"github.com/aamirmousavi/iran-bank-gateway/zarinpal"
)

type zarinpalIdentity struct {
	merchantID string
	sandbox    bool
}

func (iden *zarinpalIdentity) GetIdentityData() (string, bool, error) {
	return iden.merchantID, iden.sandbox, nil
}

func Zarinpal() {
	identity := &zarinpalIdentity{
		merchantID: "", // zarinpal MerchantID
		sandbox:    false,
	}
	zarinpalInstance := zarinpal.New(identity)

	// payment (Request for terminal)
	statusCode, paymentResponse, err := zarinpalInstance.Payment(
		zarinpal.NewPaymentRequest(
			uint64(1_000_000), // amount
			"/callback",       // callback
			"the description", // description
			nil,               // email
			nil,               // mobile
		),
	)
	if err != nil {
		log.Fatalf("zarinpal err = %#v\n", err)
	}
	fmt.Printf("status code is = %v\npayment response is = %#v\n", statusCode, paymentResponse)

	// verify (verify a transaction)
	statusCode, verifyResponse, err := zarinpalInstance.Verify(
		zarinpal.NewVerifyRequest(
			paymentResponse.Authority, // authority
			uint64(1_000_000),         // amount
		),
	)
	if err != nil {
		log.Fatalf("zarinpal err = %#v\n", err)
	}
	fmt.Printf("status code is = %v\nverify response is = %#v\n", statusCode, verifyResponse)
}

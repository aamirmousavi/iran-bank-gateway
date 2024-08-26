package example

import (
	"fmt"
	"log"

	"github.com/aamirmousavi/iran-bank-gateway/sep"
)

type setIdentity struct {
	merchantID string
	terminalID string
}

func (iden *setIdentity) GetIdentityData() (string, string, error) {
	return iden.merchantID, iden.terminalID, nil
}

func Sep() {
	identity := &setIdentity{
		merchantID: "", // sep merchant-id
		terminalID: "", // sep terminal-id
	}
	sepInstance := sep.New(identity)

	// payment (Request for terminal)
	statusCode, paymentResponse, sepErr, err := sepInstance.Payment(
		sep.NewPaymentRequest(
			uint64(1_000_000), // amount
			"/callback",       // callback
			"1",               // order id
			nil,               // mobile
		),
	)
	if sepErr != nil || err != nil {
		log.Fatalf("sep err = %#v\t err = %v\n", sepErr, err)
	}
	fmt.Printf("status code is = %v\n", statusCode)
	fmt.Printf("payment response is = %#v\n", paymentResponse)

	// verify (verify a transaction)
	statusCode, verifyResponse, err := sepInstance.Verify(
		sep.NewVerifyRequest(
			"RefId", // ref id
		),
	)
	if err != nil {
		log.Fatalf("err = %v\n", err)
	}
	fmt.Printf("status code is = %v\n", statusCode)
	fmt.Printf("verify response is = %#v\n", verifyResponse)

}

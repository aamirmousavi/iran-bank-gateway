# Iran Bank Gateway
## Contents
- [Gateways](#gateways)
    - [Behpardakht Mellat](#behpardakht-mellat)
	- [Idpay](#idpay)
## Gateways

### Behpardakht
```go
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

```
### Idpay
```go
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

```


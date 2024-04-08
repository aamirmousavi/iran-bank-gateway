package idpay

type idpay struct {
	IdentityData
}

func New(
	identityData IdentityData,
) *idpay {
	return &idpay{
		identityData,
	}
}

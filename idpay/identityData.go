package idpay

type IdentityData interface {
	GetIdentityData() (string, bool, error)
}

package zarinpal

type IdentityData interface {
	GetIdentityData() (string, bool, error)
}

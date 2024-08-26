package sep

type IdentityData interface {
	GetIdentityData() (string, string, error)
}

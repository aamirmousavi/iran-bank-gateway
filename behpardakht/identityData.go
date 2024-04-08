package behpardakht

type IdentityData interface {
	GetIdentityData() (string, string, error)
}

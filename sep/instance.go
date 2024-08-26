package sep

type sep struct {
	IdentityData
}

func New(
	identityData IdentityData,
) *sep {
	return &sep{
		identityData,
	}
}

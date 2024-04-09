package novinpal

type novinpal struct {
	IdentityData
}

func New(
	identityData IdentityData,
) *novinpal {
	return &novinpal{
		identityData,
	}
}

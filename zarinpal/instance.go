package zarinpal

type zarinpal struct {
	IdentityData
}

func New(
	identityData IdentityData,
) *zarinpal {
	return &zarinpal{
		identityData,
	}
}

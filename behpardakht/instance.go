package behpardakht

type behpardakht struct {
	IdentityData
}

func New(
	identityData IdentityData,
) *behpardakht {
	return &behpardakht{
		identityData,
	}
}

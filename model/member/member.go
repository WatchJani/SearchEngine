package member

type Member struct {
	Name           string   `json:"name"`
	Age            uint     `json:"age"`
	SecretIdentity string   `json:"secretIdentity"`
	Powers         []string `json:"powers"`
}

func New(
	Name,
	SecretIdentity string,
	Age uint,
	Powers []string,
) *Member {
	return &Member{
		Name,
		Age,
		SecretIdentity,
		Powers,
	}
}

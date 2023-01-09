package squad

import "root/model/member"

type Squad struct {
	SquadName  string          `json:"squadName"`
	HomeTown   string          `json:"homeTown"`
	Formed     uint16          `json:"formed"`
	SecretBase string          `json:"secretBase"`
	Active     bool            `json:"active"`
	Members    []member.Member `json:"members"`
}

func New(
	SquadName,
	HomeTown,
	SecretBase string,
	Formed uint16,
	Active bool,
	Members []member.Member,
) *Squad {
	return &Squad{
		SquadName,
		HomeTown,
		Formed,
		SecretBase,
		Active,
		Members,
	}
}

func NewEmpty() *Squad {
	return &Squad{}
}

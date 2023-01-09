package squad

import (
	"root/model/member"
)

func (s *Squad) FindByName(find string) *member.Member {
	alphabet := make(map[byte]int)
	alphabet['E'] = 0
	alphabet['M'] = 1
	alphabet['N'] = 3
	alphabet['['] = len(s.Members)

	//fmt.Println(alphabet)

	for myFirstIndex := alphabet[find[0]]; myFirstIndex < alphabet[find[0]+1]; myFirstIndex++ {
		if byte(find[0]) == s.Members[myFirstIndex].Name[0] {
			return &s.Members[myFirstIndex]
		}
	}

	return nil
}

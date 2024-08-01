package cap

import (
	"strings"

	"github.com/pkg/errors"
)

type MemberType int

const (
	Senior MemberType = iota
	Cadet
	CadetSponsor
	AEM
	Legislative
	Patron
)

func ParseMemberType(s string) (MemberType, error) {
	switch strings.ToLower(s) {
	case "senior":
		return Senior, nil
	case "cadet":
		return Cadet, nil
	case "cadet sponsor":
		return CadetSponsor, nil
	case "aem":
		return AEM, nil
	case "legislative":
		return Legislative, nil
	case "patron":
		return Patron, nil
	default:
		return -1, errors.Errorf("unrecognized member type: %s", s)
	}
}

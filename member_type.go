package cap

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type MemberType int

const (
	Senior MemberType = iota
	Cadet
	CadetSponsor
	AEM
	StateLegislative
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
	case "state leg":
		return StateLegislative, nil
	case "legislative":
		return Legislative, nil
	case "patron":
		return Patron, nil
	default:
		return -1, errors.Errorf("unrecognized member type: %s", s)
	}
}

func (m MemberType) String() string {
	switch m {
	case Senior:
		return "SENIOR"
	case Cadet:
		return "CADET"
	case CadetSponsor:
		return "CADET SPONSOR"
	case AEM:
		return "AEM"
	case StateLegislative:
		return "STATE LEG"
	case Legislative:
		return "LEGISLATIVE"
	case Patron:
		return "PATRON"
	}

	panic(errors.Errorf("invalid MemberType value: %d", m))
}

func (m MemberType) MarshalJSON() ([]byte, error) {
	return []byte(m.String()), nil
}

func (m *MemberType) UnmarshalJSON(raw []byte) error {
	val, err := ParseMemberType(string(raw))
	if err != nil {
		return errors.WithStack(err)
	}

	*m = val
	return nil
}

func (m MemberType) Value() (driver.Value, error) {
	return []byte(m.String()), nil
}

func (m *MemberType) Scan(raw any) error {
	s, ok := raw.(string)
	if !ok {
		return errors.Errorf("scanned value is not a string: %v", raw)
	}

	val, err := ParseMemberType(s)
	if err != nil {
		return errors.WithStack(err)
	}

	*m = val
	return nil
}

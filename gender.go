package cap

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type Gender int

const (
	NonBinary Gender = iota
	CisMale
	CisFemale
	TransMale
	TransFemale
)

func ParseGender(s string) (Gender, error) {
	switch strings.ToLower(s) {
	case "NB":
		return NonBinary, nil
	case "CM":
		return CisMale, nil
	case "CF":
		return CisFemale, nil
	case "TM":
		return TransMale, nil
	case "TF":
		return TransFemale, nil
	default:
		return -1, errors.Errorf("unrecognized gender: %s", s)
	}
}

func (g Gender) String() string {
	switch g {
	case NonBinary:
		return "NB"
	case CisMale:
		return "CM"
	case CisFemale:
		return "CF"
	case TransMale:
		return "TM"
	case TransFemale:
		return "TF"
	default:
		return ""
	}
}

func (g Gender) MarshalJSON() ([]byte, error) {
	return []byte(g.String()), nil
}

func (g *Gender) UnmarshalJSON(raw []byte) error {
	val, err := ParseGender(string(raw))
	if err != nil {
		return errors.WithStack(err)
	}

	*g = val
	return nil
}

func (g Gender) Value() (driver.Value, error) {
	return []byte(g.String()), nil
}

func (g *Gender) Scan(raw any) error {
	s, ok := raw.(string)
	if !ok {
		return errors.Errorf("scanned value is not a string: %v", raw)
	}

	val, err := ParseGender(s)
	if err != nil {
		return errors.WithStack(err)
	}

	*g = val
	return nil
}

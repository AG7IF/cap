package cap

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type UnitKind int

const (
	UnknownUnitKind UnitKind = iota
	Group
	Squadron
	Flight
)

func ParseUnitKind(s string) UnitKind {
	switch strings.ToLower(s) {
	case "group":
		return Group
	case "squadron":
		return Squadron
	case "flight":
		return Flight
	default:
		return UnknownUnitKind
	}
}

func (u UnitKind) String() string {
	switch u {
	case Group:
		return "Group"
	case Squadron:
		return "Squadron"
	case Flight:
		return "Flight"
	default:
		return ""
	}
}

func (u UnitKind) MarshalJSON() ([]byte, error) {
	return []byte(u.String()), nil
}

func (u *UnitKind) UnmarshalJSON(raw []byte) error {
	val := ParseUnitKind(string(raw))

	*u = val
	return nil
}

func (u UnitKind) Value() (driver.Value, error) {
	return []byte(u.String()), nil
}

func (u *UnitKind) Scan(raw any) error {
	s, ok := raw.(string)
	if !ok {
		return errors.Errorf("scanned value is not a string: %v", raw)
	}

	val := ParseUnitKind(s)

	*u = val
	return nil
}

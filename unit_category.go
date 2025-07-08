package cap

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type UnitCategory int

const (
	UnknownUnitCategory UnitCategory = iota
	AdminUnit
	CompositeUnit
	CadetUnit
	SeniorUnit
)

func ParseUnitCategory(s string) UnitCategory {
	switch strings.ToLower(s) {
	case "admin":
		return AdminUnit
	case "composite":
		return CompositeUnit
	case "cadet":
		return CadetUnit
	case "senior":
		return SeniorUnit
	default:
		return UnknownUnitCategory
	}
}

func (u UnitCategory) String() string {
	switch u {
	case AdminUnit:
		return "Admin"
	case CompositeUnit:
		return "Composite"
	case CadetUnit:
		return "Cadet"
	case SeniorUnit:
		return "Senior"
	default:
		return ""
	}
}

func (u UnitCategory) MarshalJSON() ([]byte, error) {
	return []byte(u.String()), nil
}

func (u *UnitCategory) UnmarshalJSON(raw []byte) error {
	val := ParseUnitCategory(string(raw))

	*u = val
	return nil
}

func (u UnitCategory) Value() (driver.Value, error) {
	return []byte(u.String()), nil
}

func (u *UnitCategory) Scan(raw any) error {
	s, ok := raw.(string)
	if !ok {
		return errors.Errorf("scanned value is not a string: %v", raw)
	}

	val := ParseUnitCategory(s)

	*u = val
	return nil
}

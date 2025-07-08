package cap

import (
	"fmt"
)

type Unit struct {
	charterNumber UnitCharterNumber
	kind          UnitKind
	category      UnitCategory
	name          string
}

func NewUnit(
	charterNumber UnitCharterNumber,
	kind UnitKind,
	category UnitCategory,
	name string,
) Unit {
	return Unit{
		charterNumber: charterNumber,
		kind:          kind,
		category:      category,
		name:          name,
	}
}

func (u Unit) CharterNumber() UnitCharterNumber {
	return u.charterNumber
}

func (u Unit) Kind() UnitKind {
	return u.kind
}

func (u Unit) Category() UnitCategory {
	return u.category
}

func (u Unit) Name() string {
	return u.name
}

func (u Unit) String() string {
	if u.charterNumber.UnitNumber() == 1 {
		return fmt.Sprintf("%s (%s)", u.charterNumber.Wing().Name(), u.kind)
	}

	if u.name == "" || u.kind == UnknownUnitKind || u.category == UnknownUnitCategory {
		return u.charterNumber.String()
	}

	return fmt.Sprintf("%s %s %s (%s)", u.name, u.category, u.kind, u.charterNumber)
}

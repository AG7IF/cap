package cap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCharterNumber(t *testing.T) {
	charter := "RMR-UT-080"
	name := "Blackhawk Cadet Squadron"

	ucn, err := ParseCharterNumber(charter, name)
	assert.NoError(t, err)

	assert.Equal(t, RockyMountainRegion, ucn.Region())
	assert.Equal(t, UTWG, ucn.Wing())
	assert.Equal(t, uint(80), ucn.UnitNumber())
	assert.Equal(t, name, ucn.UnitName())
}
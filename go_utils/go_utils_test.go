package go_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfPointLiesWithinTheUnitCircle(t *testing.T) {
	result := IsPointWithinCircle(0.5, 0.5)
	assert.Equal(t, true, result)
}

func TestIfPointLiesOutsideTheUnitCircle(t *testing.T) {
	result := IsPointWithinCircle(1.5, 1.5)
	assert.Equal(t, false, result)
}

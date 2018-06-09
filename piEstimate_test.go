package main_test

import (
	"testing"

	"github.com/amitdash291/rust_pi_estimation/go_utils"
	"github.com/stretchr/testify/assert"
)

func TestIfPointLiesWithinTheUnitCircle(t *testing.T) {
	result := go_utils.IsPointWithinCircle(0.5, 0.5)
	assert.Equal(t, result, true)
}

package incrementers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMathIsClamped(t *testing.T) {
	require := require.New(t)
	require.True(IsClamped(5, 0, 10), "5 should be clamped between 0 and 10")
	require.False(IsClamped(-5, 0, 10), "-5 should not be clamped between 0 and 10")
	require.False(IsClamped(15, 0, 10), "15 should not be clamped between 0 and 10")
}

func TestMathClamp(t *testing.T) {
	require := require.New(t)
	require.Equal(Clamp(5, 0, 10), 5, "5 should be clamped between 0 and 10")
	require.Equal(Clamp(-5, 0, 10), 0, "-5 should not be clamped between 0 and 10")
	require.Equal(Clamp(15, 0, 10), 10, "15 should not be clamped between 0 and 10")
}

func TestMathClampMin(t *testing.T) {
	require := require.New(t)
	require.Equal(ClampMin(5, 0), 5, "5 should be clamped between 0 and 10")
	require.Equal(ClampMin(-5, 0), 0, "-5 should not be clamped between 0 and 10")
}

func TestMathClampMax(t *testing.T) {
	require := require.New(t)
	require.Equal(ClampMax(5, 10), 5, "5 should be clamped below 10")
	require.Equal(ClampMax(15, 10), 10, "15 should not be clamped below 10")
}

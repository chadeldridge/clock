package rpgtools

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClockNew(t *testing.T) {
	require := require.New(t)
	c := NewClock(4)

	require.Equal(4, c.max)
	require.Equal(0, c.val)
}

func TestClockNewWithTicks(t *testing.T) {
	require := require.New(t)
	c := NewClockWithTicks(4, 3)

	require.Equal(4, c.max)
	require.Equal(3, c.val)
}

func TestClockNewFromJSON(t *testing.T) {
	require := require.New(t)
	c, err := NewClockFromJSON([]byte(`{"max":4,"val":3,"inc":1}`))

	require.NoError(err, "NewClockFromJSON() returned an error: %s", err)
	require.Equal(4, c.max)
	require.Equal(3, c.val)
}
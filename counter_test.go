package rpgtools

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCounterNew(t *testing.T) {
	require := require.New(t)
	c := NewCounter(-4, 4)

	require.Equal(-4, c.Min())
	require.Equal(4, c.Max())
	require.Equal(0, c.Value())
}

func TestCounterNewWithTicks(t *testing.T) {
	require := require.New(t)
	c := NewCounterWithTicks(-4, 4, 3)

	require.Equal(-4, c.Min())
	require.Equal(4, c.Max())
	require.Equal(3, c.Value())
}

func TestCounterNewFromJSON(t *testing.T) {
	require := require.New(t)
	c, err := NewCounterFromJSON([]byte(`{"min":-4,"max":4,"val":3,"inc":1,"orig":0}`))

	require.NoError(err, "NewCounterFromJSON() returned an error: %s", err)
	require.Equal(-4, c.Min())
	require.Equal(4, c.Max())
	require.Equal(3, c.Value())
	require.Equal(1, c.Inc())
	require.Equal(0, c.Original())
}

func TestCounter(t *testing.T) {
	require := require.New(t)
	c := NewCounter(-4, 4)

	require.False(c.IsFull())
	require.True(c.IsEmpty())
	require.False(c.IsMin())

	c.Increment()
	require.Equal(1, c.Value())

	c.Fill()
	require.Equal(4, c.Value())
	require.True(c.IsFull())

	c.Increment()
	require.Equal(4, c.Value())

	c.Decrement()
	require.Equal(3, c.Value())

	c.Empty()
	require.Equal(0, c.Value())

	c.Floor()
	require.Equal(-4, c.Value())

	c.Decrement()
	require.Equal(-4, c.Value())
}

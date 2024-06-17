package rpgtools

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClockNew(t *testing.T) {
	require := require.New(t)
	c := NewClock(4)

	require.Equal(4, c.Max())
	require.Equal(0, c.Value())
}

func TestClockNewWithTicks(t *testing.T) {
	require := require.New(t)
	c := NewClockWithTicks(4, 3)

	require.Equal(4, c.Max())
	require.Equal(3, c.Value())
}

func TestClockNewFromJSON(t *testing.T) {
	require := require.New(t)

	t.Run("valid", func(t *testing.T) {
		c, err := NewClockFromJSON([]byte(`{"nomin":false,"nomax":false,"min":0,"max":4,"val":3,"inc":1,"orig":0}`))
		require.NoError(err, "NewClockFromJSON() returned an error: %s", err)
		require.Equal(4, c.Max())
		require.Equal(3, c.Value())
		require.Equal(0, c.Original())
	})

	t.Run("invalid", func(t *testing.T) {
		// Set min to a string to force an error.
		_, err := NewClockFromJSON([]byte(`{"nomin":false,"nomax":false,"min":b,"max":4,"val":3,"inc":1,"orig":0}`))
		require.Error(err, "NewClockFromJSON() did not return an error: %s", err)
		require.Equal("expected integer", err.Error())
	})

	t.Run("invalid min", func(t *testing.T) {
		_, err := NewClockFromJSON([]byte(`{"nomin":false,"nomax":false,"min":-4,"max":4,"val":3,"inc":1,"orig":0}`))
		require.Error(err, "NewClockFromJSON() did not return an error: %s", err)
		require.Equal("invalid Clock: min must be 0", err.Error())
	})

	t.Run("invalid inc", func(t *testing.T) {
		_, err := NewClockFromJSON([]byte(`{"nomin":false,"nomax":false,"min":0,"max":4,"val":3,"inc":2,"orig":0}`))
		require.Error(err, "NewClockFromJSON() did not return an error: %s", err)
		require.Equal("invalid Clock: inc must be 1", err.Error())
	})
}

func TestClock(t *testing.T) {
	require := require.New(t)
	c := NewClock(4)

	require.False(c.IsFull())
	require.True(c.IsEmpty())

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

	c.Decrement()
	require.Equal(0, c.Value())
}

package incrementers

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
		c, err := NewClockFromJSON([]byte(`{"min":0,"max":4,"incrementer":{"inc":1,"val":2,"orig":3}}`))
		require.NoError(err, "Clock.NewFromJSON() returned an error: %s", err)
		require.Equal(2, c.Value())
		require.Equal(3, c.Original())
	})

	t.Run("invalid incrementer", func(t *testing.T) {
		_, err := NewClockFromJSON([]byte(`{"min":0,"max":0,"incrementer":{"inc":b,"val":2,"orig":3}}`))
		require.Error(err, "Clock.NewFromJSON() did not return an error")
		require.Equal("expected integer", err.Error())
	})

	t.Run("invalid min", func(t *testing.T) {
		_, err := NewClockFromJSON([]byte(`{"min":-4,"max":0,"incrementer":{"inc":1,"val":2,"orig":3}}`))
		require.Error(err, "Clock.NewFromJSON() did not return an error")
		require.Equal("invalid Clock: min must be 0", err.Error())
	})

	t.Run("invalid max", func(t *testing.T) {
		_, err := NewClockFromJSON([]byte(`{"min":0,"max":0,"incrementer":{"inc":1,"val":2,"orig":3}}`))
		require.Error(err, "Clock.NewFromJSON() did not return an error")
		require.Equal("invalid Clock: max must be greater than 0", err.Error())
	})

	t.Run("invalid inc", func(t *testing.T) {
		_, err := NewClockFromJSON([]byte(`{"min":0,"max":4,"incrementer":{"inc":4,"val":2,"orig":3}}`))
		require.Error(err, "Clock.NewFromJSON() did not return an error")
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

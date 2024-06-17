package rpgtools

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCounterNew(t *testing.T) {
	require := require.New(t)
	c := NewCounter()
	require.Equal(0, c.Value())
	require.Equal(1, c.Inc())
	require.Equal(0, c.Original())
}

func TestCounterNewWithValue(t *testing.T) {
	require := require.New(t)

	t.Run("valid", func(t *testing.T) {
		c := NewCounterWithValue(3)
		require.Equal(3, c.Value())
		require.Equal(1, c.Inc())
		require.Equal(3, c.Original())
	})

	t.Run("valid", func(t *testing.T) {
		c := NewCounterWithValue(-3)
		require.Nil(c, "NewCounterWithValue() did not return nil")
	})
}

func TestCounterNewFromJSON(t *testing.T) {
	require := require.New(t)

	t.Run("valid", func(t *testing.T) {
		c, err := NewCounterFromJSON([]byte(
			`{"nomin":false,"nomax":true,"min":0,"max":100,"val":3,"inc":1,"orig":0}`,
		))
		require.NoError(err, "NewCounterFromJSON() returned an error: %s", err)
		require.Equal(3, c.Value())
		require.Equal(1, c.Inc())
		require.Equal(0, c.Original())
	})

	t.Run("invalid", func(t *testing.T) {
		_, err := NewCounterFromJSON([]byte(
			`{"nomin":false,"nomax":true,"min":0,"max":100,"val":b,"inc":1,"orig":0}`,
		))
		require.Error(err, "NewCounterFromJSON() did not return an error")
		require.Equal("expected integer", err.Error())
	})

	t.Run("invalid min", func(t *testing.T) {
		_, err := NewCounterFromJSON([]byte(
			`{"nomin":false,"nomax":true,"min":-4,"max":100,"val":3,"inc":1,"orig":0}`,
		))
		require.Error(err, "NewCounterFromJSON() did not return an error")
		require.Equal("invalid Counter: min must be 0", err.Error())
	})
}

func TestCounter(t *testing.T) {
	require := require.New(t)
	c := NewCounter()

	require.True(c.IsEmpty())

	c.Increment()
	require.Equal(1, c.Value())

	c.SetIncrementer(3)
	c.Increment()
	require.Equal(4, c.Value())

	c.Decrement()
	require.Equal(1, c.Value())

	c.Empty()
	require.Equal(0, c.Value())

	c.Decrement()
	require.Equal(0, c.Value())
}

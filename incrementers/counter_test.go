package incrementers

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

	t.Run("zero", func(t *testing.T) {
		c := NewCounterWithValue(0)
		require.Equal(1, c.Inc())
		require.Equal(0, c.Value())
		require.Equal(0, c.Original())
	})

	t.Run("positive", func(t *testing.T) {
		c := NewCounterWithValue(3)
		require.Equal(1, c.Inc())
		require.Equal(3, c.Value())
		require.Equal(3, c.Original())
	})

	t.Run("negative", func(t *testing.T) {
		c := NewCounterWithValue(-3)
		require.Equal(1, c.Inc())
		require.Equal(0, c.Value())
		require.Equal(0, c.Original())
	})
}

func TestCounterNewFromJSON(t *testing.T) {
	require := require.New(t)

	t.Run("valid", func(t *testing.T) {
		c, err := NewCounterFromJSON([]byte(`{"min":0,"max":0,"incrementer":{"inc":1,"val":2,"orig":3}}`))
		require.NoError(err, "Counter.NewFromJSON() returned an error: %s", err)
		require.Equal(1, c.Inc())
		require.Equal(2, c.Value())
		require.Equal(3, c.Original())
	})

	t.Run("invalid incrementer", func(t *testing.T) {
		_, err := NewCounterFromJSON([]byte(`{"min":0,"max":0,"incrementer":{"inc":b,"val":2,"orig":3}}`))
		require.Error(err, "Counter.NewFromJSON() did not return an error")
		require.Equal("expected integer", err.Error())
	})

	t.Run("invalid min", func(t *testing.T) {
		_, err := NewCounterFromJSON([]byte(`{"min":-4,"max":0,"incrementer":{"inc":1,"val":2,"orig":3}}`))
		require.Error(err, "Counter.NewFromJSON() did not return an error")
		require.Equal("invalid Counter: min must be 0", err.Error())
	})

	t.Run("invalid max", func(t *testing.T) {
		_, err := NewCounterFromJSON([]byte(`{"min":0,"max":4,"incrementer":{"inc":1,"val":2,"orig":3}}`))
		require.Error(err, "Counter.NewFromJSON() did not return an error")
		require.Equal("invalid Counter: max must be 0", err.Error())
	})
}

/*
func TestCounter(t *testing.T) {
	require := require.New(t)
	c := NewCounter()

	require.True(c.IsEmpty())

	c.Inc()rement()
	require.Equal(1, c.Value()ue())

	c.SetIncrementer(3)
	c.Inc()rement()
	require.Equal(4, c.Value()ue())

	c.Decrement()
	require.Equal(1, c.Value()ue())

	c.Empty()
	require.Equal(0, c.Value()ue())

	c.Decrement()
	require.Equal(0, c.Value()ue())
}
*/

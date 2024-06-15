package incrementers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUIncrementerNewUIncrementer(t *testing.T) {
	require := require.New(t)
	c := NewUIncrementer(4)

	require.Equal(4, c.max)
	require.Equal(0, c.val)
}

func TestUIncrementerNewUIncrementerWithValue(t *testing.T) {
	require := require.New(t)
	c := NewUIncrementerWithValue(4, 3)

	require.Equal(4, c.max)
	require.Equal(3, c.val)
}

func TestUIncrementerNewUIncrementerFromJSON(t *testing.T) {
	require := require.New(t)
	c, err := NewUIncrementerFromJSON([]byte(`{"max":4,"val":3,"inc":0}`))

	require.NoError(err, "UIncrementer.NewFromJSON() returned an error: %s", err)
	require.Equal(4, c.max)
	require.Equal(3, c.val)
}

func TestUIncrementerMax(t *testing.T) {
	require := require.New(t)
	c := UIncrementer{max: 4}

	t.Run("default", func(t *testing.T) {
		require.Equal(4, c.Max())
	})

	t.Run("changed", func(t *testing.T) {
		c.max = 6
		require.Equal(6, c.Max())
	})
}

func TestUIncrementerValue(t *testing.T) {
	require := require.New(t)
	c := UIncrementer{max: 4}

	t.Run("default", func(t *testing.T) {
		require.Equal(0, c.Value())
	})

	t.Run("changed", func(t *testing.T) {
		c.val = 4
		require.Equal(4, c.Value())
	})
}

func TestUIncrementerIncrementer(t *testing.T) {
	require := require.New(t)
	c := UIncrementer{max: 4}

	t.Run("default", func(t *testing.T) {
		require.Equal(1, c.Incrementer())
	})

	t.Run("changed", func(t *testing.T) {
		c.inc = 3
		require.Equal(3, c.Incrementer())
	})
}

func TestUIncrementerIsFull(t *testing.T) {
	require := require.New(t)
	c := UIncrementer{max: 4}

	t.Run("empty", func(t *testing.T) {
		require.False(c.IsFull())
	})

	t.Run("full", func(t *testing.T) {
		c.val = 4
		require.True(c.IsFull())
	})
}

func TestUIncrementerIsEmpty(t *testing.T) {
	require := require.New(t)
	c := UIncrementer{max: 4}

	t.Run("empty", func(t *testing.T) {
		require.True(c.IsEmpty())
	})

	t.Run("full", func(t *testing.T) {
		c.val = 4
		require.False(c.IsEmpty())
	})
}

func TestUIncrementerIncrementerBy(t *testing.T) {
	require := require.New(t)
	c := UIncrementer{max: 4}

	c.IncrementBy(3)
	require.Equal(3, c.inc)
}

func TestUIncrementerAdvance(t *testing.T) {
	require := require.New(t)
	c := UIncrementer{max: 4}

	t.Run("default", func(t *testing.T) {
		c.Advance()
		require.Equal(1, c.val)
	})

	t.Run("changed", func(t *testing.T) {
		c.val = 0
		c.inc = 3
		c.Advance()
		require.Equal(3, c.val)
	})
}

func TestUIncrementerIncrement(t *testing.T) {
	require := require.New(t)
	c := UIncrementer{max: 4}

	t.Run("empty", func(t *testing.T) {
		c.Increment()
		require.Equal(1, c.val)
	})

	t.Run("full", func(t *testing.T) {
		c.val = 4
		c.Increment()
		require.Equal(4, c.val)
	})
}

func TestUIncrementerDecrement(t *testing.T) {
	require := require.New(t)
	c := UIncrementer{max: 4}

	t.Run("empty", func(t *testing.T) {
		c.Decrement()
		require.Equal(0, c.val)
	})

	t.Run("full", func(t *testing.T) {
		c.val = 4
		c.Decrement()
		require.Equal(3, c.val)
	})
}

func TestUIncrementerAdd(t *testing.T) {
	require := require.New(t)
	c := UIncrementer{max: 4}

	t.Run("empty", func(t *testing.T) {
		c.Add(3)
		require.Equal(3, c.val)
	})

	t.Run("fill", func(t *testing.T) {
		c.val = 3
		c.Add(3)
		require.Equal(4, c.val)
	})

	t.Run("full", func(t *testing.T) {
		c.val = 4
		c.Add(3)
		require.Equal(4, c.val)
	})
}

func TestUIncrementerRemove(t *testing.T) {
	require := require.New(t)
	c := UIncrementer{max: 4}

	t.Run("while empty", func(t *testing.T) {
		c.Remove(3)
		require.Equal(0, c.val)
	})

	t.Run("remove", func(t *testing.T) {
		c.val = 3
		c.Remove(2)
		require.Equal(1, c.val)
	})

	t.Run("empty", func(t *testing.T) {
		c.val = 3
		c.Remove(3)
		require.Equal(0, c.val)
	})
}

func TestUIncrementerSetValue(t *testing.T) {
	require := require.New(t)
	c := UIncrementer{max: 4}

	t.Run("minimum", func(t *testing.T) {
		c.SetValue(-1)
		require.Equal(0, c.val)
	})

	t.Run("valid", func(t *testing.T) {
		c.SetValue(3)
		require.Equal(3, c.val)
	})

	t.Run("maximum", func(t *testing.T) {
		c.SetValue(6)
		require.Equal(4, c.val)
	})
}

func TestUIncrementerSetMax(t *testing.T) {
	require := require.New(t)
	c := UIncrementer{max: 4}

	t.Run("minimum", func(t *testing.T) {
		c.SetMax(0)
		require.Equal(1, c.max)
	})

	t.Run("set", func(t *testing.T) {
		c.SetMax(6)
		require.Equal(6, c.max)
	})

	t.Run("with val", func(t *testing.T) {
		c.val = 4
		c.SetMax(3)
		require.Equal(3, c.max)
		require.Equal(3, c.val)
	})
}

func TestUIncrementerAddMax(t *testing.T) {
	require := require.New(t)
	c := UIncrementer{max: 4}

	c.AddMax(3)
	require.Equal(7, c.max)
}

func TestUIncrementerRemoveMax(t *testing.T) {
	require := require.New(t)
	c := UIncrementer{max: 6}

	t.Run("remove", func(t *testing.T) {
		c.RemoveMax(2)
		require.Equal(4, c.max)
	})

	t.Run("minimum", func(t *testing.T) {
		c.max = 4
		c.RemoveMax(6)
		require.Equal(1, c.max)
	})

	t.Run("with val", func(t *testing.T) {
		c.max = 6
		c.val = 6
		c.RemoveMax(2)
		require.Equal(4, c.max)
		require.Equal(4, c.val)
	})
}

func TestUIncrementerFill(t *testing.T) {
	require := require.New(t)
	c := UIncrementer{max: 4}

	c.Fill()
	require.Equal(4, c.val)
}

func TestUIncrementerEmpty(t *testing.T) {
	require := require.New(t)
	c := UIncrementer{max: 4}

	c.Empty()
	require.Equal(0, c.val)
}

func TestUIncrementerReset(t *testing.T) {
	require := require.New(t)
	c := UIncrementer{max: 4}

	c.val = 4
	c.Reset()
	require.Equal(0, c.val)
}

func TestUIncrementerString(t *testing.T) {
	require := require.New(t)
	c := UIncrementer{max: 4}

	t.Run("empty", func(t *testing.T) {
		require.Equal("0/4", c.String())
	})

	t.Run("with val", func(t *testing.T) {
		c.val = 3
		require.Equal("3/4", c.String())
	})
}

func TestUIncrementerMarshalJSON(t *testing.T) {
	require := require.New(t)
	c := UIncrementer{max: 4}

	t.Run("empty", func(t *testing.T) {
		data, err := c.MarshalJSON()
		require.NoError(err, "UIncrementer.MarshalJSON() returned an error: %s", err)
		require.Equal(`{"max":4,"val":0,"inc":0}`, string(data))
	})

	t.Run("with val", func(t *testing.T) {
		c.val = 3
		data, err := c.MarshalJSON()
		require.NoError(err, "UIncrementer.MarshalJSON() returned an error: %s", err)
		require.Equal(`{"max":4,"val":3,"inc":0}`, string(data))
	})
}

func TestUIncrementerUnmarshalJSON(t *testing.T) {
	require := require.New(t)
	c := UIncrementer{}

	t.Run("scan error", func(t *testing.T) {
		err := c.UnmarshalJSON([]byte(`{"max":b,"val":0,"inc":0}`))
		require.Error(err, "UIncrementer.UnmarshalJSON() did not return an error")
		require.Equal("expected integer", err.Error())
	})

	t.Run("invalid max", func(t *testing.T) {
		err := c.UnmarshalJSON([]byte(`{"max":0,"val":0,"inc":0}`))
		require.Error(err, "UIncrementer.UnmarshalJSON() did not return an error")
		require.Equal("invalid UIncrementer: max must be greater than 0", err.Error())
	})

	t.Run("invalid val", func(t *testing.T) {
		err := c.UnmarshalJSON([]byte(`{"max":4,"val":-1,"inc":0}`))
		require.Error(err, "UIncrementer.UnmarshalJSON() did not return an error")
		require.Equal("invalid UIncrementer: val must be greater than or equal to 0", err.Error())
	})

	t.Run("too many val", func(t *testing.T) {
		err := c.UnmarshalJSON([]byte(`{"max":4,"val":5,"inc":0}`))
		require.Error(err, "UIncrementer.UnmarshalJSON() did not return an error")
		require.Equal("invalid UIncrementer: val must be less than or equal to max", err.Error())
	})

	t.Run("negative inc", func(t *testing.T) {
		err := c.UnmarshalJSON([]byte(`{"max":4,"val":3,"inc":-1}`))
		require.Error(err, "UIncrementer.UnmarshalJSON() did not return an error")
		require.Equal("invalid UIncrementer: inc must be greater than or equal to 0", err.Error())
	})

	t.Run("valid", func(t *testing.T) {
		err := c.UnmarshalJSON([]byte(`{"max":4,"val":3,"inc":0}`))
		require.NoError(err, "UIncrementer.UnmarshalJSON() returned an error: %s", err)
		require.Equal(4, c.max)
		require.Equal(3, c.val)
	})
}

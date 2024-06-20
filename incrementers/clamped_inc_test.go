package incrementers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClampedIncrementerNew(t *testing.T) {
	require := require.New(t)
	c := NewClampedIncrementer(-4, 4)
	require.Equal(-4, c.min)
	require.Equal(4, c.max)
	require.Equal(1, c.Inc())
	require.Equal(0, c.Value())
	require.Equal(0, c.Original())
}

func TestClampedIncrementerNewWithValue(t *testing.T) {
	require := require.New(t)

	t.Run("zero", func(t *testing.T) {
		c := NewClampedIncrementerWithValue(-4, 4, 0)
		require.Equal(-4, c.min)
		require.Equal(4, c.max)
		require.Equal(1, c.inc)
		require.Equal(0, c.val)
		require.Equal(0, c.orig)
	})

	t.Run("positive", func(t *testing.T) {
		c := NewClampedIncrementerWithValue(-4, 4, 3)
		require.Equal(-4, c.min)
		require.Equal(4, c.max)
		require.Equal(1, c.inc)
		require.Equal(3, c.val)
		require.Equal(3, c.orig)
	})

	t.Run("negative", func(t *testing.T) {
		c := NewClampedIncrementerWithValue(-4, 4, -3)
		require.Equal(-4, c.min)
		require.Equal(4, c.max)
		require.Equal(1, c.inc)
		require.Equal(-3, c.val)
		require.Equal(-3, c.orig)
	})

	t.Run("no max", func(t *testing.T) {
		c := NewClampedIncrementerWithValue(-4, 0, 25)
		require.Equal(-4, c.min)
		require.Equal(0, c.max)
		require.Equal(1, c.inc)
		require.Equal(25, c.val)
		require.Equal(25, c.orig)
	})
}

func TestClampedIncrementerNewFromJSON(t *testing.T) {
	require := require.New(t)
	c, err := NewClampedIncrementerFromJSON([]byte(`{"min":-4,"max":4,"incrementer":{"inc":1,"val":2,"orig":3}}`))
	require.NoError(err, "ClampedIncrementer.NewFromJSON() returned an error: %s", err)
	require.Equal(-4, c.min)
	require.Equal(4, c.max)
	require.Equal(1, c.inc)
	require.Equal(2, c.val)
	require.Equal(3, c.orig)
}

func TestClampedIncrementerIsFull(t *testing.T) {
	require := require.New(t)
	c := ClampedIncrementer{min: -4, max: 4, Incrementer: Incrementer{}}

	t.Run("empty", func(t *testing.T) {
		require.False(c.IsFull())
	})

	t.Run("full", func(t *testing.T) {
		c.val = 4
		require.True(c.IsFull())
	})
}

func TestClampedIncrementerMin(t *testing.T) {
	require := require.New(t)
	c := ClampedIncrementer{min: -4, max: 4, Incrementer: Incrementer{}}
	require.Equal(-4, c.Min())
}

func TestClampedIncrementerMax(t *testing.T) {
	require := require.New(t)
	c := ClampedIncrementer{min: -4, max: 4, Incrementer: Incrementer{}}
	require.Equal(4, c.Max())
}

func TestClampedIncrementerIncrement(t *testing.T) {
	require := require.New(t)
	c := ClampedIncrementer{min: -4, max: 4, Incrementer: Incrementer{}}

	t.Run("zero", func(t *testing.T) {
		c.val = 0
		c.Increment()
		require.Equal(0, c.val)
	})

	t.Run("positive", func(t *testing.T) {
		c.val = 0
		c.inc = 1
		c.Increment()
		require.Equal(1, c.val)
	})

	t.Run("negative", func(t *testing.T) {
		c.val = 4
		c.inc = -1
		c.Increment()
		require.Equal(3, c.val)
	})

	t.Run("min", func(t *testing.T) {
		c.val = -4
		c.inc = -1
		c.Increment()
		require.Equal(-4, c.val)
	})

	t.Run("max", func(t *testing.T) {
		c.val = 4
		c.inc = 1
		c.Increment()
		require.Equal(4, c.val)
	})
}

func TestClampedIncrementerDecrement(t *testing.T) {
	require := require.New(t)
	c := ClampedIncrementer{min: -4, max: 4, Incrementer: Incrementer{}}

	t.Run("zero", func(t *testing.T) {
		c.val = 4
		c.inc = 0
		c.Decrement()
		require.Equal(4, c.val)
	})

	t.Run("positive", func(t *testing.T) {
		c.val = 4
		c.inc = 1
		c.Decrement()
		require.Equal(3, c.val)
	})

	t.Run("negative", func(t *testing.T) {
		c.val = 0
		c.inc = -1
		c.Decrement()
		require.Equal(1, c.val)
	})

	t.Run("min", func(t *testing.T) {
		c.val = -4
		c.inc = 1
		c.Decrement()
		require.Equal(-4, c.val)
	})

	t.Run("max", func(t *testing.T) {
		c.val = 4
		c.inc = -1
		c.Decrement()
		require.Equal(4, c.val)
	})
}

func TestClampedIncrementerAdd(t *testing.T) {
	require := require.New(t)
	c := ClampedIncrementer{min: -4, max: 4, Incrementer: Incrementer{}}

	t.Run("zero", func(t *testing.T) {
		c.val = 4
		c.Add(0)
		require.Equal(4, c.val)
	})

	t.Run("positive", func(t *testing.T) {
		c.val = 0
		c.Add(4)
		require.Equal(4, c.val)
	})

	t.Run("positive", func(t *testing.T) {
		c.val = 4
		c.Add(-4)
		require.Equal(0, c.val)
	})

	t.Run("min", func(t *testing.T) {
		c.val = -4
		c.Add(-1)
		require.Equal(-4, c.val)
	})

	t.Run("max", func(t *testing.T) {
		c.val = 4
		c.Add(1)
		require.Equal(4, c.val)
	})
}

func TestClampedIncrementerRemove(t *testing.T) {
	require := require.New(t)
	c := ClampedIncrementer{min: -4, max: 4, Incrementer: Incrementer{}}

	t.Run("zero", func(t *testing.T) {
		c.val = 4
		c.Remove(0)
		require.Equal(4, c.val)
	})

	t.Run("positive", func(t *testing.T) {
		c.val = 4
		c.Remove(4)
		require.Equal(0, c.val)
	})

	t.Run("negative", func(t *testing.T) {
		c.val = 0
		c.Remove(-4)
		require.Equal(4, c.val)
	})

	t.Run("min", func(t *testing.T) {
		c.val = -4
		c.Remove(1)
		require.Equal(-4, c.val)
	})

	t.Run("max", func(t *testing.T) {
		c.val = 4
		c.Remove(-1)
		require.Equal(4, c.val)
	})
}

func TestClampedIncrementerSetMin(t *testing.T) {
	require := require.New(t)
	c := ClampedIncrementer{min: -4, max: 4, Incrementer: Incrementer{}}

	t.Run("zero", func(t *testing.T) {
		c.SetMin(0)
		require.Equal(0, c.min)
	})

	t.Run("positive", func(t *testing.T) {
		c.SetMin(1)
		require.Equal(1, c.min)
	})

	t.Run("negative", func(t *testing.T) {
		c.SetMin(-1)
		require.Equal(-1, c.min)
	})
}

func TestClampedIncrementerSetMax(t *testing.T) {
	require := require.New(t)
	c := ClampedIncrementer{min: -4, max: 4, Incrementer: Incrementer{}}

	t.Run("zero", func(t *testing.T) {
		c.SetMax(0)
		require.Equal(0, c.max)
	})

	t.Run("positive", func(t *testing.T) {
		c.SetMax(1)
		require.Equal(1, c.max)
	})

	t.Run("negative", func(t *testing.T) {
		c.SetMax(-1)
		require.Equal(-1, c.max)
	})
}

func TestClampedIncrementerSetValue(t *testing.T) {
	require := require.New(t)
	c := ClampedIncrementer{min: -4, max: 4, Incrementer: Incrementer{}}

	t.Run("positive", func(t *testing.T) {
		c.SetValue(4)
		require.Equal(4, c.val)
	})

	t.Run("negative", func(t *testing.T) {
		c.SetValue(-4)
		require.Equal(-4, c.val)
	})

	t.Run("zero", func(t *testing.T) {
		c.val = 4
		c.SetValue(0)
		require.Equal(0, c.val)
	})

	t.Run("min", func(t *testing.T) {
		c.SetValue(-5)
		require.Equal(-4, c.val)
	})

	t.Run("max", func(t *testing.T) {
		c.SetValue(5)
		require.Equal(4, c.val)
	})
}

func TestClampedIncrementerSetOriginalValue(t *testing.T) {
	require := require.New(t)
	c := ClampedIncrementer{min: -4, max: 4, Incrementer: Incrementer{}}

	t.Run("positive", func(t *testing.T) {
		c.SetOriginalValue(4)
		require.Equal(4, c.orig)
	})

	t.Run("negative", func(t *testing.T) {
		c.SetOriginalValue(-4)
		require.Equal(-4, c.orig)
	})

	t.Run("zero", func(t *testing.T) {
		c.SetOriginalValue(0)
		require.Equal(0, c.orig)
	})

	t.Run("min", func(t *testing.T) {
		c.SetOriginalValue(-5)
		require.Equal(-4, c.orig)
	})

	t.Run("max", func(t *testing.T) {
		c.SetOriginalValue(5)
		require.Equal(4, c.orig)
	})

	t.Run("no max", func(t *testing.T) {
		c.SetOriginalValue(5)
		require.Equal(4, c.orig)
	})
}

func TestClampedIncrementerClamp(t *testing.T) {
	require := require.New(t)
	c := ClampedIncrementer{min: -4, max: 4, Incrementer: Incrementer{}}

	t.Run("min", func(t *testing.T) {
		c.val = -5
		c.Clamp()
		require.Equal(-4, c.val)
	})

	t.Run("max", func(t *testing.T) {
		c.val = 5
		c.Clamp()
		require.Equal(4, c.val)
	})

	t.Run("no max", func(t *testing.T) {
		c.max = 0
		c.val = 5
		c.Clamp()
		require.Equal(5, c.val)
	})
}

func TestClampedIncrementerClampOriginalValue(t *testing.T) {
	require := require.New(t)
	c := ClampedIncrementer{min: -4, max: 4, Incrementer: Incrementer{}}

	t.Run("min", func(t *testing.T) {
		c.orig = -5
		c.ClampOriginalValue()
		require.Equal(-4, c.orig)
	})

	t.Run("max", func(t *testing.T) {
		c.orig = 5
		c.ClampOriginalValue()
		require.Equal(4, c.orig)
	})

	t.Run("no max", func(t *testing.T) {
		c.max = 0
		c.orig = 5
		c.ClampOriginalValue()
		require.Equal(5, c.orig)
	})
}

func TestClampedIncrementerFill(t *testing.T) {
	require := require.New(t)
	c := ClampedIncrementer{min: -4, max: 4, Incrementer: Incrementer{}}

	c.Fill()
	require.Equal(4, c.val)
}

func TestClampedIncrementerFloor(t *testing.T) {
	require := require.New(t)
	c := ClampedIncrementer{min: -4, max: 4, Incrementer: Incrementer{}}

	c.Floor()
	require.Equal(-4, c.val)
}

func TestClampedIncrementerString(t *testing.T) {
	require := require.New(t)
	c := ClampedIncrementer{min: -4, max: 4, Incrementer: Incrementer{}}

	t.Run("default", func(t *testing.T) {
		require.Equal("0/4", c.String())
	})

	t.Run("positive", func(t *testing.T) {
		c.val = 4
		require.Equal("4/4", c.String())
	})

	t.Run("negative", func(t *testing.T) {
		c.val = -4
		require.Equal("-4/4", c.String())
	})
}

func TestClampedIncrementerMarshalJSON(t *testing.T) {
	require := require.New(t)
	c := ClampedIncrementer{min: -4, max: 4, Incrementer: Incrementer{}}

	t.Run("empty", func(t *testing.T) {
		data, err := c.MarshalJSON()
		require.NoError(err, "Incrementer.MarshalJSON() returned an error: %s", err)
		require.Equal(`{"min":-4,"max":4,"incrementer":{"inc":0,"val":0,"orig":0}}`, string(data))
	})

	t.Run("set", func(t *testing.T) {
		c = ClampedIncrementer{min: -4, max: 4, Incrementer: Incrementer{inc: 1, val: 2, orig: 3}}
		data, err := c.MarshalJSON()
		require.NoError(err, "Incrementer.MarshalJSON() returned an error: %s", err)
		require.Equal(`{"min":-4,"max":4,"incrementer":{"inc":1,"val":2,"orig":3}}`, string(data))
	})
}

func TestClampedIncrementerUnmarshalJSON(t *testing.T) {
	require := require.New(t)

	t.Run("valid", func(t *testing.T) {
		c := ClampedIncrementer{}
		err := c.UnmarshalJSON([]byte(`{"min":-4,"max":4,"incrementer":{"inc":1,"val":2,"orig":3}}`))
		require.NoError(err, "ClampedIncrementer.UnmarshalJSON() returned an error: %s", err)
		require.Equal(1, c.inc)
		require.Equal(2, c.val)
		require.Equal(3, c.orig)
	})

	t.Run("nil", func(t *testing.T) {
		c := ClampedIncrementer{}
		err := c.UnmarshalJSON(nil)
		require.Error(err, "ClampedIncrementer.UnmarshalJSON() did not return an error")
		require.Equal("Incrementer.UnmarshalJSON(): data was nil", err.Error())
	})

	t.Run("null", func(t *testing.T) {
		c := ClampedIncrementer{}
		err := c.UnmarshalJSON([]byte(`null`))
		require.NoError(err, "ClampedIncrementer.UnmarshalJSON() returned an error: %s", err)
		require.Equal(ClampedIncrementer{}, c)
	})

	t.Run("empty", func(t *testing.T) {
		c := ClampedIncrementer{}
		err := c.UnmarshalJSON([]byte(`""`))
		require.NoError(err, "ClampedIncrementer.UnmarshalJSON() returned an error: %s", err)
		require.Equal(ClampedIncrementer{}, c)
	})

	t.Run("scan error", func(t *testing.T) {
		c := ClampedIncrementer{}
		// Make "min" a string to force error.
		err := c.UnmarshalJSON([]byte(`{"min":b,"max":4,"incrementer":{"inc":1,"val":2,"orig":3}}`))
		require.Error(err, "ClampedIncrementer.UnmarshalJSON() did not return an error")
		require.Equal("expected integer", err.Error())
	})

	t.Run("invalid min", func(t *testing.T) {
		c := ClampedIncrementer{}
		err := c.UnmarshalJSON([]byte(`{"min":5,"max":4,"incrementer":{"inc":1,"val":2,"orig":3}}`))
		require.Error(err, "ClampedIncrementer.UnmarshalJSON() did not return an error")
		require.Equal("invalid ClampedIncrementer: min must be less than max", err.Error())
	})

	t.Run("val below min", func(t *testing.T) {
		c := ClampedIncrementer{}
		err := c.UnmarshalJSON([]byte(`{"min":3,"max":4,"incrementer":{"inc":1,"val":2,"orig":3}}`))
		require.Error(err, "ClampedIncrementer.UnmarshalJSON() did not return an error")
		require.Equal("invalid ClampedIncrementer: Incrementer.val must be min or greater", err.Error())
	})

	t.Run("orig below min", func(t *testing.T) {
		c := ClampedIncrementer{}
		err := c.UnmarshalJSON([]byte(`{"min":3,"max":4,"incrementer":{"inc":1,"val":3,"orig":2}}`))
		require.Error(err, "ClampedIncrementer.UnmarshalJSON() did not return an error")
		require.Equal("invalid ClampedIncrementer: Incrementer.orig must be min or greater", err.Error())
	})

	t.Run("invalid incrementer", func(t *testing.T) {
		c := ClampedIncrementer{}
		// Make "inc" a string to force error.
		err := c.UnmarshalJSON([]byte(`{"min":-4,"max":4,"incrementer":{"inc":b,"val":2,"orig":3}}`))
		require.Error(err, "ClampedIncrementer.UnmarshalJSON() did not return error")
		require.Equal("expected integer", err.Error())
	})

	t.Run("invalid val", func(t *testing.T) {
		c := ClampedIncrementer{}
		err := c.UnmarshalJSON([]byte(`{"min":-4,"max":4,"incrementer":{"inc":1,"val":6,"orig":3}}`))
		require.Error(err, "ClampedIncrementer.UnmarshalJSON() did not return an error")
		require.Equal("invalid ClampedIncrementer: Incrementer.val must min <= val <= max", err.Error())
	})

	t.Run("invalid orig", func(t *testing.T) {
		c := ClampedIncrementer{}
		err := c.UnmarshalJSON([]byte(`{"min":-4,"max":4,"incrementer":{"inc":1,"val":2,"orig":6}}`))
		require.Error(err, "ClampedIncrementer.UnmarshalJSON() did not return an error")
		require.Equal("invalid ClampedIncrementer: Incrementer.orig must min <= orig <= max", err.Error())
	})
}

/*
func TestClampedIncrementer(t *testing.T) {
	require := require.New(t)
	c := NewClampedIncrementer()

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
*/

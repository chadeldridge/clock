package incrementers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIncrementerNewIncrementer(t *testing.T) {
	require := require.New(t)
	i := NewIncrementer(0, 4)

	require.Equal(0, i.min)
	require.Equal(4, i.max)
	require.Equal(0, i.val)
}

func TestIncrementerNewIncrementerWithValue(t *testing.T) {
	require := require.New(t)
	i := NewIncrementerWithValue(0, 4, 3)

	require.Equal(0, i.min)
	require.Equal(4, i.max)
	require.Equal(3, i.val)
}

func TestIncrementerNewIncrementerFromJSON(t *testing.T) {
	require := require.New(t)
	i, err := NewIncrementerFromJSON([]byte(`{"min":0,"max":4,"val":3,"inc":1,"orig":0}`))

	require.NoError(err, "Incrementer.NewFromJSON() returned an error: %s", err)
	require.Equal(0, i.min)
	require.Equal(4, i.max)
	require.Equal(3, i.val)
	require.Equal(1, i.inc)
	require.Equal(0, i.orig)
}

func TestIncrementerMin(t *testing.T) {
	require := require.New(t)
	i := Incrementer{max: 4}

	t.Run("default", func(t *testing.T) {
		require.Equal(0, i.Min())
	})

	t.Run("positive", func(t *testing.T) {
		i.min = 2
		require.Equal(2, i.Min())
	})
	t.Run("negative", func(t *testing.T) {
		i.min = -4
		require.Equal(-4, i.Min())
	})
}

func TestIncrementerMax(t *testing.T) {
	require := require.New(t)
	i := Incrementer{max: 4}

	t.Run("default", func(t *testing.T) {
		require.Equal(4, i.Max())
	})

	t.Run("positive", func(t *testing.T) {
		i.max = 6
		require.Equal(6, i.Max())
	})

	t.Run("negative", func(t *testing.T) {
		i.max = -4
		require.Equal(-4, i.Max())
	})
}

func TestIncrementerValue(t *testing.T) {
	require := require.New(t)
	i := Incrementer{max: 4}

	t.Run("default", func(t *testing.T) {
		require.Equal(0, i.Value())
	})

	t.Run("positive", func(t *testing.T) {
		i.val = 4
		require.Equal(4, i.Value())
	})

	t.Run("negative", func(t *testing.T) {
		i.val = -4
		require.Equal(-4, i.Value())
	})
}

func TestIncrementerIncrementer(t *testing.T) {
	require := require.New(t)
	i := Incrementer{max: 4}

	t.Run("default", func(t *testing.T) {
		require.Equal(0, i.Inc())
	})

	t.Run("positive", func(t *testing.T) {
		i.inc = 1
		require.Equal(1, i.Inc())
	})

	t.Run("negative", func(t *testing.T) {
		i.inc = -1
		require.Equal(-1, i.Inc())
	})
}

func TestIncrementerOriginal(t *testing.T) {
	require := require.New(t)
	i := Incrementer{max: 4}

	t.Run("default", func(t *testing.T) {
		require.Equal(0, i.Original())
	})

	t.Run("positive", func(t *testing.T) {
		i.orig = 4
		require.Equal(4, i.Original())
	})

	t.Run("negative", func(t *testing.T) {
		i.orig = -4
		require.Equal(-4, i.Original())
	})
}

func TestIncrementerIsFull(t *testing.T) {
	require := require.New(t)
	i := Incrementer{min: -1, max: 4}

	t.Run("zero", func(t *testing.T) {
		require.False(i.IsFull())
	})

	t.Run("full", func(t *testing.T) {
		i.val = 4
		require.True(i.IsFull())
	})
}

func TestIncrementerIsEmpty(t *testing.T) {
	require := require.New(t)
	i := Incrementer{max: 4}

	t.Run("empty", func(t *testing.T) {
		require.True(i.IsEmpty())
	})

	t.Run("not empty", func(t *testing.T) {
		i.val = 1
		require.False(i.IsEmpty())
	})

	t.Run("full", func(t *testing.T) {
		i.val = 4
		require.False(i.IsEmpty())
	})
}

func TestIncrementerIsMin(t *testing.T) {
	require := require.New(t)
	i := Incrementer{min: -4, max: 4}

	t.Run("zero", func(t *testing.T) {
		require.False(i.IsMin())
	})

	t.Run("min", func(t *testing.T) {
		i.val = -4
		require.True(i.IsMin())
	})
}

func TestIncrementerIncrementerBy(t *testing.T) {
	require := require.New(t)
	i := Incrementer{max: 4}

	t.Run("default", func(t *testing.T) {
		require.Equal(0, i.inc)
	})

	t.Run("positive", func(t *testing.T) {
		i.IncrementBy(1)
		require.Equal(1, i.inc)
	})

	t.Run("negative", func(t *testing.T) {
		i.IncrementBy(-1)
		require.Equal(-1, i.inc)
	})
}

func TestIncrementerAdvance(t *testing.T) {
	require := require.New(t)
	i := Incrementer{min: -4, max: 4}

	t.Run("default", func(t *testing.T) {
		i.Advance()
		require.Equal(0, i.val)
	})

	t.Run("positive", func(t *testing.T) {
		i.val = 0
		i.inc = 1
		i.Advance()
		require.Equal(1, i.val)
	})

	t.Run("negative", func(t *testing.T) {
		i.val = 0
		i.inc = -1
		i.Advance()
		require.Equal(-1, i.val)
	})
}

func TestIncrementerIncrement(t *testing.T) {
	require := require.New(t)
	i := Incrementer{min: -4, max: 4}

	t.Run("empty", func(t *testing.T) {
		i.Increment()
		require.Equal(0, i.val)
	})

	t.Run("positive", func(t *testing.T) {
		i.val = 0
		i.inc = 1
		i.Increment()
		require.Equal(1, i.val)
	})

	t.Run("negative", func(t *testing.T) {
		i.val = 0
		i.inc = -1
		i.Increment()
		require.Equal(-1, i.val)
	})
}

func TestIncrementerDecrement(t *testing.T) {
	require := require.New(t)
	i := Incrementer{min: -4, max: 4}

	t.Run("empty", func(t *testing.T) {
		i.Decrement()
		require.Equal(0, i.val)
	})

	t.Run("positive", func(t *testing.T) {
		i.val = 0
		i.inc = 1
		i.Decrement()
		require.Equal(-1, i.val)
	})

	t.Run("negative", func(t *testing.T) {
		i.val = 0
		i.inc = -1
		i.Decrement()
		require.Equal(1, i.val)
	})
}

func TestIncrementerAdd(t *testing.T) {
	require := require.New(t)
	i := Incrementer{min: -4, max: 4}

	t.Run("empty", func(t *testing.T) {
		i.Add(3)
		require.Equal(3, i.val)
	})

	t.Run("positive", func(t *testing.T) {
		i.val = 3
		i.Add(3)
		require.Equal(4, i.val)
	})

	t.Run("positive", func(t *testing.T) {
		i.val = 3
		i.Add(-3)
		require.Equal(0, i.val)
	})

	t.Run("full", func(t *testing.T) {
		i.val = 4
		i.Add(3)
		require.Equal(4, i.val)
	})

	t.Run("min", func(t *testing.T) {
		i.val = -4
		i.Add(-3)
		require.Equal(-4, i.val)
	})
}

func TestIncrementerRemove(t *testing.T) {
	require := require.New(t)
	i := Incrementer{min: -4, max: 4}

	t.Run("empty", func(t *testing.T) {
		i.Remove(3)
		require.Equal(-3, i.val)
	})

	t.Run("positive", func(t *testing.T) {
		i.val = 3
		i.Remove(3)
		require.Equal(0, i.val)
	})

	t.Run("positive", func(t *testing.T) {
		i.val = 3
		i.Remove(-3)
		require.Equal(4, i.val)
	})

	t.Run("full", func(t *testing.T) {
		i.val = 4
		i.Remove(-3)
		require.Equal(4, i.val)
	})

	t.Run("min", func(t *testing.T) {
		i.val = -4
		i.Remove(3)
		require.Equal(-4, i.val)
	})
}

func TestIncrementerSetMin(t *testing.T) {
	require := require.New(t)
	i := Incrementer{min: -4, max: 4}

	t.Run("maximum", func(t *testing.T) {
		i.SetMin(4)
		require.Equal(3, i.min)
	})

	t.Run("set", func(t *testing.T) {
		i.SetMin(1)
		require.Equal(1, i.min)
	})

	t.Run("lower val", func(t *testing.T) {
		i.val = -4
		i.SetMin(-3)
		require.Equal(-3, i.min)
		require.Equal(-3, i.val)
	})
}

func TestIncrementerSetMax(t *testing.T) {
	require := require.New(t)
	i := Incrementer{max: 4}

	t.Run("minimum", func(t *testing.T) {
		i.SetMax(0)
		require.Equal(1, i.max)
	})

	t.Run("set", func(t *testing.T) {
		i.SetMax(6)
		require.Equal(6, i.max)
	})

	t.Run("higher val", func(t *testing.T) {
		i.val = 4
		i.SetMax(3)
		require.Equal(3, i.max)
		require.Equal(3, i.val)
	})
}

func TestIncrementerSetValue(t *testing.T) {
	require := require.New(t)
	i := Incrementer{min: -4, max: 4}

	t.Run("valid", func(t *testing.T) {
		i.SetValue(3)
		require.Equal(3, i.val)
	})

	t.Run("too low", func(t *testing.T) {
		i.SetValue(-5)
		require.Equal(-4, i.val)
	})

	t.Run("too high", func(t *testing.T) {
		i.SetValue(6)
		require.Equal(4, i.val)
	})
}

func TestIncrementerSetIncrementer(t *testing.T) {
	require := require.New(t)
	i := Incrementer{max: 4}

	t.Run("positive", func(t *testing.T) {
		i.SetIncrementer(1)
		require.Equal(1, i.inc)
	})

	t.Run("negative", func(t *testing.T) {
		i.SetIncrementer(-1)
		require.Equal(-1, i.inc)
	})
}

func TestIncrementerSetOriginalValue(t *testing.T) {
	require := require.New(t)
	i := Incrementer{min: -4, max: 4}

	t.Run("valid", func(t *testing.T) {
		i.SetOriginalValue(3)
		require.Equal(3, i.orig)
	})

	t.Run("too low", func(t *testing.T) {
		i.SetOriginalValue(-5)
		require.Equal(-4, i.orig)
	})

	t.Run("too high", func(t *testing.T) {
		i.SetOriginalValue(6)
		require.Equal(4, i.orig)
	})
}

func TestIncrementerFill(t *testing.T) {
	require := require.New(t)
	i := Incrementer{max: 4, val: 0}

	i.Fill()
	require.Equal(4, i.val)
}

func TestIncrementerEmpty(t *testing.T) {
	require := require.New(t)
	i := Incrementer{max: 4, val: 3}

	i.Empty()
	require.Equal(0, i.val)
}

func TestIncrementerFloor(t *testing.T) {
	require := require.New(t)
	i := Incrementer{min: -4, max: 4, val: 3}

	i.Floor()
	require.Equal(-4, i.val)
}

func TestIncrementerReset(t *testing.T) {
	require := require.New(t)
	i := Incrementer{max: 4, val: 3, orig: 2}

	i.val = 4
	i.Reset()
	require.Equal(2, i.val)
}

func TestIncrementerString(t *testing.T) {
	require := require.New(t)
	i := Incrementer{max: 4}

	t.Run("empty", func(t *testing.T) {
		require.Equal("0/4", i.String())
	})

	t.Run("with val", func(t *testing.T) {
		i.val = 3
		require.Equal("3/4", i.String())
	})
}

func TestIncrementerMarshalJSON(t *testing.T) {
	require := require.New(t)
	i := Incrementer{max: 4}

	t.Run("empty", func(t *testing.T) {
		data, err := i.MarshalJSON()
		require.NoError(err, "Incrementer.MarshalJSON() returned an error: %s", err)
		require.Equal(`{"min":0,"max":4,"val":0,"inc":0,"orig":0}`, string(data))
	})

	t.Run("all set", func(t *testing.T) {
		i = Incrementer{min: -4, max: 4, val: 3, inc: 1, orig: 2}
		i.val = 3
		data, err := i.MarshalJSON()
		require.NoError(err, "Incrementer.MarshalJSON() returned an error: %s", err)
		require.Equal(`{"min":-4,"max":4,"val":3,"inc":1,"orig":2}`, string(data))
	})
}

func TestIncrementerUnmarshalJSON(t *testing.T) {
	require := require.New(t)
	i := Incrementer{}

	t.Run("scan error", func(t *testing.T) {
		err := i.UnmarshalJSON([]byte(`{"max":4,"val":0,"inc":0}`))
		require.Error(err, "Incrementer.UnmarshalJSON() did not return an error")
		require.Equal("input does not match format", err.Error())
	})

	t.Run("invalid max", func(t *testing.T) {
		err := i.UnmarshalJSON([]byte(`{"min":0,"max":0,"val":0,"inc":0,"orig":0}`))
		require.Error(err, "Incrementer.UnmarshalJSON() did not return an error")
		require.Equal("invalid Incrementer: max must be greater than min", err.Error())
	})

	t.Run("higher val", func(t *testing.T) {
		err := i.UnmarshalJSON([]byte(`{"min":0,"max":4,"val":5,"inc":1,"orig":1}`))
		require.Error(err, "Incrementer.UnmarshalJSON() did not return an error")
		require.Equal("invalid Incrementer: val must be less than or equal to max", err.Error())
	})

	t.Run("lower val", func(t *testing.T) {
		err := i.UnmarshalJSON([]byte(`{"min":0,"max":4,"val":-1,"inc":1,"orig":1}`))
		require.Error(err, "Incrementer.UnmarshalJSON() did not return an error")
		require.Equal("invalid Incrementer: val must be greater than or equal to min", err.Error())
	})

	t.Run("valid", func(t *testing.T) {
		err := i.UnmarshalJSON([]byte(`{"min":-4,"max":4,"val":3,"inc":1,"orig":1}`))
		require.NoError(err, "Incrementer.UnmarshalJSON() returned an error: %s", err)
		require.Equal(-4, i.min)
		require.Equal(4, i.max)
		require.Equal(3, i.val)
		require.Equal(1, i.inc)
		require.Equal(1, i.orig)
	})
}

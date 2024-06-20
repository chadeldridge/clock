package incrementers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIncrementerNewIncrementer(t *testing.T) {
	require := require.New(t)
	i := NewIncrementer()

	require.Equal(0, i.val)
}

func TestIncrementerNewIncrementerWithValue(t *testing.T) {
	require := require.New(t)
	i := NewIncrementerWithValue(3)

	require.Equal(3, i.val)
}

func TestIncrementerNewIncrementerFromJSON(t *testing.T) {
	require := require.New(t)
	i, err := NewIncrementerFromJSON([]byte(`{"inc":1,"val":2,"orig":3}`))

	require.NoError(err, "Incrementer.NewFromJSON() returned an error: %s", err)
	require.Equal(1, i.inc)
	require.Equal(2, i.val)
	require.Equal(3, i.orig)
}

func TestIncrementerValue(t *testing.T) {
	require := require.New(t)
	i := Incrementer{}

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
	i := Incrementer{}

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
	i := Incrementer{}

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

func TestIncrementerIsZero(t *testing.T) {
	require := require.New(t)
	i := Incrementer{}

	t.Run("zero", func(t *testing.T) {
		require.True(i.IsEmpty())
	})

	t.Run("positive", func(t *testing.T) {
		i.val = 4
		require.False(i.IsEmpty())
	})

	t.Run("negative", func(t *testing.T) {
		i.val = -4
		require.False(i.IsEmpty())
	})
}

func TestIncrementerIsUnchanged(t *testing.T) {
	require := require.New(t)
	i := Incrementer{}

	t.Run("unchanged", func(t *testing.T) {
		i.orig = 4
		i.val = 4
		require.True(i.IsUnchanged())
	})

	t.Run("changed", func(t *testing.T) {
		i.orig = 4
		i.val = 3
		require.False(i.IsUnchanged())
	})
}

func TestIncrementerIncrement(t *testing.T) {
	require := require.New(t)
	i := Incrementer{}

	t.Run("zero", func(t *testing.T) {
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
	i := Incrementer{}

	t.Run("zero", func(t *testing.T) {
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
	i := Incrementer{}

	t.Run("zero", func(t *testing.T) {
		i.val = 4
		i.Add(0)
		require.Equal(4, i.val)
	})

	t.Run("positive", func(t *testing.T) {
		i.val = 0
		i.Add(4)
		require.Equal(4, i.val)
	})

	t.Run("positive", func(t *testing.T) {
		i.val = 4
		i.Add(-4)
		require.Equal(0, i.val)
	})
}

func TestIncrementerRemove(t *testing.T) {
	require := require.New(t)
	i := Incrementer{}

	t.Run("zero", func(t *testing.T) {
		i.val = 4
		i.Remove(0)
		require.Equal(4, i.val)
	})

	t.Run("positive", func(t *testing.T) {
		i.val = 4
		i.Remove(4)
		require.Equal(0, i.val)
	})

	t.Run("negative", func(t *testing.T) {
		i.val = -4
		i.Remove(-4)
		require.Equal(0, i.val)
	})
}

func TestIncrementerSetValue(t *testing.T) {
	require := require.New(t)
	i := Incrementer{}

	t.Run("positive", func(t *testing.T) {
		i.SetValue(4)
		require.Equal(4, i.val)
	})

	t.Run("negative", func(t *testing.T) {
		i.SetValue(-4)
		require.Equal(-4, i.val)
	})

	t.Run("zero", func(t *testing.T) {
		i.SetValue(0)
		require.Equal(0, i.val)
	})
}

func TestIncrementerSetIncrementer(t *testing.T) {
	require := require.New(t)
	i := Incrementer{}

	t.Run("default", func(t *testing.T) {
		require.Equal(0, i.inc)
	})

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
	i := Incrementer{}

	t.Run("positive", func(t *testing.T) {
		i.SetOriginalValue(4)
		require.Equal(4, i.orig)
	})

	t.Run("negative", func(t *testing.T) {
		i.SetOriginalValue(-4)
		require.Equal(-4, i.orig)
	})

	t.Run("zero", func(t *testing.T) {
		i.SetOriginalValue(0)
		require.Equal(0, i.orig)
	})
}

func TestIncrementerEmpty(t *testing.T) {
	require := require.New(t)
	i := Incrementer{val: 4}

	require.Equal(4, i.val)
	i.Empty()
	require.Equal(0, i.val)
}

func TestIncrementerReset(t *testing.T) {
	require := require.New(t)
	i := Incrementer{val: 4, orig: 2}

	require.Equal(4, i.val)
	i.Reset()
	require.Equal(2, i.val)
}

func TestIncrementerString(t *testing.T) {
	require := require.New(t)
	i := Incrementer{}

	t.Run("default", func(t *testing.T) {
		require.Equal("0", i.String())
	})

	t.Run("positive", func(t *testing.T) {
		i.val = 4
		require.Equal("4", i.String())
	})

	t.Run("negative", func(t *testing.T) {
		i.val = -4
		require.Equal("-4", i.String())
	})
}

func TestIncrementerMarshalJSON(t *testing.T) {
	require := require.New(t)
	i := Incrementer{}

	t.Run("empty", func(t *testing.T) {
		data, err := i.MarshalJSON()
		require.NoError(err, "Incrementer.MarshalJSON() returned an error: %s", err)
		require.Equal(`{"inc":0,"val":0,"orig":0}`, string(data))
	})

	t.Run("set", func(t *testing.T) {
		i = Incrementer{inc: 1, val: 2, orig: 3}
		data, err := i.MarshalJSON()
		require.NoError(err, "Incrementer.MarshalJSON() returned an error: %s", err)
		require.Equal(`{"inc":1,"val":2,"orig":3}`, string(data))
	})
}

func TestIncrementerUnmarshalJSON(t *testing.T) {
	require := require.New(t)
	i := Incrementer{}

	t.Run("nil", func(t *testing.T) {
		err := i.UnmarshalJSON(nil)
		require.Error(err, "Incrementer.UnmarshalJSON() did not return an error")
		require.Equal("Incrementer.UnmarshalJSON(): data was nil", err.Error())
	})

	t.Run("null", func(t *testing.T) {
		err := i.UnmarshalJSON([]byte(`null`))
		require.NoError(err, "Incrementer.UnmarshalJSON() returned an error: %w", err)
	})

	t.Run("empty quotes", func(t *testing.T) {
		err := i.UnmarshalJSON([]byte(`""`))
		require.NoError(err, "Incrementer.UnmarshalJSON() returned an error: %w", err)
	})

	t.Run("scan error", func(t *testing.T) {
		err := i.UnmarshalJSON([]byte(`{"inc":0,"val":0}`))
		require.Error(err, "Incrementer.UnmarshalJSON() did not return an error")
		require.Equal("input does not match format", err.Error())
	})

	t.Run("valid", func(t *testing.T) {
		err := i.UnmarshalJSON([]byte(`{"inc":1,"val":2,"orig":3}`))
		require.NoError(err, "Incrementer.UnmarshalJSON() returned an error: %s", err)
		require.Equal(1, i.inc)
		require.Equal(2, i.val)
		require.Equal(3, i.orig)
	})
}

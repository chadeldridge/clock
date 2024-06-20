package incrementers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUIncrementerNew(t *testing.T) {
	require := require.New(t)
	u := NewUIncrementer()
	require.Equal(0, u.Value())
	require.Equal(1, u.Inc())
	require.Equal(0, u.Original())
}

func TestUIncrementerNewWithValue(t *testing.T) {
	require := require.New(t)

	t.Run("zero", func(t *testing.T) {
		u := NewUIncrementerWithValue(0)
		require.Equal(1, u.inc)
		require.Equal(0, u.val)
		require.Equal(0, u.orig)
	})

	t.Run("positive", func(t *testing.T) {
		u := NewUIncrementerWithValue(3)
		require.Equal(1, u.inc)
		require.Equal(3, u.val)
		require.Equal(3, u.orig)
	})

	t.Run("negative", func(t *testing.T) {
		u := NewUIncrementerWithValue(-3)
		require.Equal(1, u.inc)
		require.Equal(0, u.val)
		require.Equal(0, u.orig)
	})
}

func TestUIncrementerNewFromJSON(t *testing.T) {
	require := require.New(t)
	u, err := NewUIncrementerFromJSON([]byte(`{"inc":1,"val":2,"orig":3}`))
	require.NoError(err, "UIncrementer.NewFromJSON() returned an error: %s", err)
	require.Equal(1, u.inc)
	require.Equal(2, u.val)
	require.Equal(3, u.orig)
}

func TestUIncrementerIncrement(t *testing.T) {
	require := require.New(t)
	u := UIncrementer{}

	t.Run("zero", func(t *testing.T) {
		u.val = 4
		u.Increment()
		require.Equal(4, u.val)
	})

	t.Run("positive", func(t *testing.T) {
		u.val = 0
		u.inc = 1
		u.Increment()
		require.Equal(1, u.val)
	})

	t.Run("negative", func(t *testing.T) {
		u.val = 4
		u.inc = -1
		u.Increment()
		require.Equal(3, u.val)
	})
}

func TestUIncrementerDecrement(t *testing.T) {
	require := require.New(t)
	u := UIncrementer{Incrementer{}}

	t.Run("zero", func(t *testing.T) {
		u.val = 4
		u.inc = 0
		u.Decrement()
		require.Equal(4, u.val)
	})

	t.Run("positive", func(t *testing.T) {
		u.val = 4
		u.inc = 1
		u.Decrement()
		require.Equal(3, u.val)
	})

	t.Run("negative", func(t *testing.T) {
		u.val = 0
		u.inc = -1
		u.Decrement()
		require.Equal(1, u.val)
	})
}

func TestUIncrementerAdd(t *testing.T) {
	require := require.New(t)
	u := UIncrementer{Incrementer{}}

	t.Run("zero", func(t *testing.T) {
		u.val = 4
		u.Add(0)
		require.Equal(4, u.val)
	})

	t.Run("positive", func(t *testing.T) {
		u.val = 0
		u.Add(4)
		require.Equal(4, u.val)
	})

	t.Run("positive", func(t *testing.T) {
		u.val = 4
		u.Add(-4)
		require.Equal(0, u.val)
	})
}

func TestUIncrementerRemove(t *testing.T) {
	require := require.New(t)
	u := UIncrementer{Incrementer{}}

	t.Run("zero", func(t *testing.T) {
		u.val = 4
		u.Remove(0)
		require.Equal(4, u.val)
	})

	t.Run("positive", func(t *testing.T) {
		u.val = 4
		u.Remove(4)
		require.Equal(0, u.val)
	})

	t.Run("negative", func(t *testing.T) {
		u.val = 0
		u.Remove(-4)
		require.Equal(4, u.val)
	})

	t.Run("from zero", func(t *testing.T) {
		u.val = 0
		u.Remove(4)
		require.Equal(0, u.val)
	})
}

func TestUIncrementerSetIncrementer(t *testing.T) {
	require := require.New(t)
	u := UIncrementer{Incrementer{}}

	t.Run("default", func(t *testing.T) {
		require.Equal(0, u.inc)
	})

	t.Run("positive", func(t *testing.T) {
		u.SetIncrementer(1)
		require.Equal(1, u.inc)
	})

	t.Run("negative", func(t *testing.T) {
		u.SetIncrementer(-1)
		require.Equal(-1, u.inc)
	})
}

func TestUIncrementerSetValue(t *testing.T) {
	require := require.New(t)
	u := UIncrementer{Incrementer{}}

	t.Run("positive", func(t *testing.T) {
		u.SetValue(4)
		require.Equal(4, u.val)
	})

	t.Run("negative", func(t *testing.T) {
		u.SetValue(-4)
		require.Equal(0, u.val)
	})

	t.Run("zero", func(t *testing.T) {
		u.SetValue(0)
		require.Equal(0, u.val)
	})
}

func TestUIncrementerSetOriginalValue(t *testing.T) {
	require := require.New(t)
	u := UIncrementer{Incrementer{}}

	t.Run("positive", func(t *testing.T) {
		u.SetOriginalValue(4)
		require.Equal(4, u.orig)
	})

	t.Run("negative", func(t *testing.T) {
		u.SetOriginalValue(-4)
		require.Equal(0, u.orig)
	})

	t.Run("zero", func(t *testing.T) {
		u.SetOriginalValue(0)
		require.Equal(0, u.orig)
	})
}

func TestUIncrementerEmpty(t *testing.T) {
	require := require.New(t)
	u := Incrementer{val: 4}

	require.Equal(4, u.val)
	u.Empty()
	require.Equal(0, u.val)
}

func TestUIncrementerReset(t *testing.T) {
	require := require.New(t)
	u := Incrementer{val: 4, orig: 2}

	require.Equal(4, u.val)
	u.Reset()
	require.Equal(2, u.val)
}

func TestUIncrementerString(t *testing.T) {
	require := require.New(t)
	u := UIncrementer{Incrementer{}}

	t.Run("default", func(t *testing.T) {
		require.Equal("0", u.String())
	})

	t.Run("positive", func(t *testing.T) {
		u.val = 4
		require.Equal("4", u.String())
	})

	t.Run("negative", func(t *testing.T) {
		u.val = -4
		require.Equal("-4", u.String())
	})
}

func TestUIncrementerMarshalJSON(t *testing.T) {
	require := require.New(t)
	u := UIncrementer{Incrementer{}}

	t.Run("empty", func(t *testing.T) {
		data, err := u.MarshalJSON()
		require.NoError(err, "Incrementer.MarshalJSON() returned an error: %s", err)
		require.Equal(`{"inc":0,"val":0,"orig":0}`, string(data))
	})

	t.Run("set", func(t *testing.T) {
		u = UIncrementer{Incrementer{inc: 1, val: 2, orig: 3}}
		data, err := u.MarshalJSON()
		require.NoError(err, "Incrementer.MarshalJSON() returned an error: %s", err)
		require.Equal(`{"inc":1,"val":2,"orig":3}`, string(data))
	})
}

func TestUIncrementerUnmarshalJSON(t *testing.T) {
	require := require.New(t)

	t.Run("valid", func(t *testing.T) {
		u := UIncrementer{}
		err := u.UnmarshalJSON([]byte(`{"inc":1,"val":2,"orig":3}`))
		require.NoError(err, "UIncrementer.UnmarshalJSON() returned an error: %s", err)
		require.Equal(1, u.inc)
		require.Equal(2, u.val)
		require.Equal(3, u.orig)
	})

	t.Run("invalid incrementer", func(t *testing.T) {
		u := UIncrementer{}
		// Make "inc" a string to force error.
		err := u.UnmarshalJSON([]byte(`{"inc":b,"val":2,"orig":3}`))
		require.Error(err, "UIncrementer.UnmarshalJSON() did not return error")
		require.Equal("expected integer", err.Error())
	})

	t.Run("invalid val", func(t *testing.T) {
		u := UIncrementer{}
		err := u.UnmarshalJSON([]byte(`{"inc":1,"val":-2,"orig":3}`))
		require.Error(err, "UIncrementer.UnmarshalJSON() did not return an error")
		require.Equal("invalid UIncrementer: Incrementer.val must be 0 or greater", err.Error())
	})

	t.Run("invalid orig", func(t *testing.T) {
		u := UIncrementer{}
		err := u.UnmarshalJSON([]byte(`{"inc":1,"val":2,"orig":-3}`))
		require.Error(err, "UIncrementer.UnmarshalJSON() did not return an error")
		require.Equal("invalid UIncrementer: Incrementer.orig must be 0 or greater", err.Error())
	})
}

/*
func TestUIncrementer(t *testing.T) {
	require := require.New(t)
	u := NewUIncrementer()

	require.True(c.IsEmpty())

	u.Increment()
	require.Equal(1, u.Value())

	u.SetIncrementer(3)
	u.Increment()
	require.Equal(4, u.Value())

	u.Decrement()
	require.Equal(1, u.Value())

	u.Empty()
	require.Equal(0, u.Value())

	u.Decrement()
	require.Equal(0, u.Value())
}
*/

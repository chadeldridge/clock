package clock

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClockNew(t *testing.T) {
	require := require.New(t)
	c := New(4)

	require.Equal(4, c.Steps)
	require.Equal(0, c.ticks)
}

func TestClockNewWithTicks(t *testing.T) {
	require := require.New(t)
	c := NewWithTicks(4, 3)

	require.Equal(4, c.Steps)
	require.Equal(3, c.ticks)
}

func TestClockNewFromJSON(t *testing.T) {
	require := require.New(t)
	c, err := NewFromJSON([]byte(`{"steps":4,"ticks":3}`))

	require.NoError(err, "Clock.NewFromJSON() returned an error: %s", err)
	require.Equal(4, c.Steps)
	require.Equal(3, c.ticks)
}

func TestClockTicks(t *testing.T) {
	require := require.New(t)
	c := Clock{Steps: 4}

	t.Run("default", func(t *testing.T) {
		require.Equal(0, c.Ticks())
	})

	t.Run("changed", func(t *testing.T) {
		c.ticks = 4
		require.Equal(4, c.Ticks())
	})
}

func TestClockIsFull(t *testing.T) {
	require := require.New(t)
	c := Clock{Steps: 4}

	t.Run("empty", func(t *testing.T) {
		require.False(c.IsFull())
	})

	t.Run("full", func(t *testing.T) {
		c.ticks = 4
		require.True(c.IsFull())
	})
}

func TestClockIsEmpty(t *testing.T) {
	require := require.New(t)
	c := Clock{Steps: 4}

	t.Run("empty", func(t *testing.T) {
		require.True(c.IsEmpty())
	})

	t.Run("full", func(t *testing.T) {
		c.ticks = 4
		require.False(c.IsEmpty())
	})
}

func TestClockTick(t *testing.T) {
	require := require.New(t)
	c := Clock{Steps: 4}

	t.Run("empty", func(t *testing.T) {
		c.Tick()
		require.Equal(1, c.ticks)
	})

	t.Run("full", func(t *testing.T) {
		c.ticks = 4
		c.Tick()
		require.Equal(4, c.ticks)
	})
}

func TestClockAdd(t *testing.T) {
	require := require.New(t)
	c := Clock{Steps: 4}

	t.Run("empty", func(t *testing.T) {
		c.Add(3)
		require.Equal(3, c.ticks)
	})

	t.Run("fill", func(t *testing.T) {
		c.ticks = 3
		c.Add(3)
		require.Equal(4, c.ticks)
	})

	t.Run("full", func(t *testing.T) {
		c.ticks = 4
		c.Add(3)
		require.Equal(4, c.ticks)
	})
}

func TestClockRemove(t *testing.T) {
	require := require.New(t)
	c := Clock{Steps: 4}

	t.Run("while empty", func(t *testing.T) {
		c.Remove(3)
		require.Equal(0, c.ticks)
	})

	t.Run("remove", func(t *testing.T) {
		c.ticks = 3
		c.Remove(2)
		require.Equal(1, c.ticks)
	})

	t.Run("empty", func(t *testing.T) {
		c.ticks = 3
		c.Remove(3)
		require.Equal(0, c.ticks)
	})
}

func TestClockSetSteps(t *testing.T) {
	require := require.New(t)
	c := Clock{Steps: 4}

	t.Run("minimum", func(t *testing.T) {
		c.SetSteps(0)
		require.Equal(1, c.Steps)
	})

	t.Run("set", func(t *testing.T) {
		c.SetSteps(6)
		require.Equal(6, c.Steps)
	})

	t.Run("with ticks", func(t *testing.T) {
		c.ticks = 4
		c.SetSteps(3)
		require.Equal(3, c.Steps)
		require.Equal(3, c.ticks)
	})
}

func TestClockAddSteps(t *testing.T) {
	require := require.New(t)
	c := Clock{Steps: 4}

	c.AddSteps(3)
	require.Equal(7, c.Steps)
}

func TestClockRemoveSteps(t *testing.T) {
	require := require.New(t)
	c := Clock{Steps: 6}

	t.Run("remove", func(t *testing.T) {
		c.RemoveSteps(2)
		require.Equal(4, c.Steps)
	})

	t.Run("minimum", func(t *testing.T) {
		c.Steps = 4
		c.RemoveSteps(6)
		require.Equal(1, c.Steps)
	})

	t.Run("with ticks", func(t *testing.T) {
		c.Steps = 6
		c.ticks = 6
		c.RemoveSteps(2)
		require.Equal(4, c.Steps)
		require.Equal(4, c.ticks)
	})
}

func TestClockFill(t *testing.T) {
	require := require.New(t)
	c := Clock{Steps: 4}

	c.Fill()
	require.Equal(4, c.ticks)
}

func TestClockEmpty(t *testing.T) {
	require := require.New(t)
	c := Clock{Steps: 4}

	c.Empty()
	require.Equal(0, c.ticks)
}

func TestClockReset(t *testing.T) {
	require := require.New(t)
	c := Clock{Steps: 4}

	c.ticks = 4
	c.Reset()
	require.Equal(0, c.ticks)
}

func TestClockString(t *testing.T) {
	require := require.New(t)
	c := Clock{Steps: 4}

	t.Run("empty", func(t *testing.T) {
		require.Equal("0/4", c.String())
	})

	t.Run("with ticks", func(t *testing.T) {
		c.ticks = 3
		require.Equal("3/4", c.String())
	})
}

func TestClockMarchalJSON(t *testing.T) {
	require := require.New(t)
	c := Clock{Steps: 4}

	t.Run("empty", func(t *testing.T) {
		data, err := c.MarshalJSON()
		require.NoError(err, "Clock.MarshalJSON() returned an error: %s", err)
		require.Equal(`{"steps":4,"ticks":0}`, string(data))
	})

	t.Run("with ticks", func(t *testing.T) {
		c.ticks = 3
		data, err := c.MarshalJSON()
		require.NoError(err, "Clock.MarshalJSON() returned an error: %s", err)
		require.Equal(`{"steps":4,"ticks":3}`, string(data))
	})
}

func TestClockUnmarshalJSON(t *testing.T) {
	require := require.New(t)
	c := Clock{}

	t.Run("invalid steps", func(t *testing.T) {
		err := c.UnmarshalJSON([]byte(`{"steps":0,"ticks":0}`))
		require.Error(err, "Clock.UnmarshalJSON() did not return an error")
		require.Equal("invalid clock: steps must be greater than 0", err.Error())
	})

	t.Run("invalid ticks", func(t *testing.T) {
		err := c.UnmarshalJSON([]byte(`{"steps":4,"ticks":-1}`))
		require.Error(err, "Clock.UnmarshalJSON() did not return an error")
		require.Equal("invalid clock: ticks must be greater than or equal to 0", err.Error())
	})

	t.Run("too many ticks", func(t *testing.T) {
		err := c.UnmarshalJSON([]byte(`{"steps":4,"ticks":5}`))
		require.Error(err, "Clock.UnmarshalJSON() did not return an error")
		require.Equal("invalid clock: ticks must be less than or equal to steps", err.Error())
	})

	t.Run("valid", func(t *testing.T) {
		err := c.UnmarshalJSON([]byte(`{"steps":4,"ticks":3}`))
		require.NoError(err, "Clock.UnmarshalJSON() returned an error: %s", err)
		require.Equal(4, c.Steps)
		require.Equal(3, c.ticks)
	})
}

package rpgtools

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiceNewDiceResults(t *testing.T) {
	require := require.New(t)
	r := NewDieResults(D6)
	require.Equal(D6, r.Die)
	require.Equal(0, r.Highest)
	require.Equal(6, r.Lowest)
	require.Empty(r.All)
}

func TestDiceRoll(t *testing.T) {
	require := require.New(t)
	for _, d := range []Die{D2, D4, D6, D8, D10, D12, D20, D100} {
		r := Roll(d)
		require.GreaterOrEqual(r, 1)
		require.LessOrEqual(r, int(d))
	}
}

func TestDiceDieRoll(t *testing.T) {
	require := require.New(t)
	for _, d := range []Die{D2, D4, D6, D8, D10, D12, D20, D100} {
		r := d.Roll(4)
		require.Len(r, 4)
		for _, i := range r {
			require.GreaterOrEqual(i, 1)
			require.LessOrEqual(i, int(d))
		}
	}
}

func TestDiceNewDicePool(t *testing.T) {
	require := require.New(t)
	p := NewDicePool(4)
	require.Equal(4, int(p))
}

func TestDiceDicePoolRoll(t *testing.T) {
	require := require.New(t)
	p := NewDicePool(4)
	r := p.Roll(D6)
	require.Len(r.All, 4)

	require.GreaterOrEqual(r.Highest, 1)
	require.LessOrEqual(r.Highest, 6)
	require.Contains(r.All, r.Highest)

	require.GreaterOrEqual(r.Lowest, 1)
	require.LessOrEqual(r.Lowest, 6)
	require.Contains(r.All, r.Lowest)

	require.GreaterOrEqual(r.Sum, 4)
	require.LessOrEqual(r.Sum, 24)

	for _, i := range r.All {
		require.GreaterOrEqual(i, 1)
		require.LessOrEqual(i, 6)
	}
}

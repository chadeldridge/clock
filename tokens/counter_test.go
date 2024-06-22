package tokens

import (
	"log"
	"testing"

	"github.com/chadeldridge/rpgtools/vectors"
	"github.com/stretchr/testify/require"
)

func TestCounterTokenNewCounterToken(t *testing.T) {
	require := require.New(t)
	c := NewCounterToken()
	require.Len(c.Digits, 10)
}

func TestCounterToken(t *testing.T) {
	require := require.New(t)
	c := NewCounterToken()
	c.Digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	c.Position = vectors.NewVector3(1, 2, 3)
	c.SetBackground("/usr/${userid}/share/tokens/counters/${theme}/background.png")
	require.Equal("/usr/${userid}/share/tokens/counters/${theme}/background.png", c.Background)

	c.Counter.SetValue(1359487620)
	log.Println(c.Value())
	log.Println(c.GetValue())
}

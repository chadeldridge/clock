package tokens

import (
	"strconv"

	"github.com/chadeldridge/rpgtools/incrementers"
	"github.com/chadeldridge/rpgtools/vectors"
)

var digitsPlace = []string{
	"ones",
	"tens",
	"hundreds",
	"thousands",
	"ten-thousands",
	"hundred-thousands",
	"millions",
	"ten-millions",
	"hundred-millions",
	"billions",
}

type CounterToken struct {
	incrementers.Counter                 // Value
	Position             vectors.Vector3 // Position
	Background           string          // "/usr/${userid}/share/tokens/counters/${theme}/background.png"
	Digits               []string        // Digits[0] = "/usr/${userid}/share/tokens/counters/${theme}/0.png"
}

func NewCounterToken() CounterToken {
	return CounterToken{
		Counter:  incrementers.NewCounter(),
		Position: vectors.Vector3{},
		Digits:   make([]string, 10),
	}
}

func (c *CounterToken) SetBackground(background string) {
	c.Background = background
}

func (c *CounterToken) GetValue() map[string]string {
	v := make(map[string]string)
	i := 0
	s := strconv.Itoa(c.Value())

	for d := len(s) - 1; d >= 0; d-- {
		j, _ := strconv.Atoi(string(s[d]))
		v[digitsPlace[i]] = c.Digits[j]
		i++
	}

	return v
}

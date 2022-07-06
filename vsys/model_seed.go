package vsys

import (
	"crypto/rand"
	"fmt"
	"strings"
)

type Seed struct {
	Str
}

func NewSeed(s string) (*Seed, error) {
	words := strings.Fields(s)
	if len(words) != 15 {
		return nil, fmt.Errorf("NewSeed: Seed must contain 15 words")
	}

	for _, w := range words {
		if !WORDS_SET.Contains(w) {
			return nil, fmt.Errorf("NewSeed: Seed must contain valid words")
		}
	}

	return &Seed{Str(s)}, nil
}

func NewRandSeed() (*Seed, error) {
	s, err := NewRandSeedStr()
	if err != nil {
		return nil, fmt.Errorf("NewRandSeed: %w", err)
	}
	return NewSeed(s)
}

func NewRandSeedStr() (string, error) {
	cnt := 2048
	words := []string{}

	for i := 0; i < 5; i++ {
		r := make([]byte, 4)
		_, err := rand.Read(r)

		if err != nil {
			return "", fmt.Errorf("NewRandSeedStr: %w", err)
		}

		var x int
		x += int(r[3])
		x += int(r[2]) << 8
		x += int(r[1]) << 16
		x += int(r[0]) << 24

		w1 := x % cnt
		w2 := (x/cnt + w1) % cnt
		w3 := (x/cnt/cnt + w2) % cnt

		words = append(words, WORDS[w1])
		words = append(words, WORDS[w2])
		words = append(words, WORDS[w3])
	}
	return strings.Join(words, " "), nil
}

func (s *Seed) String() string {
	return fmt.Sprintf("%T(%s)", s, s.Str)
}

package slot

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSimple_screen(t *testing.T) {
	tests := []struct {
		name   string
		expect screen
	}{
		{"simple", screen{{"Eggs", "Mayonnaise", "Steak"}, {"Eggs", "Cheese", "Bacon"}, {"Sausage", "Bacon", "Steak"}, {"Mayonnaise", "Bacon", "Ham"}, {"Butter", "Steak", "Mayonnaise"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := AtkinsDiet{random: rand.New(rand.NewSource(0)), strips: defAtkinsDietReelStrip()}
			assert.Equal(t, tt.expect, s.screen())
		})
	}
}

func BenchmarkSimple_Spin(b *testing.B) {
	s := AtkinsDiet{random: rand.New(rand.NewSource(time.Now().Unix()))}
	for i := 0; i < b.N; i++ {
		s.Spin("", 1)
	}
}

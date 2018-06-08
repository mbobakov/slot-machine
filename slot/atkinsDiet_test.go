package slot

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSimple_screenFromStops(t *testing.T) {
	tests := []struct {
		name   string
		input  stops
		expect screen
	}{
		{"simple", stops{1, 2, 3, 4, 5}, screen{{"Mayonnaise", "Ham", "Sausage"}, {"Steak", "Sausage", "Cheese"}, {"Scale", "Cheese", "Mayonnaise"}, {"Butter", "Bacon", "Cheese"}, {"Sausage", "Butter", "Bacon"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := AtkinsDiet{random: rand.New(rand.NewSource(0)), strips: defAtkinsDietReelStrip()}
			assert.Equal(t, tt.expect, s.screenFromStops(tt.input))
		})
	}
}

func BenchmarkSimple_Spin(b *testing.B) {
	s := AtkinsDiet{random: rand.New(rand.NewSource(time.Now().Unix()))}
	for i := 0; i < b.N; i++ {
		s.Spin("", 1)
	}
}

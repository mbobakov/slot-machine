package slot

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSimple_screen(t *testing.T) {
	tests := []struct {
		name     string
		reelSize uint8
		expect   screen
	}{
		{"simple", 5, screen{{4, 0, 1}, {4, 0, 1}, {3, 4, 0}, {1, 2, 3}, {0, 1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := AtkinsDiet{random: rand.New(rand.NewSource(0)), reelSize: tt.reelSize}
			assert.Equal(t, tt.expect, s.screen())
		})
	}
}

func BenchmarkSimple_Spin(b *testing.B) {
	s := AtkinsDiet{reelSize: 32, scatter: 3, random: rand.New(rand.NewSource(time.Now().Unix())), wild: 0}
	for i := 0; i < b.N; i++ {
		s.Spin(1)
	}
}

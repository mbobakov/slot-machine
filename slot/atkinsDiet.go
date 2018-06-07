package slot

import (
	"math/rand"
	"sync"
	"time"
)

// AtkinsDiet slot machine
type AtkinsDiet struct {
	paylines []payLine
	strips   strips

	mu          sync.RWMutex
	coefficient coefficient
	random      *rand.Rand
}

// NewAtkinsDiet slot-machine
func NewAtkinsDiet() *AtkinsDiet {
	return &AtkinsDiet{
		random:      rand.New(rand.NewSource(time.Now().Unix())),
		coefficient: defAtkinsDietCoefficients(),
		paylines:    defAtkinsDietPayLines(),
		strips:      defAtkinsDietReelStrip(),
	}
}

func (s *AtkinsDiet) screen() screen {
	var scr screen
	for i := 0; i < 5; i++ {
		// sTops
		s.mu.Lock() // cause random is not thread safe
		stopIdx := s.random.Intn(int(31))
		s.mu.Unlock()

		// midlle and bottom
		for j := 0; j < 3; j++ {
			scr[i][j] = s.strips[i][stopIdx]
			if stopIdx == 31 {
				stopIdx = 0
				continue
			}
			stopIdx++
		}
	}
	return scr
}

// Spin the Simple slot machine. Return bonus,free spins and STops
func (s *AtkinsDiet) Spin(multypl int) (int, bool, [5]string) {
	var (
		freeSpin bool
		bonus    int
	)
	src := s.screen()

	for _, p := range s.paylines {
		if !freeSpin {
			freeSpin = s.hasFreeSpin(src, p)
		}
		bonus += s.bonus(src, p)
	}
	return bonus, freeSpin, [5]string{src[0][0], src[1][0], src[2][0], src[3][0], src[4][0]}
}

package slot

import (
	"math/rand"
	"sync"
	"time"
)

// AtkinsDiet slot machine
type AtkinsDiet struct {
	reelSize,
	wild,
	scatter uint8
	paylines []payLine

	mu          sync.RWMutex
	coefficient coefficient
	random      *rand.Rand
}

// NewAtkinsDiet slot-machine
func NewAtkinsDiet() *AtkinsDiet {
	return &AtkinsDiet{
		random:      rand.New(rand.NewSource(time.Now().Unix())),
		reelSize:    32,
		wild:        0,
		scatter:     15,
		coefficient: defAtkinsDietCoefficients(),
		paylines:    defAtkinsDietPayLines(),
	}
}

func (s *AtkinsDiet) screen() screen {
	var scr screen
	for i := 0; i < 5; i++ {
		// sTops
		s.mu.Lock() // cause random is not thread safe
		scr[i][0] = uint8(s.random.Intn(int(s.reelSize)))
		s.mu.Unlock()
		// midlle and bottom
		for j := 1; j < 3; j++ {
			if scr[i][j-1] == s.reelSize-1 {
				scr[i][j] = 0
				continue
			}
			scr[i][j] = scr[i][j-1] + 1
		}
	}
	return scr
}

// Spin the Simple slot machine. Return bonus,free spins and STops
func (s *AtkinsDiet) Spin(multypl int) (int, bool, [5]uint8) {
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
	return bonus, freeSpin, [5]uint8{src[0][0], src[1][0], src[2][0], src[3][0], src[4][0]}
}

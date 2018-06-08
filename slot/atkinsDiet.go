package slot

import (
	"math/rand"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// AtkinsDiet slot machine
type AtkinsDiet struct {
	paylines []payLine
	strips   strips
	log      logrus.FieldLogger

	mu          sync.RWMutex
	coefficient coefficient
	random      *rand.Rand
}

// NewAtkinsDiet slot-machine
func NewAtkinsDiet(logger logrus.FieldLogger) *AtkinsDiet {
	return &AtkinsDiet{
		random:      rand.New(rand.NewSource(time.Now().Unix())),
		coefficient: defAtkinsDietCoefficients(),
		paylines:    defAtkinsDietPayLines(),
		strips:      defAtkinsDietReelStrip(),
		log:         logger,
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
func (s *AtkinsDiet) Spin(reqID string, multypl int) (int, bool, [5]string) {
	var (
		freeSpin bool
		bonus    int
		l        = s.log.WithField("reqID", reqID)
	)
	src := s.screen()
	l.Debugf("Screen is: %s", src.String())

	for i, p := range s.paylines {
		if !freeSpin {
			freeSpin = s.hasFreeSpin(src, p)
		}
		payout := s.bonus(src, p)
		l.Debugf("Payline #'%d' with elements: '%s-%s-%s-%s-%s' won payout: %d and freespin: %t",
			i, src[0][p[0]], src[1][p[1]], src[2][p[2]], src[3][p[3]], src[4][p[4]], payout, freeSpin)
		bonus += payout
	}
	return bonus, freeSpin, [5]string{src[0][0], src[1][0], src[2][0], src[3][0], src[4][0]}
}

package slot

// bonus for the payline on screen
// Also deal with the wild elements
func (s *AtkinsDiet) bonus(d screen, pl payLine) int {
	var (
		elCount int
		cell    = d[0][pl[1]]
		onWild  = cell == s.wild
	)
	for i := 1; i < 5; i++ {
		col := i
		line := pl[i]
		if onWild && d[col][line] != s.wild {
			onWild = false
			cell = d[col][line]
			elCount++
			continue
		}
		if cell != d[col][line] && s.wild != d[col][line] {
			break
		}
		elCount++
	}
	if elCount <= 1 {
		elCount = s.scatterCount(d, pl) - 1
		cell = s.scatter
	}
	if elCount < 1 {
		return 0
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.coefficient[cell][elCount-1]
}

func (s *AtkinsDiet) hasFreeSpin(d screen, pl payLine) bool {
	return s.scatterCount(d, pl) >= 3
}

func (s *AtkinsDiet) scatterCount(d screen, pl payLine) int {
	var scLongest, scCurrent int
	for i := 0; i < 5; i++ {
		switch {
		case i == 4:
			if d[i][pl[i]] == s.scatter {
				scCurrent++
			}
			if scCurrent > scLongest {
				scLongest = scCurrent
			}
		case d[i][pl[i]] != s.scatter && scCurrent != 0:
			if scCurrent > scLongest {
				scLongest = scCurrent
			}
			scCurrent = 0
		case d[i][pl[i]] == s.scatter:
			scCurrent++
		}
	}
	return scLongest
}

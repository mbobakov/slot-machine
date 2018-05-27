package slot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple_bonus(t *testing.T) {
	tests := []struct {
		name      string
		expect    int
		inputLine payLine
		IntputScr screen
	}{
		{"1X4=1000", 1000, payLine{0, 0, 0, 0}, screen{{1, 0, 0}, {1, 0, 0}, {1, 0, 0}, {1, 0, 0}, {1, 0, 0}}},
		{"1wild", 1000, payLine{0, 0, 0, 0}, screen{{0, 0, 0}, {1, 0, 0}, {1, 0, 0}, {1, 0, 0}, {1, 0, 0}}},
		{"wildInMiddle", 1000, payLine{0, 0, 0, 0}, screen{{1, 0, 0}, {1, 0, 0}, {0, 0, 0}, {1, 0, 0}, {1, 0, 0}}},
		{"noBonus", 0, payLine{0, 0, 0, 0}, screen{{2, 0, 0}, {1, 0, 0}, {1, 0, 0}, {1, 0, 0}, {1, 0, 0}}},
		{"scatter", 300, payLine{0, 0, 0, 0}, screen{{3, 0, 0}, {3, 0, 0}, {3, 0, 0}, {3, 0, 0}, {3, 0, 0}}},
		{"scatterInMiddle", 2, payLine{0, 0, 0, 0}, screen{{33, 0, 0}, {34, 0, 0}, {31, 0, 0}, {3, 0, 0}, {3, 0, 0}}},
		// TODO: Add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := AtkinsDiet{scatter: 3, coefficient: defAtkinsDietCoefficients()}
			assert.Equal(t, tt.expect, s.bonus(tt.IntputScr, tt.inputLine))
		})
	}
}

func TestSimple_scatterCount(t *testing.T) {
	tests := []struct {
		name         string
		expect       int
		inputPayline payLine
		intputScr    screen
	}{
		{"hasAtTheEnd", 3, payLine{0, 0, 0, 0, 0}, screen{{0, 0, 0}, {0, 0, 0}, {3, 0, 0}, {3, 0, 0}, {3, 0, 0}}},
		{"hasAtTheMiddle", 3, payLine{0, 0, 0, 0, 0}, screen{{0, 0, 0}, {3, 0, 0}, {3, 0, 0}, {3, 0, 0}, {0, 0, 0}}},
		{"notFound", 1, payLine{0, 0, 0, 0, 0}, screen{{3, 0, 0}, {0, 0, 0}, {3, 0, 0}, {0, 0, 0}, {3, 0, 0}}},
		{"notFound2", 1, payLine{0, 0, 0, 0, 0}, screen{{0, 0, 0}, {3, 0, 0}, {0, 0, 0}, {3, 0, 0}, {0, 0, 0}}},
		// TODO: Add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := AtkinsDiet{scatter: 3}
			assert.Equal(t, tt.expect, s.scatterCount(tt.intputScr, tt.inputPayline))
		})
	}
}

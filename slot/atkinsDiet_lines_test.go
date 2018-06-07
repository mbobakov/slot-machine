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
		{"1X4=1000", 1000, payLine{0, 0, 0, 0}, screen{{"Steak", "", ""}, {"Steak", "", ""}, {"Steak", "", ""}, {"Steak", "", ""}, {"Steak", "", ""}}},
		{"1wild", 1000, payLine{0, 0, 0, 0}, screen{{"Atkins", "", ""}, {"Steak", "", ""}, {"Steak", "", ""}, {"Steak", "", ""}, {"Steak", "", ""}}},
		{"wildInMiddle", 1000, payLine{0, 0, 0, 0}, screen{{"Steak", "", ""}, {"Steak", "", ""}, {"Atkins", "", ""}, {"Steak", "", ""}, {"Steak", "", ""}}},
		{"noBonus", 0, payLine{0, 0, 0, 0}, screen{{"Ham", "", ""}, {"Steak", "", ""}, {"Steak", "", ""}, {"Steak", "", ""}, {"Steak", "", ""}}},
		// TODO: Add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := AtkinsDiet{coefficient: defAtkinsDietCoefficients()}
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
		{"hasAtTheEnd", 3, payLine{0, 0, 0, 0, 0}, screen{{"", "", ""}, {"", "", ""}, {"Scale", "", ""}, {"Scale", "", ""}, {"Scale", "", ""}}},
		{"hasAtTheMiddle", 3, payLine{0, 0, 0, 0, 0}, screen{{"", "", ""}, {"Scale", "", ""}, {"Scale", "", ""}, {"Scale", "", ""}, {"", "", ""}}},
		{"notFound", 1, payLine{0, 0, 0, 0, 0}, screen{{"Scale", "", ""}, {"", "", ""}, {"Scale", "", ""}, {"", "", ""}, {"Scale", "", ""}}},
		{"notFound2", 1, payLine{0, 0, 0, 0, 0}, screen{{"", "", ""}, {"Scale", "", ""}, {"", "", ""}, {"Scale", "", ""}, {"", "", ""}}},
		// TODO: Add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := AtkinsDiet{}
			assert.Equal(t, tt.expect, s.scatterCount(tt.intputScr, tt.inputPayline))
		})
	}
}

package slot

func defAtkinsDietCoefficients() coefficient {
	return coefficient{
		0: {5, 50, 500, 5000},
		1: {5, 40, 200, 1000},
		2: {2, 30, 150, 500},
		3: {2, 25, 100, 300},
		4: {0, 20, 75, 200},
		5: {0, 20, 75, 200},
		6: {0, 15, 50, 100},
		7: {0, 15, 50, 100},
		8: {0, 10, 25, 50},
		9: {0, 10, 25, 50},
	}
}

func defAtkinsDietPayLines() []payLine {
	return []payLine{
		{1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0},
		{2, 2, 2, 2, 2},
		{0, 1, 2, 1, 0},
		{2, 1, 0, 1, 2},
		{1, 0, 0, 0, 1},
		{1, 2, 2, 2, 1},
		{0, 0, 1, 2, 2},
		{2, 2, 1, 0, 0},
		{1, 0, 1, 2, 1},
		{1, 2, 1, 0, 1},
		{0, 1, 1, 1, 0},
		{2, 1, 1, 1, 2},
		{0, 1, 0, 1, 0},
		{2, 1, 2, 1, 2},
		{1, 1, 0, 1, 1},
		{1, 1, 2, 1, 1},
		{0, 0, 2, 0, 0},
		{2, 2, 0, 2, 2},
		{0, 2, 2, 2, 0},
	}
}
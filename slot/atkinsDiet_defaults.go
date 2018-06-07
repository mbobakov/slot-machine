package slot

func defAtkinsDietCoefficients() coefficient {
	return coefficient{
		"Atkins":       {5, 50, 500, 5000},
		"Steak":        {5, 40, 200, 1000},
		"Ham":          {2, 30, 150, 500},
		"BuffaloWings": {2, 25, 100, 300},
		"Sausage":      {0, 20, 75, 200},
		"Eggs":         {0, 20, 75, 200},
		"Butter":       {0, 15, 50, 100},
		"Cheese":       {0, 15, 50, 100},
		"Bacon":        {0, 10, 25, 50},
		"Mayonnaise":   {0, 10, 25, 50},
	}
}

func defAtkinsDietReelStrip() strips {
	return strips{
		0: {"Scale", "Mayonnaise", "Ham", "Sausage", "Bacon", "Eggs", "Cheese", "Mayonnaise", "Sausage", "Butter", "BuffaloWings", "Bacon", "Eggs", "Mayonnaise", "Steak", "BuffaloWings", "Butter", "Cheese", "Eggs", "Atkins", "Bacon", "Mayonnaise", "Ham", "Cheese", "Eggs", "Scale", "Butter", "Bacon", "Sausage", "BuffaloWings", "Steak", "Butter"},
		1: {"Mayonnaise", "BuffaloWings", "Steak", "Sausage", "Cheese", "Mayonnaise", "Ham", "Butter", "Bacon", "Steak", "Sausage", "Mayonnaise", "Ham", "Atkins", "Butter", "Eggs", "Cheese", "Bacon", "Sausage", "BuffaloWings", "Scale", "Mayonnaise", "Butter", "Cheese", "Bacon", "Eggs", "BuffaloWings", "Mayonnaise", "Steak", "Ham", "Cheese", "Bacon"},
		2: {"Ham", "Butter", "Eggs", "Scale", "Cheese", "Mayonnaise", "Butter", "Ham", "Sausage", "Bacon", "Steak", "BuffaloWings", "Butter", "Mayonnaise", "Cheese", "Sausage", "Eggs", "Bacon", "Mayonnaise", "BuffaloWings", "Ham", "Sausage", "Bacon", "Cheese", "Eggs", "Atkins", "BuffaloWings", "Bacon", "Butter", "Cheese", "Mayonnaise", "Steak"},
		3: {"Ham", "Cheese", "Atkins", "Scale", "Butter", "Bacon", "Cheese", "Sausage", "Steak", "Eggs", "Bacon", "Mayonnaise", "Sausage", "Cheese", "Butter", "Ham", "Mayonnaise", "Bacon", "BuffaloWings", "Sausage", "Cheese", "Eggs", "Butter", "BuffaloWings", "Bacon", "Mayonnaise", "Eggs", "Ham", "Sausage", "Steak", "Mayonnaise", "Bacon"},
		4: {"Bacon", "Scale", "Steak", "Ham", "Cheese", "Sausage", "Butter", "Bacon", "BuffaloWings", "Cheese", "Sausage", "Ham", "Butter", "Steak", "Mayonnaise", "Eggs", "Sausage", "Ham", "Atkins", "Butter", "BuffaloWings", "Mayonnaise", "Eggs", "Ham", "Bacon", "Butter", "Steak", "Mayonnaise", "Sausage", "Eggs", "Cheese", "BuffaloWings"},
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

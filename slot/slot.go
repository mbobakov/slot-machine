package slot

type (
	screen      [5][3]uint8
	coefficient map[uint8][4]int // map[<element>]{bonus for 2, bonus for 3, bonus for 4, bonus for 5}
	payLine     [5]int           // [column]line
)

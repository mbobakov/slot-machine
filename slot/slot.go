package slot

const (
	wild    = "Atkins"
	scatter = "Scale"
)

type (
	screen      [5][3]string
	strips      [5][32]string
	coefficient map[string][4]int // map[<element>]{bonus for 2, bonus for 3, bonus for 4, bonus for 5}
	payLine     [5]int            // [column]line
)

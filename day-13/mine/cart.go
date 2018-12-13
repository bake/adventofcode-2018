package mine

// direction in which a cart can move.
type direction int

func (d direction) String() string {
	dirs := map[direction]byte{
		north: '^', south: 'v',
		west: '<', east: '>',
	}
	return string(dirs[d])
}

const (
	north direction = iota
	south
	west
	east
)

type turn int

func (t turn) String() string {
	turns := map[turn]byte{
		left:     'l',
		right:    'r',
		straight: 's',
	}
	return string(turns[t])
}

const (
	left turn = iota
	straight
	right
)

// Cart has coordinates and a direction in which it moves.
type Cart struct {
	X, Y    int
	Crashed bool
	dir     direction
	turn    turn
}

func (c Cart) String() string {
	if c.Crashed {
		return "X"
	}
	return c.dir.String()
}

package game

type entityKind int

const (
	kindElf entityKind = iota
	kindGoblin
)

type Entity struct {
	kind   entityKind
	x, y   int
	AP, HP int
}

func (e Entity) X() int { return e.x }
func (e Entity) Y() int { return e.y }

func (e Entity) Alive() bool { return e.HP > 0 }

func (e Entity) String() string {
	switch e.kind {
	case kindElf:
		return "E"
	case kindGoblin:
		return "G"
	}
	return "U"
}

type Entities []*Entity

func (es Entities) At(x, y int) *Entity {
	for _, e := range es {
		if e.x == x && e.y == y {
			return e
		}
	}
	return nil
}

func (es Entities) Kind(kind entityKind) Entities {
	var res Entities
	for _, e := range es {
		if e.kind == kind {
			res = append(res, e)
		}
	}
	return res
}

func (es Entities) Alive() Entities {
	var res Entities
	for _, e := range es {
		if e.Alive() {
			res = append(res, e)
		}
	}
	return res
}

func (es Entities) Len() int { return len(es) }
func (es Entities) Less(i, j int) bool {
	return es[i].y < es[j].y || (es[i].y == es[j].y && es[i].x < es[j].x)
}
func (es Entities) Swap(i, j int) { es[i], es[j] = es[j], es[i] }

package grid

// XY holds x and y in an array.
type XY [2]int

// X returns the x value of a XY.
func (xy XY) X() int { return xy[0] }

// Y returns the y value of a XY.
func (xy XY) Y() int { return xy[1] }

// SortByX implements the sort interface.
type SortByX []XY

func (xys SortByX) Len() int           { return len(xys) }
func (xys SortByX) Less(i, j int) bool { return xys[i].X() < xys[j].X() }
func (xys SortByX) Swap(i, j int)      { xys[i], xys[j] = xys[j], xys[i] }

// SortByY implements the sort interface.
type SortByY []XY

func (xys SortByY) Len() int           { return len(xys) }
func (xys SortByY) Less(i, j int) bool { return xys[i].Y() < xys[j].Y() }
func (xys SortByY) Swap(i, j int)      { xys[i], xys[j] = xys[j], xys[i] }

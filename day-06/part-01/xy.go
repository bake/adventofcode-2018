package main

// xy holds x and y in an array.
type xy [2]int

func (xy xy) x() int { return xy[0] }
func (xy xy) y() int { return xy[1] }

type sortByX []xy

func (xys sortByX) Len() int           { return len(xys) }
func (xys sortByX) Less(i, j int) bool { return xys[i].x() < xys[j].x() }
func (xys sortByX) Swap(i, j int)      { xys[i], xys[j] = xys[j], xys[i] }

type sortByY []xy

func (xys sortByY) Len() int           { return len(xys) }
func (xys sortByY) Less(i, j int) bool { return xys[i].y() < xys[j].y() }
func (xys sortByY) Swap(i, j int)      { xys[i], xys[j] = xys[j], xys[i] }

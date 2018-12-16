package game

import (
	"fmt"
	"image"
	"image/color"
)

func (g *Game) String() string {
	var str string
	for y := range g.Grid {
		var entities Entities
		for x := range g.Grid[y] {
			if e := g.Entities.At(x, y); e != nil && e.Alive() {
				entities = append(entities, e)
				str += e.String()
				continue
			}
			str += string(g.Grid[y][x])
		}
		str += "  "
		for _, e := range entities {
			str += fmt.Sprintf(" %s(%d)", e, e.HP)
		}
		str += "\n"
	}
	return str
}

func (g *Game) Image() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, len(g.Grid[0]), len(g.Grid)))
	for y := range g.Grid {
		for x := range g.Grid[y] {
			if e := g.Entities.At(x, y); e != nil && e.Alive() {
				hp := uint8(255 / 200 * e.HP)
				switch e.kind {
				case kindElf:
					img.Set(x, y, color.RGBA{R: 255 - hp, G: hp, A: 255})
				case kindGoblin:
					img.Set(x, y, color.RGBA{R: 255 - hp, B: hp, A: 255})
				}
				continue
			}
			switch g.Grid[y][x] {
			case '#':
				img.Set(x, y, color.Black)
			case '.':
				img.Set(x, y, color.White)
			}
		}
	}
	return img
}

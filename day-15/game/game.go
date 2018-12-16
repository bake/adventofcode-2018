package game

import (
	"math"
	"sort"

	"github.com/bakerolls/adventofcode-2018/day-15/pathfinding"
)

type Game struct {
	turns    int
	Grid     Grid
	Entities Entities
	Path     pathfinding.SearchFunc
}

func New(g [][]byte, p pathfinding.SearchFunc) *Game {
	var entities Entities
	for y, row := range g {
		for x, cell := range row {
			switch cell {
			case 'E':
				entities = append(entities, &Entity{kindElf, x, y, 3, 200})
				g[y][x] = '.'
			case 'G':
				entities = append(entities, &Entity{kindGoblin, x, y, 3, 200})
				g[y][x] = '.'
			}
		}
	}
	return &Game{
		Grid:     Grid(g),
		Entities: entities,
		Path:     p,
	}
}

func (g *Game) Turn() bool {
	entities := g.Entities.Alive()
	sort.Sort(entities)
	for _, entity := range entities {
		if !entity.Alive() {
			continue
		}
		var enemies Entities
		switch entity.kind {
		case kindElf:
			enemies = entities.Kind(kindGoblin).Alive()
		case kindGoblin:
			enemies = entities.Kind(kindElf).Alive()
		}
		if len(enemies) == 0 {
			return false
		}
		if !g.fight(entity, enemies) {
			g.move(entity, enemies)
			g.fight(entity, enemies)
		}
	}
	g.turns++
	return true
}

func (g *Game) move(entity *Entity, enemies Entities) {
	nearestDist := math.MaxInt32
	nearestField := pathfinding.XYer(entity)
	for _, enemy := range enemies {
		path, ok := g.Path(g, enemy, entity)
		if ok && 2 < len(path) && len(path) < nearestDist {
			nearestDist = len(path)
			nearestField = path[1]
		}
	}
	entity.x = nearestField.X()
	entity.y = nearestField.Y()
}

func (g *Game) fight(entity *Entity, enemies Entities) bool {
	var enemy *Entity
	ds := []struct{ x, y int }{{0, -1}, {-1, 0}, {+1, 0}, {0, +1}}
	for _, d := range ds {
		e := enemies.At(entity.x+d.x, entity.y+d.y)
		if e == nil {
			continue
		}
		if enemy == nil || e.HP < enemy.HP {
			enemy = e
		}
	}
	if enemy == nil {
		return false
	}
	enemy.HP -= entity.AP
	return true
}

// Walkable implements the pathfinding.Walker interface.
func (g *Game) Walkable(x, y int) bool {
	if !g.Grid.Walkable(x, y) {
		return false
	}
	return g.Entities.Alive().At(x, y) == nil
}

func (g *Game) Outcome() (turns, hp int) {
	for _, e := range g.Entities.Alive() {
		hp += e.HP
	}
	return g.turns, hp
}

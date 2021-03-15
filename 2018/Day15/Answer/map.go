package day15

type Coordinate struct {
	X, Y int
}

var offsets = []Coordinate{
	{0, -1},
	{-1, 0},
	{1, 0},
	{0, 1},
}

type Map map[int]map[int]*Tile

func (m Map) SetTile(t *Tile, x, y int) {
	if m[y] == nil {
		m[y] = make(map[int]*Tile)
	}
	m[y][x] = t
	t.X = x
	t.Y = y
	t.Map = m
}

func (m Map) Tile(x, y int) *Tile {
	if m[y] == nil {
		return nil
	}
	return m[y][x]
}

func (m Map) FindWalkableTiles(t *Tile) (map[*Tile]int, map[*Tile]*Tile) {
	frontier := []*Tile{t}
	distance := map[*Tile]int{t: 0}
	cameFrom := map[*Tile]*Tile{t: nil}

	for len(frontier) > 0 {
		current := frontier[0]
		frontier = frontier[1:]

		for _, next := range current.WalkableNeighbors() {
			if _, ok := distance[next]; !ok {
				frontier = append(frontier, next)
				distance[next] = distance[current] + 1
				cameFrom[next] = current
			}
		}
	}

	return distance, cameFrom
}

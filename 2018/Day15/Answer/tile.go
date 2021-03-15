package day15

type Tile struct {
	Kind int
	X, Y int
	Map  Map
	Unit *Unit
}

func (t Tile) WalkableNeighbors() []*Tile {
	var neighbors []*Tile

	for _, offset := range offsets {
		if n := t.Map.Tile(t.X+offset.X, t.Y+offset.Y); n != nil && n.Kind == KindSpace {
			neighbors = append(neighbors, n)
		}
	}

	return neighbors
}

type SortableTiles []*Tile

func (s SortableTiles) Len() int {
	return len(s)
}

func (s SortableTiles) Less(i, j int) bool {
	if s[i].Y == s[j].Y {
		return s[i].X < s[j].X
	}
	return s[i].Y < s[j].Y
}

func (s SortableTiles) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

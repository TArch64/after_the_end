package pathfinder

import (
	"cmp"
	"slices"

	"after_the_end/db/model"
	"after_the_end/helper/axial"
)

type OpenedHex struct {
	coord    *axial.Coord
	priority uint
}

type Finder struct {
	Location *model.Location
	from     *axial.Coord
	to       *axial.Coord
	hexMap   map[string]*model.LocationHex
	cameFrom map[string]*axial.Coord
	costs    map[string]uint
	open     []*OpenedHex
	current  *OpenedHex
}

func New(location *model.Location, from, to *axial.Coord) *Finder {
	finder := &Finder{
		Location: location,
		from:     from,
		to:       to,
		hexMap:   make(map[string]*model.LocationHex),
		cameFrom: make(map[string]*axial.Coord),
		costs:    make(map[string]uint),
	}

	finder.buildHexMap()
	return finder
}

func (f *Finder) Find() []*axial.Coord {
	f.costs[f.from.StringKey()] = 0
	f.open = append(f.open, &OpenedHex{coord: f.from, priority: 0})

	for {
		if !f.nextOpened() {
			return nil
		}

		if f.current.coord.StringKey() == f.to.StringKey() {
			return f.buildPath()
		}

		for _, coord := range f.current.coord.Neighbours() {
			if _, ok := f.hexMap[coord.StringKey()]; ok {
				f.handleNeighbour(coord)
			}
		}

		slices.SortFunc(f.open, func(a, b *OpenedHex) int {
			return cmp.Compare(a.priority, b.priority)
		})
	}
}

func (f *Finder) buildHexMap() {
	for _, hex := range f.Location.Hexes {
		f.hexMap[hex.Coord.StringKey()] = hex
	}
}

func (f *Finder) nextOpened() bool {
	if len(f.open) == 0 {
		return false
	}

	f.current = f.open[0]
	f.open = f.open[1:]
	return true
}

func (f *Finder) handleNeighbour(coord *axial.Coord) {
	coordKey := coord.StringKey()
	currentKey := f.current.coord.StringKey()
	cost := f.costs[currentKey] + 1

	if previousCost, visited := f.costs[coordKey]; !visited || cost < previousCost {
		f.costs[coordKey] = cost
		f.cameFrom[coordKey] = f.current.coord
		priority := cost + coord.Distance(f.to)

		if visited {
			f.updateOpenedPriority(coordKey, priority)
		} else {
			f.addOpened(coord, priority)
		}
	}
}

func (f *Finder) updateOpenedPriority(key string, priority uint) {
	for _, opened := range f.open {
		if opened.coord.StringKey() == key {
			opened.priority = priority
			break
		}
	}
}

func (f *Finder) addOpened(coord *axial.Coord, priority uint) {
	f.open = append(f.open, &OpenedHex{
		coord:    coord,
		priority: priority,
	})
}

func (f *Finder) buildPath() []*axial.Coord {
	current := f.to
	path := []*axial.Coord{f.to}

	for {
		if from, ok := f.cameFrom[current.StringKey()]; ok {
			path = append(path, from)
			current = from
		} else {
			break
		}
	}

	slices.Reverse(path)
	return path
}

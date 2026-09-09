package main

func (m *Match) IsWalkable(x, y int) bool {
	if x < 0 || x >= boardSize || y < 0 || y >= boardSize {
		return false
	}
	if m.Walls[x][y] {
		return false
	}
	otherHero := &m.Heroes[1-m.Active]
	if otherHero.X == x && otherHero.Y == y {
		return false
	}
	return true
}

func (m *Match) MovementCost(x, y int) int {
	if x == 2 && y == 4 {
		return 2
	}
	return 1
}

func (m *Match) CanMoveActiveHero(dx, dy int) bool {
	if m.Finished {
		return false
	}
	if dx > 1 || dx < -1 {
		return false
	}
	if dy > 1 || dy < -1 {
		return false
	}
	if dx == 0 && dy == 0 {
		return false
	}
	if dx != 0 && dy != 0 {
		return false
	}

	hero := &m.Heroes[m.Active]

	destx := hero.X + dx
	desty := hero.Y + dy

	if !m.IsWalkable(destx, desty) {
		return false
	}

	if hero.MP < m.MovementCost(destx, desty) {
		return false
	}

	return true
}

func (m *Match) MoveActiveHero(dx, dy int) bool {
	if !m.CanMoveActiveHero(dx, dy) {
		return false
	}

	hero := &m.Heroes[m.Active]

	hero.X = hero.X + dx
	hero.Y = hero.Y + dy
	hero.MP -= m.MovementCost(hero.X, hero.Y)
	return true
}

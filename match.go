package main

const (
	boardSize = 10
	maxAP     = 1
	maxMP     = 3
)

type Hero struct {
	X, Y int
	HP   int
	AP   int
	MP   int
}

type Match struct {
	Heroes [2]Hero
	Active int
}

func NewMatch() *Match {
	return &Match{
		Heroes: [2]Hero{
			{X: 1, Y: 4, HP: 20, AP: maxAP, MP: maxMP},
			{X: 8, Y: 4, HP: 20, AP: maxAP, MP: maxMP},
		},
		Active: 0,
	}
}

func (m *Match) EndTurn() {
	m.Active = 1 - m.Active

	hero := &m.Heroes[m.Active]
	hero.AP = maxAP
	hero.MP = maxMP
}

func (m *Match) MoveActiveHero(dx, dy int) bool {
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

	if hero.MP < 1 {
		return false
	}

	destx := hero.X + dx
	desty := hero.Y + dy

	if destx < 0 || destx >= boardSize || desty < 0 || desty >= boardSize {
		return false
	}

	otherHero := &m.Heroes[1-m.Active]
	if destx == otherHero.X && desty == otherHero.Y {
		return false
	}

	hero.X = destx
	hero.Y = desty
	hero.MP -= 1
	return true
}

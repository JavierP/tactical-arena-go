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
	Heroes   [2]Hero
	Active   int
	Walls    [boardSize][boardSize]bool
	Finished bool
}

func NewMatch() *Match {
	localMatch := &Match{
		Heroes: [2]Hero{
			{X: 1, Y: 4, HP: 20, AP: maxAP, MP: maxMP},
			{X: 8, Y: 4, HP: 20, AP: maxAP, MP: maxMP},
		},
		Active: 0,
	}
	localMatch.Walls[4][3] = true
	localMatch.Walls[4][4] = true
	localMatch.Walls[4][5] = true
	return localMatch
}

func (m *Match) EndTurn() {
	if m.Finished {
		return
	}
	m.Active = 1 - m.Active

	hero := &m.Heroes[m.Active]
	hero.AP = maxAP
	hero.MP = maxMP
}

func (m *Match) MoveActiveHero(dx, dy int) bool {
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

	if hero.MP < 1 {
		return false
	}

	destx := hero.X + dx
	desty := hero.Y + dy

	// Checking board limits
	if destx < 0 || destx >= boardSize || desty < 0 || desty >= boardSize {
		return false
	}
	// Checking for Walls
	if m.Walls[destx][desty] {
		return false
	}
	// checking for other Player
	otherHero := &m.Heroes[1-m.Active]
	if destx == otherHero.X && desty == otherHero.Y {
		return false
	}

	hero.X = destx
	hero.Y = desty
	hero.MP -= 1
	return true
}

func (m *Match) Strike() bool {
	if m.Finished {
		return false
	}
	hero := &m.Heroes[m.Active]
	otherHero := &m.Heroes[1-m.Active]

	if hero.AP < 1 {
		return false
	}

	horizontalDistance := otherHero.X - hero.X
	verticalDistance := otherHero.Y - hero.Y

	if horizontalDistance < 0 {
		horizontalDistance = -horizontalDistance
	}
	if verticalDistance < 0 {
		verticalDistance = -verticalDistance
	}

	totalDistance := horizontalDistance + verticalDistance

	if totalDistance != 1 {
		return false
	}

	hero.AP -= 1
	otherHero.HP -= 5

	if otherHero.HP <= 0 {
		otherHero.HP = 0
		m.Finished = true
	}
	return true
}

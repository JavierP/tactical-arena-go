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

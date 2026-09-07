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

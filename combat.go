package main

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

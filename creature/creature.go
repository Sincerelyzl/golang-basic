package creature

import "fmt"

var (
	defaultPlayerHealth = 2000
	defaultPlayerLevel  = 1
)

type Creature interface {
	ShowLifePoint()
	GetLifePoint() int
	RecieveDamage(dmg *int)
}

type player struct {
	LifePoint int
	Level     int
}

type monster struct {
	LifePoint int
	Level     int
}

func (p player) ShowLifePoint() {
	fmt.Println("player LifePoint : ", p.LifePoint)
}

func (m monster) ShowLifePoint() {
	fmt.Println("monster LifePoint : ", m.LifePoint)
}

func (p player) GetLifePoint() int {
	return p.LifePoint
}

func (m monster) GetLifePoint() int {
	return m.LifePoint
}

func (p *player) RecieveDamage(dmg *int) {
	p.LifePoint -= *dmg
}

func (m *monster) RecieveDamage(dmg *int) {
	m.LifePoint -= *dmg
}

func NewPlayer(lifepoint int) player {
	player := player{
		LifePoint: defaultPlayerHealth,
		Level:     defaultPlayerLevel,
	}

	if lifepoint != 0 {
		player.LifePoint = lifepoint
	}

	return player
}

func AttackEvent(attacker Creature, reciever Creature, dmg int) {
	reciever.ShowLifePoint()
	fmt.Println("Attacker ", attacker)
	fmt.Println("Reciever ", reciever)
	fmt.Println("Reciever Is Taken Damaged For ", dmg, " damage")
	reciever.RecieveDamage(&dmg)
	reciever.ShowLifePoint()
}

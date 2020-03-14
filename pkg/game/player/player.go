package player

import (
	"github.com/Sighr/hero-realms/pkg/game/cards"
)

type Player struct {
	Hand      cards.Container
	Deck      cards.Container
	Discard   cards.Container
	Field     cards.Container
	Creatures cards.Container
	Hp        int
	TempHp    int
	TempDmg   int
	TempHeal  int
}

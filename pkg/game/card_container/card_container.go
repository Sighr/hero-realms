package card_container

import "github.com/Sighr/hero-realms/pkg/game/cards"

type CardContainer interface {
	AddCard(*cards.Card)
	RemoveCard(*cards.Card)
	GetCards() []*cards.Card // do not modify cards returned!
}

package cards

type CardContainer interface {
	AddCard(*Card)
	RemoveCard(*Card)
	GetCards() []*Card // do not modify cards returned!
}

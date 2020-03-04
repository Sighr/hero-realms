package cards

const (
	FactionGreen = "WILD"
	FactionYellow = "IMPERIAL"
	FactionBlue = "CONTRABAND"
	FactionRed = "DEMON"
)

type Card struct {
	id int
	uniqueId int
	name string
	faction string
	creature Creature
	cost int
	special Action
	synergy Action
	sacrifice Action
}
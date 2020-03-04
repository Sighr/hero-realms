package cards

type Action interface {
	Perform(interface {})
	GetAvailableTargets() string
}
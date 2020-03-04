package stack

type Stack interface {
	Shuffle()
	Push(interface{})
	Pop() (interface{}, error)
}
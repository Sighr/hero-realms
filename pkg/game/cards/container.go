package cards

import "math/rand"

type Container struct {
	arr []*Card
}

func (cs *Container) Push(card *Card) {
	cs.arr = append(cs.arr, card)
}

func (cs *Container) Pop() (*Card, error) {
	if len(cs.arr) > 0 {
		card := cs.arr[len(cs.arr)-1]
		cs.arr = cs.arr[:len(cs.arr)-1]
		return card, nil
	} else {
		return nil, &EmptyStack{}
	}
}

func (cs *Container) Shuffle() {
	for i := 1; i < len(cs.arr); i++ {
		num := rand.Intn(i)
		cs.arr[i], cs.arr[num] = cs.arr[num], cs.arr[i]
	}
}

func (cs *Container) AddCard(card *Card) {
	cs.Push(card)
}

func (cs *Container) RemoveCard(card *Card) {
	for idx, val := range cs.arr {
		if val == card {
			cs.arr = append(cs.arr[:idx], cs.arr[idx+1])
		}
	}
}

func (cs *Container) GetCards() []*Card {
	// potential insecurity
	return cs.arr
}

type EmptyStack struct{}

func (e *EmptyStack) Error() string {
	return "The stack is empty"
}

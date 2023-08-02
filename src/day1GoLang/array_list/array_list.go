package arrayList

import (
	"errors"
)

type ArrayList struct {
	length int
	list   []int
}

func (a *ArrayList) append(val int) {
	a.length++

	a.list = append(a.list, val)
}

func (a *ArrayList) get(idx int) (int, error) {
	if idx < 0 || idx >= a.length {
		return 0, errors.New("index out of bounds")
	}

	return a.list[idx], nil
}

func (a *ArrayList) prepend(val int) {
	a.length++

	if a.length == 1 {
		a.list = append(a.list, val)
		return
	}

	// Shift the array one space to the right.
	a.list = append([]int{val}, a.list...)
}

func (a *ArrayList) remove(val int) (int, error) {
	var removedIdx int

	for i := 0; i < a.length; i++ {
		if a.list[i] == val {
			removedIdx = i
			break
		}
	}

	if removedIdx > 0 || a.list[removedIdx] == val {
		return a.removeAt(removedIdx)
	}

	return val, errors.New("value not found")
}

func (a *ArrayList) removeAt(idx int) (int, error) {
	if idx < 0 || idx >= a.length {
		return 0, errors.New("index out of bounds")
	}

	a.length--
	val := a.list[idx]
	a.list = append(a.list[:idx], a.list[idx+1:]...)

	return val, nil
}

package arrayList

import (
	"testing"
)

func TestArrayList(t *testing.T) {
	list := &ArrayList{list: [1]int{} }

	list.append(5)
	list.append(7)
	list.append(9)

	getAndCheck(t, list, 2, 9)	
	removeAtAndCheck(t, list, 1, 7)
	check(t, list.length, 2)

	list.append(11)

	removeAtAndCheck(t, list, 1, 9)

	if _, err := list.remove(9); err != "index out of bounds" {
		t.Errorf("should throw error 'index out of bounds'")
	}

	removeAtAndCheck(t, list, 0, 5)
	removeAtAndCheck(t, list, 0, 11)
	check(t, list.length, 0)

	list.prepend(5)
	list.prepend(7)
	list.prepend(9)

	getAndCheck(t, list, 2, 5)
	getAndCheck(t, list, 0, 9)
	removeAndCheck(t, list, 9, 9)
	check(t, list.length, 2)
	getAndCheck(t, list, 0, 7) 
}

func check(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func getAndCheck(t testing.TB, list *ArrayList, index, want int) {
	t.Helper()

	if got, err := list.get(index); err == nil {
		check(t, got, want)
	} else {
		t.Errorf("%s", err)
	}
}

func removeAtAndCheck(t testing.TB, list *ArrayList, index, want int) {
	t.Helper()

	if got, err := list.removeAt(index); err == nil {
		check(t, got, want)
	} else {
		t.Errorf("%s", err)
	}
}

func removeAndCheck(t testing.TB, list *ArrayList, val, want int) {
	t.Helper()

	if got, err := list.remove(val); err == nil {
		check(t, got, want)
	} else {
		t.Errorf("%s", err)
	}
}


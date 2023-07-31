package arrayList

type ArrayList struct {
	capacity int
	length int
	list *[]int
}

func (a *ArrayList) append(val int) int {
	a.length++
}

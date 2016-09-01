package library

type Archive struct {
	Year     int
	Articles []Article
}

type Archives []Archive

func (a Archives) Len() int {
	return len(a)
}

func (a Archives) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Archives) Less(i, j int) bool {
	return a[i].Year > a[j].Year
}

package foo

type Foo struct {
	Bar
}

type Bar struct {
	Count int
}

func (b *Bar) Double() int {
	return b.Count * 2
}

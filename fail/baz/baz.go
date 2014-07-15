package baz

type Baz struct{
	name string
}

func (b *Baz) DoSomething() string {
	return "did something"
}

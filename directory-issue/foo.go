package foo

type Baz struct {
	name string
}

func NewBaz() *Baz {
	return &Baz{}
}

func (b *Baz) DoSomething() string {
	return "doing sth."
}

type Foo struct {}

func (f *Foo) TryMe() {
    bazz := NewBaz()
    str := bazz.DoSomething()
    fmt.Println(str)
}

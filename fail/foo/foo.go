package foo

import(
    "fmt"
    "github.com/ilovewchao/example/baz"
)

type Foo struct {}

func (f *Foo) TryMe() {
    var bazz baz.Baz
    str := bazz.DoSomething()
    fmt.Println(str)
}

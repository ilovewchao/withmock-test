package foo

import(
    "fmt"
    "github.com/ilovewchao/withmock-test/success/baz"
)

type Foo struct {}

func (f *Foo) TryMe() {
    var bazz baz.Baz
    str := bazz.DoSomething()
    fmt.Println(str)
}

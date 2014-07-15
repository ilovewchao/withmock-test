package foo

import(
	"encoding/json"
    "fmt"
    "github.com/wchao/example/baz"
)

type Foo struct {}

func (f *Foo) TryMe() {
    var bazz baz.Baz
	
	err := json.Unmarshal([]byte(`{"name": "cody"}`), &bazz)
	if err != nil {
		fmt.Println(err)
		return
	}
	
    err = bazz.DoSomething()
	if err != nil {
		fmt.Printf("Error: %s",err)
	}
    return
}

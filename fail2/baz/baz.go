package baz

import (
	"fmt"
)

type Baz struct{
	name string
}

func (b *Baz) DoSomething() error {
	if len(b.name) == 0 {
		return fmt.Errorf("has no name")
	}
	return nil
}

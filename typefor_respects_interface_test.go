package reflectnotes

import (
	"fmt"
	"reflect"
)

func ExampleTypeForRespectsInterface() {
	type I interface {
		String() string
	}

	fmt.Println(reflect.TypeFor[I]())

	// Output:
	// reflectnotes.I
}

package reflectnotes

import (
	"fmt"
	"reflect"
	"time"
)

func ExampleValueOfBypassesInterface() {
	type I interface {
		String() string
	}
	var c = time.Now()
	var i I = c

	fmt.Println("c value type:", reflect.ValueOf(c).Type())
	fmt.Println("i value type:", reflect.ValueOf(i).Type())

	// Output:
	// c value type: time.Time
	// i value type: time.Time
}

func ExampleValueOfBypassesInterface_nil() {
	type I interface {
		String() string
	}
	var i I = nil
	fmt.Println("i is valid?", reflect.ValueOf(i).IsValid())
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("IsNil() panicked")
		}
	}()
	fmt.Println("i is nil?", reflect.ValueOf(i).IsNil())

	// Output:
	// i is valid? false
	// IsNil() panicked
}
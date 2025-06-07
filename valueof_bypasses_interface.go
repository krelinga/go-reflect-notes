package reflectnotes

// [reflect.ValueOf] bypasses interfaces.
//
// When you pass a value to [reflect.ValueOf], it returns a [reflect.Value] that represents the
// concrete type of the value, not the interface type.
//
// If you pass a nil value to [reflect.ValueOf], it returns a [reflect.Value] that is not valid. This means
// that you cannot call methods on it, even methods that you might expect to work (like [reflect.Value.IsNil]).
//
// If you want to inspect interface types you should use [reflect.TypeFor] instead, see [TypeForRespectsInterface].
func ValueOfBypassesInterface() {}

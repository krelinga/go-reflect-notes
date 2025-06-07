package reflectnotes

// [reflect.ValueOf] bypasses interfaces.
//
// When you pass a value to [reflect.ValueOf], it returns a [reflect.Value] that represents the
// concrete type of the value, not the interface type.
func ValueOfBypassesInterface() {}

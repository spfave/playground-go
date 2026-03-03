package reflection

import "reflect"

// Initial implementation
// func walk(x any, fn func(input string)) {
// 	val := reflect.ValueOf(x)
// 	name := val.Field(0)
// 	fn(name.String())
// }

// ----------------------------------------------------------------------------------- //
func walk(x any, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {

	}
}

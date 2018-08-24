package display

import (
	"fmt"
	"reflect"
	"github.com/elycruz/go-programming-lang-book/chp12/format"
)

func Display (name string, x interface{}, maxDepth int) {
	fmt.Printf("Display %s (%T):\n", name, x)
	if maxDepth <= 0 {
		maxDepth = 10
	}
	display(name, 0, maxDepth, reflect.ValueOf(x))
}

func display (path string, currDepth int, maxDepth int, v reflect.Value) {
	if currDepth >= maxDepth {
		return
	}
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), currDepth + 1, maxDepth, v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fPath, currDepth + 1, maxDepth, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			path2 := fmt.Sprintf("%s[%s]", path,
				format.FormatAtom(key))
			vAtIndex := v.MapIndex(key)
			display(path2, currDepth + 1, maxDepth, vAtIndex)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), currDepth + 1, maxDepth, v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type)
			display(path + ".value", currDepth + 1, maxDepth, v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, format.FormatAtom(v))
	}
}

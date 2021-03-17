// Copyright (c) 2021, Hugues Guilleus. All rights reserved.
// Use of this source code is governed by a BBSD 3-Clause License
// that can be found in the LICENSE file.

package reflectjs

import (
	"github.com/HuguesGuilleus/go-workerglobalscope/reflectjs/uint8array"
	"reflect"
	"strconv"
	"syscall/js"
	"time"
)

// Transform a Go value into a JavaScript value.
//
// For the type: js.Value, js.Wrapper, js.Func, nil, bool, string,
// int, int8, int16, int32, uint, uint8, uint16, uint32, float32,
// float64 it's like js.ValueOf function.
//
//  | Go                     | JavaScript             |
//  | ---------------------- | ---------------------- |
//  | int64                  | BigInt                 |
//  | uint64                 | BigInt                 |
//  | []byte                 | Uint8array             |
//  | time.Time              | Date                   |
//  | map[...]bool           | Set                    |
//  | map[...]...            | Map                    |
//  | struct{...}            | Object                 |
//  | []... (slice or array) | Array                  |
//
// Other type like function or channel panics.
func ToJs(v interface{}) js.Value {
	switch v := v.(type) {
	case js.Value:
		return v
	case js.Wrapper:
		return v.JSValue()
	case js.Func, nil, bool, string,
		int, int8, int16, int32,
		uint, uint8, uint16, uint32,
		float32, float64:
		return js.ValueOf(v)
	case uint64:
		return BigInt.Invoke(strconv.FormatInt(int64(v), 10))
	case int64:
		return BigInt.Invoke(strconv.FormatInt(v, 10))
	case []byte:
		return uint8array.New(v)
	case time.Time:
		if v.IsZero() {
			return Date.New(0)
		}
		j, _ := v.MarshalText()
		return Date.New(string(j))
	default:
		return toJsReflect(reflect.ValueOf(v))
	}
}

func toJsReflect(v reflect.Value) js.Value {
	t := v.Type()
	switch k := t.Kind(); k {
	case reflect.Ptr:
		return ToJs(reflect.Indirect(v).Interface())

	case reflect.Array, reflect.Slice:
		l := v.Len()
		array := Array.New(l)
		for i := 0; i < l; i++ {
			array.SetIndex(i, ToJs(v.Index(i).Interface()))
		}
		return array

	case reflect.Map:
		if t.Elem().Kind() == reflect.Bool {
			set := Set.New()
			r := v.MapRange()
			for r.Next() {
				set.Call("add", ToJs(r.Key().Interface()))
			}
			return set
		} else {
			m := Map.New()
			r := v.MapRange()
			for r.Next() {
				m.Call("set", ToJs(r.Key().Interface()), ToJs(r.Value().Interface()))
			}
			return m
		}

	case reflect.Struct:
		obj := Object.New()
		num := v.Type().NumField()
		for i := 0; i < num; i++ {
			f := t.Field(i)
			if f.PkgPath == "" {
				obj.Set(f.Name, ToJs(v.Field(i).Interface()))
			}
		}
		return obj

	default:
		panic("Unvalid type: " + k.String())
	}
}

var (
	BigInt js.Value = js.Global().Get("BigInt")
	Date   js.Value = js.Global().Get("Date")
	Object js.Value = js.Global().Get("Object")
	Array  js.Value = js.Global().Get("Array")
	Map    js.Value = js.Global().Get("Map")
	Set    js.Value = js.Global().Get("Set")
)

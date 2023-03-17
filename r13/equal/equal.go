// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Package equal zapewnia relację głębokiej równoważności dla dowolnych wartości.
package equal

import (
	"reflect"
	"unsafe"
)

//!+
func equal(x, y reflect.Value, seen map[comparison]bool) bool {
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}
	if x.Type() != y.Type() {
		return false
	}

	// ...pominięta kontrola cykli (zostanie pokazana później)...

	//!-
	//!+cyclecheck
	// Kontrola cykli.
	if x.CanAddr() && y.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		yptr := unsafe.Pointer(y.UnsafeAddr())
		if xptr == yptr {
			return true // identyczne referencje
		}
		c := comparison{xptr, yptr, x.Type()}
		if seen[c] {
			return true // już widziane
		}
		seen[c] = true
	}
	//!-cyclecheck
	//!+
	switch x.Kind() {
	case reflect.Bool:
		return x.Bool() == y.Bool()

	case reflect.String:
		return x.String() == y.String()

	// ...dla zwięzłości pominięte przypadki liczbowe...

	//!-
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		return x.Int() == y.Int()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		return x.Uint() == y.Uint()

	case reflect.Float32, reflect.Float64:
		return x.Float() == y.Float()

	case reflect.Complex64, reflect.Complex128:
		return x.Complex() == y.Complex()
	//!+
	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return x.Pointer() == y.Pointer()

	case reflect.Ptr, reflect.Interface:
		return equal(x.Elem(), y.Elem(), seen)

	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		for i := 0; i < x.Len(); i++ {
			if !equal(x.Index(i), y.Index(i), seen) {
				return false
			}
		}
		return true

	// ...dla zwięzłości pominięte przypadki struktury i mapy...
	//!-
	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if !equal(x.Field(i), y.Field(i), seen) {
				return false
			}
		}
		return true

	case reflect.Map:
		if x.Len() != y.Len() {
			return false
		}
		for _, k := range x.MapKeys() {
			if !equal(x.MapIndex(k), y.MapIndex(k), seen) {
				return false
			}
		}
		return true
		//!+
	}
	panic("nieosiagalne")
}

//!-

//!+comparison
// Equal raportuje, czy x i y są głęboko równoważne.
//!-comparison
//
// Klucze mapy są zawsze porównywane za pomocą operatora ==, a nie głeboko.
// (Ma to znaczenie dla kluczy zawierających wskaźniki lub interfejsy.)
//!+comparison
func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}

type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}

//!-comparison

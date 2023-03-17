// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Package params zapewnia oparty na refleksji parser dla parametrów URL.
package params

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

//!+Unpack

// Unpack zapełnia pola struktury wskazane przez ptr 
// z parametrów żądania HTTP w req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Tworzy mapę pól z kluczami w postaci efektywnych nazw.
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem() // zmienna struktury
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // reflect.StructField
		tag := fieldInfo.Tag           // reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	// Aktualizuje pole struktury dla każdego parametru w żądaniu.
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // ignorowanie nierozpoznanych parametrów HTTP
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

//!-Unpack

//!+populate
func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("nieobsługiwany rodzaj %s", v.Type())
	}
	return nil
}

//!-populate

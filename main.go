package bitfield

import (
	"fmt"
	"reflect"
	"strconv"
)

func Unpack(source uint64, result any) error {
	theType := reflect.TypeOf(result).Elem()
	theValue := reflect.ValueOf(result).Elem()
	for i, n := 0, theType.NumField(); i < n; i++ {
		field := theType.Field(i)
		if bitTag, ok := field.Tag.Lookup("bit"); ok {
			bit, err := strconv.Atoi(bitTag)
			if err != nil {
				return fmt.Errorf("%s: expect number in `bit` tag", field.Name)
			}
			bitValue := source & ((1 << bit) - 1)

			theField := theValue.Field(i)
			if theField.CanUint() {
				// println(field.Name, "set", bitValue, "as uint")
				theField.SetUint(bitValue)
			} else if theField.CanInt() {
				// println(field.Name, "set", bitValue, "as int")
				theField.SetInt(int64(bitValue))
			} else {
				return fmt.Errorf("%s: expected the field int or uint", field.Name)
			}
			source >>= bit
		}
	}
	return nil
}

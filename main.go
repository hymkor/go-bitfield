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

func Pack(source any) (uint64, error) {
	theType := reflect.TypeOf(source).Elem()
	theValue := reflect.ValueOf(source).Elem()

	var result uint64
	for i := theType.NumField() - 1; i >= 0; i-- {
		field := theType.Field(i)
		if bitTag, ok := field.Tag.Lookup("bit"); ok {
			bit, err := strconv.Atoi(bitTag)
			if err != nil {
				return 0, fmt.Errorf("%s: expect number in `bit` tag", field.Name)
			}
			result <<= bit
			mask := ((1 << bit) - 1)

			theField := theValue.Field(i)
			if theField.CanUint() {
				// println(field.Name, "set", bitValue, "as uint")
				result |= theField.Uint() & uint64(mask)
			} else if theField.CanInt() {
				// println(field.Name, "set", bitValue, "as int")
				result |= uint64(theField.Int()) & uint64(mask)
			} else {
				return 0, fmt.Errorf("%s: expected the field int or uint", field.Name)
			}
		}
	}
	return result, nil
}

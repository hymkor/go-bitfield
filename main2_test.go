package bitfield_test

import (
	"testing"

	"github.com/hymkor/go-bitfield"
)

type DosDateNG1 struct {
	Second int  `bit:"5"`
	Min    uint `bit:"X"`
	Hour   int  `bit:"5"`
}

func TestUnpackNg1(t *testing.T) {
	var dt DosDateNG1
	if err := bitfield.Unpack(0x7423, &dt); err == nil {
		t.Fatal("Any error is not found in invalid structure")
	} else {
		println("OK: expect error:", err.Error())
	}
}

type DosDateNG2 struct {
	Second int  `bit:"5"`
	Min    uint `bit:"6"`
	Hour   *int `bit:"5"`
}

func TestUnpackNg2(t *testing.T) {
	var dt DosDateNG2
	if err := bitfield.Unpack(0x7423, &dt); err == nil {
		t.Fatal("Any error is not found in invalid structure")
	} else {
		println("OK: expect error:", err.Error())
	}
}

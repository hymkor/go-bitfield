package bitfield_test

import (
	"testing"

	"github.com/hymkor/go-bitfield"
)

type DosDate struct {
	Second int  `bit:"5"`
	Min    uint `bit:"6"`
	Hour   int  `bit:"5"`
}

func TestRead(t *testing.T) {
	var dt DosDate
	if err := bitfield.Read(0x7423, &dt); err != nil {
		t.Fatal(err.Error())
		return
	}
	if dt.Second != 3 {
		t.Fatalf("Second: expect 3, but %d", dt.Second)
	}
	if dt.Min != 33 {
		t.Fatalf("Min: expect 33, but %d", dt.Min)
	}
	if dt.Hour != 14 {
		t.Fatalf("Hour: expect 14,but %d", dt.Hour)
	}
}

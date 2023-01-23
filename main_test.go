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

func TestUnpack(t *testing.T) {
	const packedValue = 0x7423

	var dt DosDate
	if err := bitfield.Unpack(uint64(packedValue), &dt); err != nil {
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

func TestPack(t *testing.T) {
	const packedValue = 0x7423

	var dt = &DosDate{
		Second: 3,
		Min:    33,
		Hour:   14,
	}

	value, err := bitfield.Pack(dt)
	if err != nil {
		t.Fatalf("pack: %s", err.Error())
	}
	if value != packedValue {
		t.Fatalf("pack: expect %d, but %d", packedValue, value)
	}
}

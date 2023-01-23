package bitfield_test

import (
	"testing"

	"github.com/hymkor/go-bitfield"
)

const packedValue = 0x7423

type DosDate struct {
	Second int  `bit:"5"`
	Min    uint `bit:"6"`
	Hour   int  `bit:"5"`
}

func TestUnpack(t *testing.T) {
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

const (
	secondBit = 5
	minBit    = 6
	hourBit   = 5
)

func TestUnpackInline(t *testing.T) {
	dt := bitfield.UnpackInline(packedValue, secondBit, minBit, hourBit)
	if dt[0] != 3 {
		t.Fatalf("Second: expect 3, but %d", dt[0])
	}
	if dt[1] != 33 {
		t.Fatalf("Min: expect 33, but %d", dt[1])
	}
	if dt[2] != 14 {
		t.Fatalf("Hour: expect 14,but %d", dt[2])
	}
}

func TestPackInline(t *testing.T) {
	dt := bitfield.PackInline(secondBit, 3, minBit, 33, hourBit, 14)
	if dt != packedValue {
		t.Fatalf("PackInline expect %d,but %d", packedValue, dt)
	}
}

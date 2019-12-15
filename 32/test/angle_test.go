package geomTest

import (
	. "github.com/tadeuszjt/geom/32"
	"math"
	"testing"
)

func TestMakeAngle(t *testing.T) {
	cases := []struct {
		rad, result float32
	}{
		{0, 0},
		{2 * math.Pi, 0},
		{math.Pi, math.Pi},
		{3 * math.Pi, math.Pi},
		{-3 * math.Pi, math.Pi},
		{-3 * math.Pi / 2, math.Pi / 2},
		{nan, nan},
		{pInf, nan},
		{nInf, nan},
	}

	for _, c := range cases {
		expected := c.result
		actual := float32(MakeAngle(c.rad))
		if !floatIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestAnglePlus(t *testing.T) {
	cases := []struct {
		a, b, result float32
	}{
		{0, 0, 0},
		{1, 2, 3},
		{3 * math.Pi, math.Pi, 0},
		{-3 * math.Pi, -3 * math.Pi / 2, 3 * math.Pi / 2},
		{nan, -3 * math.Pi / 2, nan},
		{pInf, -3 * math.Pi / 2, nan},
		{math.Pi, nInf, nan},
	}

	for _, c := range cases {
		expected := c.result
		a := MakeAngle(c.a)
		b := MakeAngle(c.b)
		actual := float32(a.Plus(b))
		if !floatIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

var a Angle

func BenchmarkMakeAngle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a = MakeAngle(12)
	}
}

package geomTest

import (
	. "github.com/tadeuszjt/geom/generic"
	"testing"
)

func cuboidIdentical(a, b Cuboid[float64]) bool {
	return vec3Identical(a.Min, b.Min) && vec3Identical(a.Max, b.Max)
}

func TestCuboidCentred(t *testing.T) {
	cases := []struct {
		width, height, depth float64
		result               Cuboid[float64]
	}{
		{0, 0, 0, Cuboid[float64]{}},
		{10, 20, 30, Cuboid[float64]{Min: Vec3[float64]{-5, -10, -15}, Max: Vec3[float64]{5, 10, 15}}},
	}

	for _, c := range cases {
		expected := c.result
		actual := CuboidCentred(c.width, c.height, c.depth)
		if !cuboidIdentical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}

func TestCuboidOrigin(t *testing.T) {
	cases := []struct {
		width, height, depth float64
		result               Cuboid[float64]
	}{
		{0, 0, 0, Cuboid[float64]{}},
		{10, 20, 30, Cuboid[float64]{Min: Vec3[float64]{}, Max: Vec3[float64]{10, 20, 30}}},
	}

	for _, c := range cases {
		expected := c.result
		actual := CuboidOrigin(c.width, c.height, c.depth)
		if !cuboidIdentical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}

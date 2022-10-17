package geomTest

import (
	"github.com/tadeuszjt/geom/generic"
	//"math"
	"testing"
)

func polyIdentical(a, b geom.Poly[float64]) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !vec2Identical(a[i], b[i]) {
			return false
		}
	}
	return true
}

func TestPolyConvert(t *testing.T) {
	a := geom.Poly[float32]{
		{1, 2},
		{0.001, 0},
		{3, -5.7},
	}

	expected := geom.Poly[float64]{
		{1, 2},
		{0.001, 0},
		{3, -5.7},
	}

	actual := geom.PolyConvert[float32, float64](a)

	if !polyIdentical(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestPolyCopy(t *testing.T) {
	cases := []struct {
		poly   geom.Poly[float64]
		result geom.Poly[float64]
	}{
		{
			[]geom.Vec2[float64]{},
			[]geom.Vec2[float64]{},
		},
		{
			[]geom.Vec2[float64]{{1, 2}},
			[]geom.Vec2[float64]{{1, 2}},
		},
		{
			[]geom.Vec2[float64]{{1, 2}, {3.4, 5.6}},
			[]geom.Vec2[float64]{{1, 2}, {3.4, 5.6}},
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := geom.PolyCopy(c.poly)

		if !polyIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
		if len(expected) > 0 && &expected[0] == &actual[0] {
			t.Errorf("didn't copy, same pointer.")
		}
	}
}

func TestPolyArea(t *testing.T) {
	cases := []struct {
		poly geom.Poly[float64]
		area float64
	}{
		{[]geom.Vec2[float64]{{0, 0}, {1, 0}, {1, 1}, {0, 1}}, 1.0},
		{[]geom.Vec2[float64]{{0, 0}, {1, 0}, {1, 1}}, 0.5},
		{[]geom.Vec2[float64]{{0, 0}, {2, 0}, {2, 2}}, 2.0},
		{[]geom.Vec2[float64]{
			{0, 0}, {2, 0}, {2, 2}, {1, 3}, {0, 2}},
			5.0,
		},
	}

	for _, c := range cases {
		expected := c.area
		actual := c.poly.Area()

		if !floatIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestPolyCentroid(t *testing.T) {
	cases := []struct {
		poly     geom.Poly[float64]
		centroid geom.Vec2[float64]
	}{
		{
			[]geom.Vec2[float64]{{0, 0}, {2, 0}, {2, 2}, {0, 2}},
			geom.Vec2[float64]{1, 1},
		},
		{
			[]geom.Vec2[float64]{{2, 0}, {4, 2}, {2, 4}, {0, 2}},
			geom.Vec2[float64]{2, 2},
		},
	}

	for _, c := range cases {
		expected := c.centroid
		actual := c.poly.Centroid()

		if !vec2Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestPolyMomentOfInertia(t *testing.T) {
	cases := []struct {
		poly geom.Poly[float64]
		moi  float64
	}{
		{
			[]geom.Vec2[float64]{{2, -1.5}, {2, 1.5}, {-2, 1.5}, {-2, -1.5}},
			(1. / 12.) * (4 * 3) * (4*4 + 3*3),
		},
	}

	for _, c := range cases {
		expected := c.moi
		actual := c.poly.MomentOfInertia()

		if !floatIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestPolyContains(t *testing.T) {
	cases := []struct {
		poly   geom.Poly[float64]
		point  geom.Vec2[float64]
		result bool
	}{
		{
			geom.Poly[float64]{{0, 0}, {1, 0}, {1, 1}, {0, 1}},
			geom.Vec2[float64]{0.5, 0.5},
			true,
		},
		{
			geom.Poly[float64]{{0, 0}, {1, 0}, {1, 1}, {0, 1}},
			geom.Vec2[float64]{1.0001, 0.5},
			false,
		},
		{
			geom.Poly[float64]{{0, 0}, {1, 0}, {0, 1}},
			geom.Vec2[float64]{0.5, 0.5},
			true,
		},
		{
			geom.Poly[float64]{{0, 0}, {1, 0}, {1, -1}, {2, 2}},
			geom.Vec2[float64]{0.5, 0},
			true,
		},
		{
			geom.Poly[float64]{{0, 0}, {1, 0}, {1, -1}, {2, 2}},
			geom.Vec2[float64]{0.5, -0.00001},
			false,
		},
		{
			geom.Poly[float64]{{0, 0}, {1, 0}, {1, -1}, {2, 2}},
			geom.Vec2[float64]{1, 0},
			true,
		},
		{
			geom.Poly[float64]{{0, 0}, {1, 0}, {1, -1}, {2, 2}},
			geom.Vec2[float64]{0.9999, -0.5},
			false,
		},
		{
			geom.Poly[float64]{{0, 0}, {1, 0}, {1, -1}, {2, 2}},
			geom.Vec2[float64]{1, -0.5},
			true,
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.poly.Contains(c.point)

		if expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

package geomTest

import (
	"github.com/tadeuszjt/geom/generic"
	//"math"
	"testing"
)

func TestPolyArea(t *testing.T) {
	cases := []struct {
		verts []geom.Vec2[float64]
		area  float64
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
		actual := geom.PolyArea(c.verts)

		if !floatIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestPolyCentroid(t *testing.T) {
	cases := []struct {
		verts    []geom.Vec2[float64]
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
		actual := geom.PolyCentroid(c.verts)

		if !vec2Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestPolyMomentOfInertia(t *testing.T) {
	cases := []struct {
		verts    []geom.Vec2[float64]
		moi float64
	}{
		{
			[]geom.Vec2[float64]{{2, -1.5}, {2, 1.5}, {-2, 1.5}, {-2, -1.5}},
			(1. / 12.) * (4*3) * (4*4+3*3),
		},
	}

	for _, c := range cases {
		expected := c.moi
		actual := geom.PolyMomentOfInertia(c.verts)

		if !floatIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

package geomTest

import (
	. "github.com/tadeuszjt/geom/generic"
	"testing"
)

func rectIdentical(a, b Rect[float64]) bool {
	return vec2Identical(a.Min, b.Min) && vec2Identical(a.Max, b.Max)
}

func TestRectZero(t *testing.T) {
	expected := Rect[float64]{Vec2[float64]{0, 0}, Vec2[float64]{0, 0}}
	actual := Rect[float64]{}
	if !rectIdentical(expected, actual) {
		t.Errorf("expected: %v, got: %v", expected, actual)
	}
}

func TestRectOrigin(t *testing.T) {
	expected := Rect[float64]{Vec2[float64]{0, 0}, Vec2[float64]{123, .456}}
	actual := RectOrigin[float64](123, .456)
	if !rectIdentical(expected, actual) {
		t.Errorf("expected: %v, got: %v", expected, actual)
	}
}

func TestRectCentred(t *testing.T) {
	cases := []struct {
		width, height float64
		result        Rect[float64]
	}{
		{0, 0, Rect[float64]{}},
		{10, 20, Rect[float64]{Min: Vec2[float64]{-5, -10}, Max: Vec2[float64]{5, 10}}},
		{12, 22, Rect[float64]{Min: Vec2[float64]{-6, -11}, Max: Vec2[float64]{6, 11}}},
		{nan, pInf, Rect[float64]{Min: Vec2[float64]{nan, nInf}, Max: Vec2[float64]{nan, pInf}}},
		{nInf, nan, Rect[float64]{Min: Vec2[float64]{pInf, nan}, Max: Vec2[float64]{nInf, nan}}},
	}

	for _, c := range cases {
		expected := c.result
		actual := RectCentred(c.width, c.height)
		if !rectIdentical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}

func TestRectCentredAt(t *testing.T) {
	cases := []struct {
		width, height float64
		position      Vec2[float64]
		result        Rect[float64]
	}{
		{0, 0, Vec2[float64]{0, 0}, Rect[float64]{}},
		{10, 20, Vec2[float64]{0, 0}, Rect[float64]{Min: Vec2[float64]{-5, -10}, Max: Vec2[float64]{5, 10}}},
		{10, 20, Vec2[float64]{3, 4}, Rect[float64]{Min: Vec2[float64]{-2, -6}, Max: Vec2[float64]{8, 14}}},
		{0, 0, Vec2[float64]{3, 4}, Rect[float64]{Min: Vec2[float64]{3, 4}, Max: Vec2[float64]{3, 4}}},
		{0.3, 0.8, Vec2[float64]{-2.3, 4}, Rect[float64]{Min: Vec2[float64]{-2.45, 3.6}, Max: Vec2[float64]{-2.15, 4.4}}},
		{-3, 0, Vec2[float64]{1, 2}, Rect[float64]{Min: Vec2[float64]{2.5, 2}, Max: Vec2[float64]{-0.5, 2}}},
	}

	for _, c := range cases {
		expected := c.result
		actual := RectCentredAt(c.width, c.height, c.position)
		if !rectIdentical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}

func TestMakeRect(t *testing.T) {
	cases := []struct {
		x, y, w, h float64
		result     Rect[float64]
	}{
		{0, 0, 0, 0, Rect[float64]{}},
		{
			3, 4, 10, 20,
			Rect[float64]{Min: Vec2[float64]{3, 4}, Max: Vec2[float64]{13, 24}},
		},
		{
			-2.3, 4, 0.3, 0.8,
			Rect[float64]{Min: Vec2[float64]{-2.3, 4}, Max: Vec2[float64]{-2, 4.8}},
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := MakeRect(c.x, c.y, c.w, c.h)
		if !rectIdentical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}

func TestRectWidth(t *testing.T) {
	cases := []struct {
		rect  Rect[float64]
		width float64
	}{
		{Rect[float64]{Min: Vec2[float64]{0, 0}, Max: Vec2[float64]{0, 0}}, 0},
		{Rect[float64]{Min: Vec2[float64]{0, 0}, Max: Vec2[float64]{10, 20}}, 10},
		{Rect[float64]{Min: Vec2[float64]{1.4, 3.2}, Max: Vec2[float64]{2.3, 4.5}}, 0.9},
		{Rect[float64]{Min: Vec2[float64]{-8.2, 1.2}, Max: Vec2[float64]{11.3, 4.5}}, 19.5},
	}

	for _, c := range cases {
		expected := c.width
		actual := c.rect.Width()
		if !floatIdentical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}

func TestRectHeight(t *testing.T) {
	cases := []struct {
		rect   Rect[float64]
		height float64
	}{
		{Rect[float64]{Min: Vec2[float64]{0, 0}, Max: Vec2[float64]{0, 0}}, 0},
		{Rect[float64]{Min: Vec2[float64]{0, 0}, Max: Vec2[float64]{10, 20}}, 20},
		{Rect[float64]{Min: Vec2[float64]{1.4, 3.2}, Max: Vec2[float64]{2.3, 4.5}}, 1.3},
		{Rect[float64]{Min: Vec2[float64]{8.2, -1.2}, Max: Vec2[float64]{11.3, 4.5}}, 5.7},
	}

	for _, c := range cases {
		expected := c.height
		actual := c.rect.Height()
		if !floatIdentical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}

func TestRectContains(t *testing.T) {
	cases := []struct {
		rect    Rect[float64]
		points  []Vec2[float64]
		results []bool
	}{
		{
			Rect[float64]{},
			[]Vec2[float64]{{0, 0}, {0.00001, 0}, {0, 0.0001}, {-0.0001, 0}, {0, -0.00001}},
			[]bool{true, false, false, false, false},
		},
		{
			Rect[float64]{Min: Vec2[float64]{0, 0}, Max: Vec2[float64]{10, 10}},
			[]Vec2[float64]{{0, 0}, {-1, -1}, {1, -1}, {10, 10}, {10, 10.00001}},
			[]bool{true, false, false, true, false},
		},
		{
			Rect[float64]{Min: Vec2[float64]{-0.1, -0.2}, Max: Vec2[float64]{0.1, 0.2}},
			[]Vec2[float64]{{-0.1, -0.2}, {0.1, -0.21}, {0.1, 0.2}, {0, 0}, {nan, 0}},
			[]bool{true, false, true, true, false},
		},
		{
			Rect[float64]{Min: Vec2[float64]{1, 2}, Max: Vec2[float64]{3, 4}},
			[]Vec2[float64]{{1, 2}, {0.8, 1.9}, {2.8, 2.2}, {3.1, 0.9}, {1.1, 4}, {0.9, 4}},
			[]bool{true, false, true, false, true, false},
		},
		{
			Rect[float64]{Min: Vec2[float64]{100, 1.3}, Max: Vec2[float64]{120, 1.8}},
			[]Vec2[float64]{{110, 1.2}, {110, 1.3}, {110, 1.7}, {110, 1.9}},
			[]bool{false, true, true, false},
		},
	}

	for _, c := range cases {
		for i := range c.points {
			expected := c.results[i]
			actual := c.rect.Contains(c.points[i])
			if actual != expected {
				t.Errorf(
					"rect: %v, point: %v, expected: %v, got: %v",
					c.rect,
					c.points[i],
					expected,
					actual,
				)
			}
		}
	}
}

func TestRectVerts(t *testing.T) {
	cases := []struct {
		rect   Rect[float64]
		result [4]Vec2[float64]
	}{
		{
			Rect[float64]{},
			[4]Vec2[float64]{{0, 0}, {0, 0}, {0, 0}, {0, 0}},
		},
		{
			RectCentred[float64](2, 2),
			[4]Vec2[float64]{{-1, -1}, {1, -1}, {1, 1}, {-1, 1}},
		},
		{
			RectCentredAt(2, 2, Vec2[float64]{-5, 5}),
			[4]Vec2[float64]{{-6, 4}, {-4, 4}, {-4, 6}, {-6, 6}},
		},
		{
			RectCentredAt(2, pInf, Vec2[float64]{nInf, 5}),
			[4]Vec2[float64]{{nInf, nInf}, {nInf, nInf}, {nInf, pInf}, {nInf, pInf}},
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.rect.Verts()

		for i := range expected {
			if !vec2Identical(expected[i], actual[i]) {
				t.Errorf("expected: %v, got: %v", expected, actual)
				break
			}
		}
	}
}

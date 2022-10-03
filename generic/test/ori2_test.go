package geomTest

import (
	. "github.com/tadeuszjt/geom/generic"
	"math"
	"testing"
)

func ori2Identical(a, b Ori2[float64]) bool {
	return floatIdentical(a.X, b.X) &&
		floatIdentical(a.Y, b.Y) &&
		floatIdentical(a.Theta, b.Theta)
}

func TestOri2(t *testing.T) {
	expected := Ori2[float64]{0, 0, 0}
	actual := Ori2[float64]{}
	if !ori2Identical(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestVec2(t *testing.T) {
	cases := []struct {
		o Ori2[float64]
		v Vec2[float64]
	}{
		{Ori2[float64]{}, Vec2[float64]{}},
		{Ori2[float64]{1, 2, 3}, Vec2[float64]{1, 2}},
		{Ori2[float64]{.1, .2, .3}, Vec2[float64]{.1, .2}},
		{Ori2[float64]{-1, -2, 3}, Vec2[float64]{-1, -2}},
		{Ori2[float64]{nan, pInf, nInf}, Vec2[float64]{nan, pInf}},
	}

	for _, c := range cases {
		expected := c.v
		actual := c.o.Vec2()
		if !vec2Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestOri2PlusEquals(t *testing.T) {
	cases := []struct {
		v, b, result Ori2[float64]
	}{
		{Ori2[float64]{}, Ori2[float64]{}, Ori2[float64]{}},
		{Ori2[float64]{0, 0, 0}, Ori2[float64]{1, 2, 3}, Ori2[float64]{1, 2, 3}},
		{Ori2[float64]{0, 0, 0}, Ori2[float64]{-1, -2, -3}, Ori2[float64]{-1, -2, -3}},
		{Ori2[float64]{1, 2, math.Pi}, Ori2[float64]{4, 5, math.Pi}, Ori2[float64]{5, 7, 2 * math.Pi}},
		{Ori2[float64]{1, 2, 3}, Ori2[float64]{-4, -5, -6}, Ori2[float64]{-3, -3, -3}},
		{Ori2[float64]{nan, 2, 3}, Ori2[float64]{4, 5, 0}, Ori2[float64]{nan, 7, 3}},
	}

	for _, c := range cases {
		expected := c.result
		c.v.PlusEquals(c.b)
		actual := c.v
		if !ori2Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestOri2Mat3Transform(t *testing.T) {
	cases := []struct {
		o      Ori2[float64]
		v      Vec2[float64]
		result Vec3[float64]
	}{
		{Ori2[float64]{}, Vec2[float64]{}, Vec3[float64]{0, 0, 1}},
		{Ori2[float64]{1, 2, 0}, Vec2[float64]{0, 0}, Vec3[float64]{1, 2, 1}},
		{Ori2[float64]{1, 2, 0}, Vec2[float64]{3, 4}, Vec3[float64]{4, 6, 1}},
		{Ori2[float64]{3, 4, math.Pi / 2}, Vec2[float64]{1, 2}, Vec3[float64]{1, 5, 1}},
		{Ori2[float64]{3, 4, -math.Pi / 2}, Vec2[float64]{1, 2}, Vec3[float64]{5, 3, 1}},
		{Ori2[float64]{-2, 8, math.Pi}, Vec2[float64]{3, -2}, Vec3[float64]{-5, 10, 1}},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.o.Mat3Transform().TimesVec2(c.v, 1)
		if !vec3Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestOri2Vec3(t *testing.T) {
	cases := []struct {
		o      Ori2[float64]
		result Vec3[float64]
	}{
		{Ori2[float64]{}, Vec3[float64]{}},
		{Ori2[float64]{1, 2, 3}, Vec3[float64]{1, 2, 3}},
		{Ori2[float64]{-1, -2, -3}, Vec3[float64]{-1, -2, -3}},
		{Ori2[float64]{0.001, 0.002, 0.003}, Vec3[float64]{0.001, 0.002, 0.003}},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.o.Vec3()
		if !vec3Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestOri2Convert(t *testing.T) {
	var a Ori2[float32]
	var b Ori2[float64]

	a = Ori2[float32]{1.2, 2.3, 3.4}
	b = Ori2Convert[float32, float64](a)

	expected := Ori2[float64]{1.2, 2.3, 3.4}
	actual := b

	if !ori2Identical(expected, actual) {
		t.Errorf("expected: %v, got: %v", expected, actual)
	}
}

func TestOri2Times(t *testing.T) {
	cases := []struct {
		a, b, result Ori2[float64]
	}{
		{Ori2[float64]{}, Ori2[float64]{}, Ori2[float64]{}},
		{Ori2[float64]{0, 0, 0}, Ori2[float64]{1, 2, 3}, Ori2[float64]{0, 0, 0}},
		{Ori2[float64]{1, 0.2, 3}, Ori2[float64]{0.4, 5, 0.6}, Ori2[float64]{0.4, 1, 1.8}},
		{Ori2[float64]{-1, -2, -3}, Ori2[float64]{4, 5, 6}, Ori2[float64]{-4, -10, -18}},
		{Ori2[float64]{nan, pInf, nInf}, Ori2[float64]{-4, -5, -6}, Ori2[float64]{nan, nInf, pInf}},
		{Ori2[float64]{nan, pInf, nInf}, Ori2[float64]{0, 0, 0}, Ori2[float64]{nan, nan, nan}},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.a.Times(c.b)
		if !ori2Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestOri2Dot(t *testing.T) {
	cases := []struct {
		a, b   Ori2[float64]
		result float64
	}{
		{Ori2[float64]{}, Ori2[float64]{}, 0},
		{Ori2[float64]{1, 2, 3}, Ori2[float64]{4, 3, -2}, 4},
		{Ori2[float64]{1, 2, 3}, Ori2[float64]{-1, -2, -3}, -14},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.a.Dot(c.b)
		if !floatIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

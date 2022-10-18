package geomTest

import (
	. "github.com/tadeuszjt/geom/generic"
	"math"
	"testing"
)

func mat3Identical(a, b Mat3[float64]) bool {
	for i := range a {
		if !floatIdentical(a[i], b[i]) {
			return false
		}
	}
	return true
}

func TestMat3Identity(t *testing.T) {
	expected := Mat3[float64]{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
	actual := Mat3Identity[float64]()
	if !mat3Identical(expected, actual) {
		t.Errorf("expected: %v, got: %v", expected, actual)
	}
}

func TestMat3Translation(t *testing.T) {
	expected := Mat3[float64]{
		1, 0, 3.4,
		0, 1, 5.6,
		0, 0, 1,
	}
	actual := Mat3Translation(Vec2[float64]{3.4, 5.6})
	if !mat3Identical(expected, actual) {
		t.Errorf("expected: %v, got: %v", expected, actual)
	}
}

func TestMat3Rotation(t *testing.T) {
	f707 := math.Sqrt(0.5)

	expected := Mat3[float64]{
		f707, -f707, 0,
		f707, f707, 0,
		0, 0, 1,
	}
	actual := Mat3Rotation[float64](math.Pi / 4)
	if !mat3Identical(expected, actual) {
		t.Errorf("expected: %v, got: %v", expected, actual)
	}
}

func TestMat3TimesVec2(t *testing.T) {
	for _, c := range []struct {
		mat    Mat3[float64]
		vec    Vec2[float64]
		bias   float64
		result Vec2[float64]
	}{
		{Mat3Identity[float64](), Vec2[float64]{1, 1}, 1, Vec2[float64]{1, 1}},
		{
			Mat3[float64]{
				1, 2, 3,
				4, 5, 6,
				7, 8, 9,
			},
			Vec2[float64]{10, 11},
			1,
			Vec2[float64]{35, 101},
		},
		{
			Mat3[float64]{
				-3, pInf, 2.2,
				0, -38, 7,
				nan, 8, -0.1,
			},
			Vec2[float64]{-1, -2},
			-3,
			Vec2[float64]{nInf, 55},
		},
		{
			Mat3[float64]{
				pInf, 0, 0,
				nInf, 0, 0,
				0.001, -0.002, 0.003,
			},
			Vec2[float64]{0, 1},
			2,
			Vec2[float64]{nan, nan},
		},
		{Mat3Translation(Vec2[float64]{1, 2}), Vec2[float64]{3, 4}, 1, Vec2[float64]{4, 6}},
		{Mat3Rotation[float64](math.Pi / 2), Vec2[float64]{2, 1}, 1, Vec2[float64]{-1, 2}},
	} {
		expected := c.result
		actual := c.mat.TimesVec2(c.vec, c.bias)
		if !vec2Identical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}

func TestMat3Camera2D(t *testing.T) {
	camera := Rect[float64]{
		Min: Vec2[float64]{10, 16},
		Max: Vec2[float64]{50, 32},
	}
	display := Rect[float64]{
		Min: Vec2[float64]{-1, -2},
		Max: Vec2[float64]{3, 4},
	}
	mat := Mat3Camera2D(camera, display)

	cases := []struct {
		point, result Vec2[float64]
	}{
		{Vec2[float64]{30, 24}, Vec2[float64]{1, 1}},
		{Vec2[float64]{10, 16}, Vec2[float64]{-1, -2}},
		{Vec2[float64]{50, 16}, Vec2[float64]{3, -2}},
		{Vec2[float64]{50, 32}, Vec2[float64]{3, 4}},
		{Vec2[float64]{10, 32}, Vec2[float64]{-1, 4}},
	}

	for _, c := range cases {
		actual := mat.TimesVec2(c.point, 1)
		expected := c.result
		if !vec2Identical(expected, actual) {
			t.Errorf("point: %v: expected: %v, got: %v", c.point, expected, actual)
		}
	}
}

func TestMat3Times(t *testing.T) {
	cases := []struct {
		a, b, result Mat3[float64]
	}{
		{
			Mat3Identity[float64](),
			Mat3Identity[float64](),
			Mat3Identity[float64](),
		},
		{
			Mat3[float64]{
				1, 2, 3,
				4, 5, 6,
				7, 8, 9,
			},
			Mat3[float64]{
				1, 2, 3,
				4, 5, 6,
				7, 8, 9,
			},
			Mat3[float64]{
				30, 36, 42,
				66, 81, 96,
				102, 126, 150,
			},
		},
		{
			Mat3[float64]{
				1, 2, 3,
				4, 5, 6,
				7, 8, 9,
			},
			Mat3[float64]{
				.1, .2, .3,
				.4, .5, .6,
				.7, .8, .9,
			},
			Mat3[float64]{
				3.0, 3.6, 4.2,
				6.6, 8.1, 9.6,
				10.2, 12.6, 15.0,
			},
		},
	}

	for _, c := range cases {
		actual := c.a.Product(c.b)
		expected := c.result
		if !mat3Identical(expected, actual) {
			t.Errorf("a: %v Times b: %v, expected: %v, got: %v",
				c.a, c.b, expected, actual)
		}
	}
}

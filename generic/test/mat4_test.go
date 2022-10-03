package geomTest

import (
	. "github.com/tadeuszjt/geom/generic"
	"math"
	"testing"
)

func mat4Identical(a, b Mat4[float64]) bool {
	for i := range a {
		if !floatIdentical(a[i], b[i]) {
			return false
		}
	}
	return true
}

func TestMat4Identical(t *testing.T) {
	cases := []struct {
		a, b   Mat4[float64]
		result bool
	}{
		{Mat4Identity[float64](), Mat4Identity[float64](), true},
		{
			Mat4[float64]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			Mat4[float64]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			true,
		},
		{
			Mat4[float64]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15.0001},
			Mat4[float64]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			false,
		},
		{
			Mat4[float64]{0, 1, 2, 3, 4, 5, 6, nan, 8, 9, 10, 11, 12, 13, 14, 15},
			Mat4[float64]{0, 1, 2, 3, 4, 5, 6, -nan, 8, 9, 10, 11, 12, 13, 14, 15},
			true,
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := mat4Identical(c.a, c.b)

		if expected != actual {
			t.Errorf("a: %v, b: %v, expected: %v, got: %v",
				c.a, c.b, expected, actual)
		}
	}
}

func TestMat4TransformVec3(t *testing.T) {
	cases := []struct {
		m      Mat4[float64]
		v      Vec3[float64]
		w      float64
		result Vec3[float64]
	}{
		{Mat4Identity[float64](), Vec3[float64]{}, 1, Vec3[float64]{}},
		{
			Mat4[float64]{
				0, 1, 2, 3,
				4, 5, 6, 7,
				8, 9, 10, 11,
				12, 13, 14, 15,
			},
			Vec3[float64]{1, 2, 3},
			4,
			Vec3[float64]{20. / 140., 60. / 140., 100. / 140.},
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.m.TransformVec3(c.v, c.w)

		if !vec3Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestMat4Perspective(t *testing.T) {
	cases := []struct {
		r, l, t, b, n, f float64
		world, screen    []Vec3[float64]
	}{
		{
			1, -1, 1, -1, 1, 4,
			[]Vec3[float64]{
				{1, 1, -1},
				{-1, -1, -1},
				{1, 1, -4},
				{-1, -1, -4},
			},
			[]Vec3[float64]{
				{1, 1, -1},
				{-1, -1, -1},
				{1. / 4., 1. / 4., 1},
				{-1. / 4., -1. / 4., 1},
			},
		},
		{
			2, -2, 1, -1, 1, 4,
			[]Vec3[float64]{
				{1, 1, -1},
				{-1, -1, -1},
				{1, 1, -4},
				{-1, -1, -4},
			},
			[]Vec3[float64]{
				{.5, 1, -1},
				{-.5, -1, -1},
				{.5 / 4., 1. / 4., 1},
				{-.5 / 4., -1. / 4., 1},
			},
		},
	}

	for _, c := range cases {
		p := Mat4Perspective(c.r, c.l, c.t, c.b, c.n, c.f)

		for i := range c.world {
			expected := c.screen[i]
			actual := p.TransformVec3(c.world[i], 1)

			if !vec3Identical(expected, actual) {
				t.Errorf("expected: %v, actual: %v", expected, actual)
			}
		}
	}
}

func TestMat4Translation(t *testing.T) {
	cases := []struct {
		x, y, z float64
		result  Mat4[float64]
	}{
		{
			0, 0, 0,
			Mat4Identity[float64](),
		},
		{
			1.23, -2.34, 23.12,
			Mat4[float64]{
				1, 0, 0, 1.23,
				0, 1, 0, -2.34,
				0, 0, 1, 23.12,
				0, 0, 0, 1,
			},
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := Mat4Translation[float64](Vec3[float64]{c.x, c.y, c.z})
		if !mat4Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestMat4Product(t *testing.T) {
	cases := []struct {
		a, b, result Mat4[float64]
	}{
		{
			Mat4Identity[float64](),
			Mat4Identity[float64](),
			Mat4Identity[float64](),
		},
		{
			Mat4[float64]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			Mat4[float64]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			Mat4[float64]{
				56, 62, 68, 74,
				152, 174, 196, 218,
				248, 286, 324, 362,
				344, 398, 452, 506,
			},
		},
		{
			Mat4[float64]{
				5, 7, 9, 10,
				2, 3, 3, 8,
				8, 10, 2, 3,
				3, 3, 4, 8,
			},
			Mat4[float64]{
				3, 10, 12, 18,
				12, 1, 4, 9,
				9, 10, 12, 2,
				3, 12, 4, 10,
			},
			Mat4[float64]{
				210, 267, 236, 271,
				93, 149, 104, 149,
				171, 146, 172, 268,
				105, 169, 128, 169,
			},
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.a.Product(c.b)

		if !mat4Identical(expected, actual) {
			t.Errorf("a: %v, b: %v, expected: %v, got: %v", c.a, c.b, expected, actual)
		}
	}
}

func TestMat4RotationX(t *testing.T) {
	f707 := math.Sqrt(0.5)

	cases := []struct {
		rad    float64
		result Mat4[float64]
	}{
		{
			0,
			Mat4Identity[float64](),
		},
		{
			math.Pi / 4,
			Mat4[float64]{
				1, 0, 0, 0,
				0, f707, -f707, 0,
				0, f707, f707, 0,
				0, 0, 0, 1,
			},
		},
		{
			-3 * math.Pi / 4,
			Mat4[float64]{
				1, 0, 0, 0,
				0, -f707, f707, 0,
				0, -f707, -f707, 0,
				0, 0, 0, 1,
			},
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := Mat4RotationX(c.rad)

		if !mat4Identical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}

func TestMat4RotationY(t *testing.T) {
	f707 := math.Sqrt(0.5)

	cases := []struct {
		rad    float64
		result Mat4[float64]
	}{
		{0, Mat4Identity[float64]()},
		{
			float64(math.Pi / 4),
			Mat4[float64]{
				f707, 0, f707, 0,
				0, 1, 0, 0,
				-f707, 0, f707, 0,
				0, 0, 0, 1,
			},
		},
		{
			-3 * float64(math.Pi / 4),
			Mat4[float64]{
				-f707, 0, -f707, 0,
				0, 1, 0, 0,
				f707, 0, -f707, 0,
				0, 0, 0, 1,
			},
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := Mat4RotationY(c.rad)

		if !mat4Identical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}

func TestMat4RotationZ(t *testing.T) {
	f707 := math.Sqrt(0.5)

	cases := []struct {
		rad    float64
		result Mat4[float64]
	}{
		{
			0,
			Mat4Identity[float64](),
		},
		{
			float64(math.Pi / 4),
			Mat4[float64]{
				f707, -f707, 0, 0,
				f707, f707, 0, 0,
				0, 0, 1, 0,
				0, 0, 0, 1,
			},
		},
		{
			-3 * float64(math.Pi / 4),
			Mat4[float64]{
				-f707, f707, 0, 0,
				-f707, -f707, 0, 0,
				0, 0, 1, 0,
				0, 0, 0, 1,
			},
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := Mat4RotationZ(c.rad)

		if !mat4Identical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}

func TestMat4RollPitchYaw(t *testing.T) {
	f707 := math.Sqrt(0.5)

	cases := []struct {
		r, p, y float64
		result  Mat4[float64]
	}{
		{
			0, 0, 0,
			Mat4Identity[float64](),
		},
		{
			float64(math.Pi / 4), 0, 0,
			Mat4RotationZ(float64(math.Pi / 4)),
		},
		{
			0, float64(math.Pi / 4), 0,
			Mat4RotationX(float64(math.Pi / 4)),
		},
		{
			0, 0, float64(math.Pi / 4),
			Mat4RotationY(float64(math.Pi / 4)),
		},
		{
			float64(math.Pi / 2), float64(math.Pi / 4), float64(math.Pi / 2),
			Mat4[float64]{
				f707, 0, f707, 0,
				f707, 0, -f707, 0,
				0, 1, 0, 0,
				0, 0, 0, 1,
			},
		},
		{
			-float64(math.Pi / 2), -float64(math.Pi / 2), float64(math.Pi / 4),
			Mat4[float64]{
				f707, f707, 0, 0,
				0, 0, 1, 0,
				f707, -f707, 0, 0,
				0, 0, 0, 1,
			},
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := Mat4RollPitchYaw(c.r, c.p, c.y)

		if !mat4Identical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}

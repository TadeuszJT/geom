package geomTest

import (
	. "github.com/tadeuszjt/geom/32"
	"math"
	"testing"
)

func mat4Identical(a, b Mat4) bool {
	for i := range a {
		if !floatIdentical(a[i], b[i]) {
			return false
		}
	}
	return true
}

func TestMat4Identical(t *testing.T) {
	cases := []struct {
		a, b   Mat4
		result bool
	}{
		{Mat4Identity(), Mat4Identity(), true},
		{
			Mat4{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			Mat4{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			true,
		},
		{
			Mat4{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15.0001},
			Mat4{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			false,
		},
		{
			Mat4{0, 1, 2, 3, 4, 5, 6, nan, 8, 9, 10, 11, 12, 13, 14, 15},
			Mat4{0, 1, 2, 3, 4, 5, 6, -nan, 8, 9, 10, 11, 12, 13, 14, 15},
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

func TestMat4Translation(t *testing.T) {
	cases := []struct {
		x, y, z float32
		result  Mat4
	}{
		{
			0, 0, 0,
			Mat4Identity(),
		},
		{
			1.23, -2.34, 23.12,
			Mat4{
				1, 0, 0, 1.23,
				0, 1, 0, -2.34,
				0, 0, 1, 23.12,
				0, 0, 0, 1,
			},
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := Mat4Translation(Vec3{c.x, c.y, c.z})
		if !mat4Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestMat4Product(t *testing.T) {
	cases := []struct {
		a, b, result Mat4
	}{
		{
			Mat4Identity(),
			Mat4Identity(),
			Mat4Identity(),
		},
		{
			Mat4{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			Mat4{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			Mat4{
				56, 62, 68, 74,
				152, 174, 196, 218,
				248, 286, 324, 362,
				344, 398, 452, 506,
			},
		},
		{
			Mat4{
				5, 7, 9, 10,
				2, 3, 3, 8,
				8, 10, 2, 3,
				3, 3, 4, 8,
			},
			Mat4{
				3, 10, 12, 18,
				12, 1, 4, 9,
				9, 10, 12, 2,
				3, 12, 4, 10,
			},
			Mat4{
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
	f707 := float32(math.Sqrt(0.5))

	cases := []struct {
		rad    float32
		result Mat4
	}{
		{
			0,
			Mat4Identity(),
		},
		{
			float32(math.Pi / 4),
			Mat4{
				1, 0, 0, 0,
				0, f707, -f707, 0,
				0, f707, f707, 0,
				0, 0, 0, 1,
			},
		},
		{
			float32(-3 * math.Pi / 4),
			Mat4{
				1, 0, 0, 0,
				0, -f707, f707, 0,
				0, -f707, -f707, 0,
				0, 0, 0, 1,
			},
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := Mat4RotationX(MakeAngle(c.rad))

		if !mat4Identical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}

func TestMat4RotationY(t *testing.T) {
	f707 := float32(math.Sqrt(0.5))

	cases := []struct {
		rad    float32
		result Mat4
	}{
		{
			0,
			Mat4Identity(),
		},
		{
			float32(math.Pi / 4),
			Mat4{
				f707, 0, f707, 0,
				0, 1, 0, 0,
				-f707, 0, f707, 0,
				0, 0, 0, 1,
			},
		},
		{
			float32(-3 * math.Pi / 4),
			Mat4{
				-f707, 0, -f707, 0,
				0, 1, 0, 0,
				f707, 0, -f707, 0,
				0, 0, 0, 1,
			},
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := Mat4RotationY(MakeAngle(c.rad))

		if !mat4Identical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}

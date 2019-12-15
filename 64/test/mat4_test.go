package geomTest

import (
	. "github.com/tadeuszjt/geom/64"
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

package geomTest

import (
	. "github.com/tadeuszjt/geom/generic"
	"math"
	"testing"
)

func vec3Identical(a, b Vec3[float64]) bool {
	return floatIdentical(a.X, b.X) &&
		floatIdentical(a.Y, b.Y) &&
		floatIdentical(a.Z, b.Z)
}

func TestVec3Vec2(t *testing.T) {
	for _, c := range []struct {
		Vec3[float64]
		Vec2[float64]
	}{
		{Vec3[float64]{0, 0, 0}, Vec2[float64]{0, 0}},
		{Vec3[float64]{1, 2, 3}, Vec2[float64]{1, 2}},
		{Vec3[float64]{-1, -2, -3}, Vec2[float64]{-1, -2}},
		{Vec3[float64]{nan, nInf, pInf}, Vec2[float64]{nan, nInf}},
	} {
		expected := c.Vec2
		actual := c.Vec3.Vec2()
		if !vec2Identical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}

func TestVec3Dot(t *testing.T) {
	cases := []struct {
		a, b   Vec3[float64]
		result float64
	}{
		{Vec3[float64]{}, Vec3[float64]{}, 0},
		{Vec3[float64]{1, 2, 3}, Vec3[float64]{4, 5, 6}, 32},
		{Vec3[float64]{0, 0, 0}, Vec3[float64]{4, 5, 6}, 0},
		{Vec3[float64]{-1, -2, -3}, Vec3[float64]{4, 5, 6}, -32},
		{Vec3[float64]{-1, 2, -3}, Vec3[float64]{4, 5, 6}, -12},
		{Vec3[float64]{-1, nan, -3}, Vec3[float64]{4, 5, 6}, nan},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.a.Dot(c.b)
		if !floatIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestVec3Times(t *testing.T) {
	cases := []struct {
		a, b, result Vec3[float64]
	}{
		{Vec3[float64]{}, Vec3[float64]{}, Vec3[float64]{}},
		{Vec3[float64]{1, 2, 3}, Vec3[float64]{4, 5, 6}, Vec3[float64]{4, 10, 18}},
		{Vec3[float64]{-1, 2, -3}, Vec3[float64]{4, -5, 6}, Vec3[float64]{-4, -10, -18}},
		{Vec3[float64]{nan, pInf, nInf}, Vec3[float64]{-4, -5, -6}, Vec3[float64]{nan, nInf, pInf}},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.a.Times(c.b)
		if !vec3Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestVec3ScaledBy(t *testing.T) {
	cases := []struct {
		scalar    float64
		v, result Vec3[float64]
	}{
		{0, Vec3[float64]{}, Vec3[float64]{}},
		{0, Vec3[float64]{1, 2, 3}, Vec3[float64]{0, 0, 0}},
		{1, Vec3[float64]{1, 2, 3}, Vec3[float64]{1, 2, 3}},
		{-1, Vec3[float64]{1, 2, 3}, Vec3[float64]{-1, -2, -3}},
		{100, Vec3[float64]{1, 2, 3}, Vec3[float64]{100, 200, 300}},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.v.ScaledBy(c.scalar)
		if !vec3Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestVec3Normal(t *testing.T) {
	f707 := float64(math.Sqrt(.5))

	cases := []struct {
		v, result Vec3[float64]
	}{
		{Vec3[float64]{}, Vec3[float64]{}},
		{Vec3[float64]{2, 0, 0}, Vec3[float64]{1, 0, 0}},
		{Vec3[float64]{2, 2, 0}, Vec3[float64]{f707, f707, 0}},
		{Vec3[float64]{0, 4, -4}, Vec3[float64]{0, f707, -f707}},
		{Vec3[float64]{-3, -4, -5}, Vec3[float64]{-3. / (f707 * 10), -4. / (f707 * 10), -5. / (f707 * 10)}},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.v.Normal()

		if !vec3Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestVec3Ori2(t *testing.T) {
	cases := []struct {
		v Vec3[float64]
		o Ori2[float64]
	}{
		{Vec3[float64]{}, Ori2[float64]{}},
		{Vec3[float64]{1, 2, 3}, Ori2[float64]{1, 2, 3}},
		{Vec3[float64]{-1, -2, -3}, Ori2[float64]{-1, -2, -3}},
		{Vec3[float64]{nan, pInf, nInf}, Ori2[float64]{nan, pInf, nInf}},
	}

	for _, c := range cases {
		expected := c.o
		actual := c.v.Ori2()
		if !ori2Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestVec3Len(t *testing.T) {
	cases := []struct {
		v Vec3[float64]
		l float64
	}{
		{Vec3[float64]{}, 0},
		{Vec3[float64]{1, 0, 0}, 1},
		{Vec3[float64]{0, 1, 0}, 1},
		{Vec3[float64]{0, 0, 1}, 1},
		{Vec3[float64]{3, 4, 0}, 5},
		{Vec3[float64]{0, -3, 4}, 5},
		{Vec3[float64]{3, 4, 5}, math.Sqrt(3*3 + 4*4 + 5*5)},
	}

	for _, c := range cases {
		expected := c.l
		actual := c.v.Len()
		if !floatIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestVec3Pitch(t *testing.T) {
	cases := []struct {
		v Vec3[float64]
		p float64
	}{
		{Vec3[float64]{}, 0},
		{Vec3[float64]{3, 4, 0}, -math.Atan(4.0 / 3.0)},
		{Vec3[float64]{3, 0, 2}, 0},
		{Vec3[float64]{0, 3, -3}, -float64(math.Pi / 4)},
		{Vec3[float64]{-4, nInf, 3}, float64(math.Pi / 2)},
		{Vec3[float64]{-4, pInf, 2}, -float64(math.Pi / 2)},
	}

	for _, c := range cases {
		expected := c.p
		actual := c.v.Pitch()
		if !floatIdentical(expected, actual) {
			t.Errorf("v: %v: expected: %v, actual: %v", c.v, expected, actual)
		}
	}
}

func TestVec3Yaw(t *testing.T) {
	cases := []struct {
		v Vec3[float64]
		y float64
	}{
		{Vec3[float64]{}, 0},
		{Vec3[float64]{3, -32, 3}, float64(math.Pi / 4)},
		{Vec3[float64]{3, -6, -3}, 3 * float64(math.Pi / 4)},
		{Vec3[float64]{-3, -6, 3}, -float64(math.Pi / 4)},
		{Vec3[float64]{-3, -3, -3}, -3 * float64(math.Pi / 4)},
		{Vec3[float64]{0, nInf, -3}, float64(math.Pi)},
	}

	for _, c := range cases {
		expected := c.y
		actual := c.v.Yaw()
		if !floatIdentical(expected, actual) {
			t.Errorf("v: %v: expected: %v, actual: %v", c.v, expected, actual)
		}
	}
}

func TestVec3Rand(t *testing.T) {
	const num = 1000

	cube := Cuboid[float64]{Min: Vec3[float64]{-1, -2, -3}, Max: Vec3[float64]{4, 5, 6}}
	vecs := [num]Vec3[float64]{}

	for i := range vecs {
		vecs[i] = Vec3Rand(cube)
	}

	for i := range vecs {
		if !cube.Contains(vecs[i]) {
			t.Errorf("%v does not contain %v", cube, vecs[i])
		}
	}
}

func TestVec3NormPitchYaw(t *testing.T) {
	f707 := math.Sqrt(0.5)

	cases := []struct {
		p, y float64
		v    Vec3[float64]
	}{
		{0, 0, Vec3[float64]{0, 0, 1}},
		{float64(math.Pi / 4), 0, Vec3[float64]{0, -f707, f707}},
		{-float64(math.Pi / 4), 0, Vec3[float64]{0, f707, f707}},
		{-float64(math.Pi / 4), float64(math.Pi / 2), Vec3[float64]{f707, f707, 0}},
	}

	for _, c := range cases {
		expected := c.v
		actual := Vec3NormPitchYaw(c.p, c.y)
		if !vec3Identical(expected, actual) {
			t.Errorf("v: %v: expected: %v, actual: %v", c.v, expected, actual)
		}
	}
}

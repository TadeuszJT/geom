package geomTest

import (
	. "github.com/tadeuszjt/geom/32"
	"math"
	"testing"
)

func vec3Identical(a, b Vec3) bool {
	return floatIdentical(a.X, b.X) &&
		floatIdentical(a.Y, b.Y) &&
		floatIdentical(a.Z, b.Z)
}

func TestVec3Vec2(t *testing.T) {
	for _, c := range []struct {
		Vec3
		Vec2
	}{
		{Vec3{0, 0, 0}, Vec2{0, 0}},
		{Vec3{1, 2, 3}, Vec2{1, 2}},
		{Vec3{-1, -2, -3}, Vec2{-1, -2}},
		{Vec3{nan, nInf, pInf}, Vec2{nan, nInf}},
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
		a, b   Vec3
		result float32
	}{
		{Vec3{}, Vec3{}, 0},
		{Vec3{1, 2, 3}, Vec3{4, 5, 6}, 32},
		{Vec3{0, 0, 0}, Vec3{4, 5, 6}, 0},
		{Vec3{-1, -2, -3}, Vec3{4, 5, 6}, -32},
		{Vec3{-1, 2, -3}, Vec3{4, 5, 6}, -12},
		{Vec3{-1, nan, -3}, Vec3{4, 5, 6}, nan},
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
		a, b, result Vec3
	}{
		{Vec3{}, Vec3{}, Vec3{}},
		{Vec3{1, 2, 3}, Vec3{4, 5, 6}, Vec3{4, 10, 18}},
		{Vec3{-1, 2, -3}, Vec3{4, -5, 6}, Vec3{-4, -10, -18}},
		{Vec3{nan, pInf, nInf}, Vec3{-4, -5, -6}, Vec3{nan, nInf, pInf}},
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
		scalar    float32
		v, result Vec3
	}{
		{0, Vec3{}, Vec3{}},
		{0, Vec3{1, 2, 3}, Vec3{0, 0, 0}},
		{1, Vec3{1, 2, 3}, Vec3{1, 2, 3}},
		{-1, Vec3{1, 2, 3}, Vec3{-1, -2, -3}},
		{100, Vec3{1, 2, 3}, Vec3{100, 200, 300}},
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
    f707 := float32(math.Sqrt(.5))

	cases := []struct {
		v, result Vec3
	}{
        {Vec3{}, Vec3{}},
        {Vec3{2, 0, 0}, Vec3{1, 0, 0}},
        {Vec3{2, 2, 0}, Vec3{f707, f707, 0}},
        {Vec3{0, 4, -4}, Vec3{0, f707, -f707}},
        {Vec3{-3, -4, -5}, Vec3{-3./(f707*10), -4./(f707*10), -5./(f707*10)}},
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
		v Vec3
		o Ori2
	}{
		{Vec3{}, Ori2{}},
		{Vec3{1, 2, 3}, Ori2{1, 2, 3}},
		{Vec3{-1, -2, -3}, Ori2{-1, -2, -3}},
		{Vec3{nan, pInf, nInf}, Ori2{nan, pInf, MakeAngle(nInf)}},
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
		v Vec3
		l float32
	}{
		{Vec3{}, 0},
		{Vec3{1, 0, 0}, 1},
		{Vec3{0, 1, 0}, 1},
		{Vec3{0, 0, 1}, 1},
		{Vec3{3, 4, 0}, 5},
		{Vec3{0, -3, 4}, 5},
		{Vec3{3, 4, 5}, float32(math.Sqrt(3*3 + 4*4 + 5*5))},
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
		v Vec3
		p Angle
	}{
		{Vec3{}, 0},
		{Vec3{3, 4, 0}, -Angle(math.Atan(4.0 / 3.0))},
		{Vec3{3, 0, 2}, 0},
		{Vec3{0, 3, -3}, -Angle45Deg},
		{Vec3{-4, nInf, 3}, Angle90Deg},
		{Vec3{-4, pInf, 2}, -Angle90Deg},
	}

	for _, c := range cases {
		expected := float32(c.p)
		actual := float32(c.v.Pitch())
		if !floatIdentical(expected, actual) {
			t.Errorf("v: %v: expected: %v, actual: %v", c.v, expected, actual)
		}
	}
}

func TestVec3Yaw(t *testing.T) {
	cases := []struct {
		v Vec3
		y Angle
	}{
		{Vec3{}, 0},
		{Vec3{3, -32, 3}, Angle45Deg},
		{Vec3{3, -6, -3}, 3 * Angle45Deg},
		{Vec3{-3, -6, 3}, -Angle45Deg},
		{Vec3{-3, -3, -3}, -3 * Angle45Deg},
		{Vec3{0, nInf, -3}, AnglePi},
	}

	for _, c := range cases {
		expected := float32(c.y)
		actual := float32(c.v.Yaw())
		if !floatIdentical(expected, actual) {
			t.Errorf("v: %v: expected: %v, actual: %v", c.v, expected, actual)
		}
	}
}

func TestVec3Rand(t *testing.T) {
    const num = 1000

    cube := Cuboid{Min: Vec3{-1, -2, -3}, Max: Vec3{4, 5, 6}}
    vecs := [num]Vec3{}

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
	f707 := float32(math.Sqrt(0.5))

	cases := []struct {
		p, y Angle
		v    Vec3
	}{
		{0, 0, Vec3{0, 0, 1}},
		{Angle45Deg, 0, Vec3{0, -f707, f707}},
		{-Angle45Deg, 0, Vec3{0, f707, f707}},
		{-Angle45Deg, Angle90Deg, Vec3{f707, f707, 0}},
	}

	for _, c := range cases {
		expected := c.v
		actual := Vec3NormPitchYaw(c.p, c.y)
		if !vec3Identical(expected, actual) {
			t.Errorf("v: %v: expected: %v, actual: %v", c.v, expected, actual)
		}
	}
}

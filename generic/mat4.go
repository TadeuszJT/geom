package geom

import "math"

type Mat4[T Num] [16]T

func Mat4Identity[T Num]() Mat4[T] {
	return Mat4[T]{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

/* Transforms a homogenous coordinate
 * Using w=1 will perform a normal 'affine' transformation
 */
func (m Mat4[T]) TransformVec3(v Vec3[T], w T) Vec3[T] {
	a := m[0]*v.X + m[1]*v.Y + m[2]*v.Z + m[3]*w
	b := m[4]*v.X + m[5]*v.Y + m[6]*v.Z + m[7]*w
	c := m[8]*v.X + m[9]*v.Y + m[10]*v.Z + m[11]*w
	d := m[12]*v.X + m[13]*v.Y + m[14]*v.Z + m[15]*w

	return Vec3[T]{X: a / d, Y: b / d, Z: c / d}
}

/* Must be used with a homogenous coordinate transformation such as OpenGl or TransformVec3().
 * Looks down Z axis
 */
func Mat4Perspective[T Num](r, l, t, b, n, f T) Mat4[T] {
	return Mat4[T]{
		2 * n / (r - l), 0, (r + l) / (r - l), 0,
		0, 2 * n / (t - b), (t + b) / (t - b), 0,
		0, 0, -(f + n) / (f - n), -2 * f * n / (f - n),
		0, 0, -1, 0,
	}
}

func Mat4Translation[T Num](v Vec3[T]) Mat4[T] {
	return Mat4[T]{
		1, 0, 0, v.X,
		0, 1, 0, v.Y,
		0, 0, 1, v.Z,
		0, 0, 0, 1,
	}
}

func Mat4RotationX[T Num](rad T) Mat4[T] {
	s := T(math.Sin(float64(rad)))
	c := T(math.Cos(float64(rad)))

	return Mat4[T]{
		1, 0, 0, 0,
		0, c, -s, 0,
		0, s, c, 0,
		0, 0, 0, 1,
	}
}

func Mat4RotationY[T Num](rad T) Mat4[T] {
	s := T(math.Sin(float64(rad)))
	c := T(math.Cos(float64(rad)))

	return Mat4[T]{
		c, 0, s, 0,
		0, 1, 0, 0,
		-s, 0, c, 0,
		0, 0, 0, 1,
	}
}

func Mat4RotationZ[T Num](rad T) Mat4[T] {
	s := T(math.Sin(float64(rad)))
	c := T(math.Cos(float64(rad)))

	return Mat4[T]{
		c, -s, 0, 0,
		s, c, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func Mat4Scalar[T Num](x, y, z T) Mat4[T] {
	return Mat4[T]{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	}
}

func Mat4RollPitchYaw[T Num](r, p, y T) Mat4[T] {
	rz := Mat4RotationZ(r)
	rx := Mat4RotationX(p)
	ry := Mat4RotationY(y)
	return ry.Product(rx).Product(rz)
}

func (a Mat4[T]) Product(b Mat4[T]) (m Mat4[T]) {
	for i := range m {
		r, c := i/4, i%4
		for j := 0; j < 4; j++ {
			m[i] += a[r*4+j] * b[j*4+c]
		}
	}

	return
}

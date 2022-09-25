package geom

import "math"

type Mat4 [16]float32

func Mat4Identity() Mat4 {
	return Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

/* Transforms a homogenous coordinate
 * Using w=1 will perform a normal 'affine' transformation
 */
func (m Mat4) TransformVec3(v Vec3, w float32) Vec3 {
	a := m[0]*v.X + m[1]*v.Y + m[2]*v.Z + m[3]*w
	b := m[4]*v.X + m[5]*v.Y + m[6]*v.Z + m[7]*w
	c := m[8]*v.X + m[9]*v.Y + m[10]*v.Z + m[11]*w
	d := m[12]*v.X + m[13]*v.Y + m[14]*v.Z + m[15]*w

	return Vec3{X: a / d, Y: b / d, Z: c / d}
}

/* Must be used with a homogenous coordinate transformation such as OpenGl or TransformVec3().
 * Looks down Z axis
 */
func Mat4Perspective(r, l, t, b, n, f float32) Mat4 {
	return Mat4{
		2 * n / (r - l), 0, (r + l) / (r - l), 0,
		0, 2 * n / (t - b), (t + b) / (t - b), 0,
		0, 0, -(f + n) / (f - n), -2 * f * n / (f - n),
		0, 0, -1, 0,
	}
}

func Mat4Translation(v Vec3) Mat4 {
	return Mat4{
		1, 0, 0, v.X,
		0, 1, 0, v.Y,
		0, 0, 1, v.Z,
		0, 0, 0, 1,
	}
}

func Mat4RotationX(rad Angle) Mat4 {
	s := float32(math.Sin(float64(rad)))
	c := float32(math.Cos(float64(rad)))

	return Mat4{
		1, 0, 0, 0,
		0, c, -s, 0,
		0, s, c, 0,
		0, 0, 0, 1,
	}
}

func Mat4RotationY(rad Angle) Mat4 {
	s := float32(math.Sin(float64(rad)))
	c := float32(math.Cos(float64(rad)))

	return Mat4{
		c, 0, s, 0,
		0, 1, 0, 0,
		-s, 0, c, 0,
		0, 0, 0, 1,
	}
}

func Mat4RotationZ(rad Angle) Mat4 {
	s := float32(math.Sin(float64(rad)))
	c := float32(math.Cos(float64(rad)))

	return Mat4{
		c, -s, 0, 0,
		s, c, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func Mat4Scalar(x, y, z float32) Mat4 {
	return Mat4{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	}
}

func Mat4RollPitchYaw(r, p, y Angle) Mat4 {
	rz := Mat4RotationZ(r)
	rx := Mat4RotationX(p)
	ry := Mat4RotationY(y)
	return ry.Product(rx).Product(rz)
}

func (a Mat4) Product(b Mat4) (m Mat4) {
	for i := range m {
		r, c := i/4, i%4
		for j := 0; j < 4; j++ {
			m[i] += a[r*4+j] * b[j*4+c]
		}
	}

	return
}

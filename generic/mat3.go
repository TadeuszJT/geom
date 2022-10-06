package geom

import (
	"math"
)

type Mat3[T Num] [9]T

func Mat3Identity[T Num]() Mat3[T] {
	return Mat3[T]{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
}

func Mat3Scalar[T Num](x, y T) Mat3[T] {
	return Mat3[T]{
		x, 0, 0,
		0, y, 0,
		0, 0, 1,
	}
}

func Mat3Translation[T Num](v Vec2[T]) Mat3[T] {
	return Mat3[T]{
		1, 0, v.X,
		0, 1, v.Y,
		0, 0, 1,
	}
}

func Mat3Rotation[T Num](theta T) Mat3[T] {
	c := T(math.Cos(float64(theta)))
	s := T(math.Sin(float64(theta)))
	return Mat3[T]{
		c, -s, 0,
		s, c, 0,
		0, 0, 1,
	}
}

func (a Mat3[T]) Product(b Mat3[T]) Mat3[T] {
	return Mat3[T]{
		a[0]*b[0] + a[1]*b[3] + a[2]*b[6],
		a[0]*b[1] + a[1]*b[4] + a[2]*b[7],
		a[0]*b[2] + a[1]*b[5] + a[2]*b[8],

		a[3]*b[0] + a[4]*b[3] + a[5]*b[6],
		a[3]*b[1] + a[4]*b[4] + a[5]*b[7],
		a[3]*b[2] + a[4]*b[5] + a[5]*b[8],

		a[6]*b[0] + a[7]*b[3] + a[8]*b[6],
		a[6]*b[1] + a[7]*b[4] + a[8]*b[7],
		a[6]*b[2] + a[7]*b[5] + a[8]*b[8],
	}
}

func (m Mat3[T]) TimesVec2(v Vec2[T], bias T) Vec3[T] {
	return Vec3[T]{
		X: m[0]*v.X + m[1]*v.Y + m[2]*bias,
		Y: m[3]*v.X + m[4]*v.Y + m[5]*bias,
		Z: m[6]*v.X + m[7]*v.Y + m[8]*bias,
	}
}

func Mat3Camera2D[T Num](camera, display Rect[T]) Mat3[T] {
	sx := display.Width() / camera.Width()
	sy := display.Height() / camera.Height()

	tx := display.Min.X - sx*camera.Min.X
	ty := display.Min.Y - sy*camera.Min.Y

	return Mat3[T]{
		sx, 0, tx,
		0, sy, ty,
		0, 0, 1,
	}
}

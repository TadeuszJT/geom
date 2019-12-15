package geom

import "math"

/*
Conventions:
Right-Handed coordinate system
	Rotation around thumb axis, fingers curled
Column-Major
	P' = Ry*Rx*T*P - does Ry first

	| Xx Yx Zx Tx |   | X |
	| Xy Yy Zy Ty | * | Y |
	| Xz Yz Zz Tz |   | Z |
	| 0  0  0  1  |   | 1 |

*/

type Mat4 [16]float32

func Mat4Identity() Mat4 {
	return Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func Mat4Perspective(display Rect, displayNear, displayFar, fov, near, far float32) Mat4 {
	arInv := display.Height() / display.Width()
	d := float32(math.Atan((float64(fov) * math.Pi) / (180 * 2)))
	a := (-near - far) / (near - far)
	b := (2 * far * near) / (near - far)

	return Mat4{
		-d * arInv, 0, 0, 0,
		0, d, 0, 0,
		0, 0, a, b,
		0, 0, 1, 0,
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

func (a Mat4) Product(b Mat4) (m Mat4) {
	for i := range m {
		r, c := i/4, i%4
		for j := 0; j < 4; j++ {
			m[i] += a[r*4+j] * b[j*4+c]
		}
	}

	return
}

package geom

import "math"

type Ori2 struct {
	X, Y  float32
	Theta Angle
}

func MakeOri2(pos Vec2, theta Angle) Ori2 {
    return Ori2{pos.X, pos.Y, theta}
}

func (o Ori2) Vec2() Vec2 {
	return Vec2{o.X, o.Y}
}

func (o Ori2) Vec3() Vec3 {
	return Vec3{o.X, o.Y, float32(o.Theta)}
}

func (a *Ori2) PlusEquals(b Ori2) {
	a.X += b.X
	a.Y += b.Y
	a.Theta = MakeAngle(float32(a.Theta + b.Theta))
}

func (o Ori2) Mat3Transform() Mat3 {
	sin := float32(math.Sin(float64(o.Theta)))
	cos := float32(math.Cos(float64(o.Theta)))
	return Mat3{
		cos, -sin, o.X,
		sin, cos, o.Y,
		0, 0, 1,
	}
}

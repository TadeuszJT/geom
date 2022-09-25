package geom

import "math"

type Ori2 struct {
	X, Y, Theta float64
}

func (o Ori2) Vec2() Vec2 {
	return Vec2{o.X, o.Y}
}

func (o Ori2) Vec3() Vec3 {
	return Vec3{o.X, o.Y, o.Theta}
}

func (a Ori2) Times(b Ori2) Ori2 {
	return Ori2{a.X * b.X, a.Y * b.Y, a.Theta * b.Theta}
}

func (a Ori2) Dot(b Ori2) float64 {
    return a.X*b.X + a.Y*b.Y + a.Theta*b.Theta
}

func (a *Ori2) PlusEquals(b Ori2) {
	a.X += b.X
	a.Y += b.Y
	a.Theta += b.Theta
}

func (o Ori2) ScaledBy(f float64) Ori2 {
	return Ori2{f * o.X, f * o.Y, f * o.Theta}
}

func (o Ori2) Mat3Transform() Mat3 {
	sin := math.Sin(o.Theta)
	cos := math.Cos(o.Theta)
	return Mat3{
		cos, -sin, o.X,
		sin, cos, o.Y,
		0, 0, 1,
	}
}

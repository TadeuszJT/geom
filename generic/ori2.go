package geom

import "math"

type Ori2[T Num] struct {
	X, Y, Theta  T
}

func Ori2Convert[A, B Num](a Ori2[A]) Ori2[B] {
    return Ori2[B]{B(a.X), B(a.Y), B(a.Theta)}
}

func MakeOri2[T Num](pos Vec2[T], theta T) Ori2[T] {
	return Ori2[T]{pos.X, pos.Y, theta}
}

func (o Ori2[T]) Vec2() Vec2[T] {
	return Vec2[T]{o.X, o.Y}
}

func (o Ori2[T]) Vec3() Vec3[T] {
	return Vec3[T]{o.X, o.Y, o.Theta}
}

func (o Ori2[T]) ScaledBy(f T) Ori2[T] {
	return Ori2[T]{o.X * f, o.Y * f, o.Theta * f}
}

func (a *Ori2[T]) PlusEquals(b Ori2[T]) {
	a.X += b.X
	a.Y += b.Y
	a.Theta += b.Theta
}

func (o Ori2[T]) Mat3Transform() Mat3[T] {
	sin := T(math.Sin(float64(o.Theta)))
	cos := T(math.Cos(float64(o.Theta)))
	return Mat3[T]{
		cos, -sin, o.X,
		sin, cos, o.Y,
		0, 0, 1,
	}
}

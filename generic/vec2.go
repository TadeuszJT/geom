package geom

import (
	"math"
	"math/rand"
)

type Vec2[T Num] struct {
	X, Y T
}

func Vec2Convert[A, B Num](v Vec2[A]) Vec2[B] {
	return Vec2[B]{B(v.X), B(v.Y)}
}

func (v Vec2[T]) Ori2() Ori2[T] {
	return Ori2[T]{v.X, v.Y, 0}
}

func (v Vec2[T]) Perpendicular() Vec2[T] {
    return Vec2[T]{-v.Y, v.X}
}

func (a Vec2[T]) Plus(b Vec2[T]) Vec2[T] {
	return Vec2[T]{a.X + b.X, a.Y + b.Y}
}

func (a Vec2[T]) Minus(b Vec2[T]) Vec2[T] {
	return Vec2[T]{a.X - b.X, a.Y - b.Y}
}

func (a Vec2[T]) Dot(b Vec2[T]) T {
	return a.X*b.X + a.Y*b.Y
}

func (a Vec2[T]) Cross(b Vec2[T]) T {
	return a.X*b.Y - a.Y*b.X
}

func (v Vec2[T]) ScaledBy(f T) Vec2[T] {
	return Vec2[T]{v.X * f, v.Y * f}
}

func (v Vec2[T]) RotatedBy(rad T) Vec2[T] {
	rad64 := float64(rad)
	sin := T(math.Sin(rad64))
	cos := T(math.Cos(rad64))

	return Vec2[T]{cos*v.X - sin*v.Y, sin*v.X + cos*v.Y}
}

func (v Vec2[T]) Len2() T {
	return v.X*v.X + v.Y*v.Y
}

func (v Vec2[T]) Len() T {
	return T(math.Sqrt(float64(v.Len2())))
}

func (v Vec2[T]) Theta() T {
	return T(math.Atan2(float64(v.Y), float64(v.X)))
}

func (v Vec2[T]) Normal() Vec2[T] {
	x64 := float64(v.X)
	y64 := float64(v.Y)

	len := v.Len()
	if len == 0 || math.IsNaN(x64) || math.IsNaN(y64) {
		return Vec2[T]{}
	}

	switch {
	case math.IsInf(x64, 0) && math.IsInf(y64, 0):
		return Vec2[T]{}
	case math.IsInf(x64, 1):
		return Vec2[T]{1, 0}
	case math.IsInf(x64, -1):
		return Vec2[T]{-1, 0}
	case math.IsInf(y64, 1):
		return Vec2[T]{0, 1}
	case math.IsInf(y64, -1):
		return Vec2[T]{0, -1}
	}

	return v.ScaledBy(1 / len)
}

func (a *Vec2[T]) PlusEquals(b Vec2[T]) {
	a.X += b.X
	a.Y += b.Y
}

func Vec2Rand[T Num](r Rect[T]) Vec2[T] {
	return Vec2[T]{
		T(rand.Float64())*r.Width() + r.Min.X,
		T(rand.Float64())*r.Height() + r.Min.Y,
	}
}

func Vec2RandNormal[T Num]() Vec2[T] {
	theta := rand.Float64() * 2 * math.Pi
	return Vec2[T]{
		T(math.Sin(theta)),
		T(math.Cos(theta)),
	}
}

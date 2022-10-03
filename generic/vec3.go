package geom

import (
	"math"
	"math/rand"
)

type Vec3[T Num] struct {
	X, Y, Z T
}

func (v Vec3[T]) Vec2() Vec2[T] {
	return Vec2[T]{v.X, v.Y}
}

func (a Vec3[T]) Dot(b Vec3[T]) T {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a Vec3[T]) Plus(b Vec3[T]) Vec3[T] {
	return Vec3[T]{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (a Vec3[T]) Minus(b Vec3[T]) Vec3[T] {
	return Vec3[T]{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func (a Vec3[T]) Times(b Vec3[T]) Vec3[T] {
	return Vec3[T]{a.X * b.X, a.Y * b.Y, a.Z * b.Z}
}

func (v Vec3[T]) ScaledBy(f T) Vec3[T] {
	return Vec3[T]{f * v.X, f * v.Y, f * v.Z}
}

func (v Vec3[T]) Normal() Vec3[T] {
	l := v.Len()
	if l == 0 {
		return Vec3[T]{}
	}
	return Vec3[T]{v.X / l, v.Y / l, v.Z / l}
}

func (v Vec3[T]) Ori2() Ori2[T] {
	return Ori2[T]{v.X, v.Y, v.Z}
}

func (v Vec3[T]) Len2() T {
	return v.Dot(v)
}

func (v Vec3[T]) Len() T {
	return T(math.Sqrt(float64(v.Dot(v))))
}

func (v Vec3[T]) Pitch() T {
	xzNorm := math.Sqrt(float64(v.X*v.X + v.Z*v.Z))
	return T(math.Atan2(float64(-v.Y), xzNorm))
}

func (v Vec3[T]) Yaw() T {
	return T(math.Atan2(float64(v.X), float64(v.Z)))
}

func (v *Vec3[T]) PlusEquals(b Vec3[T]) {
	v.X += b.X
	v.Y += b.Y
	v.Z += b.Z
}

func (v *Vec3[T]) MinusEquals(b Vec3[T]) {
	v.X -= b.X
	v.Y -= b.Y
	v.Z -= b.Z
}

func Vec3Rand[T Num](space Cuboid[T]) Vec3[T] {
	return Vec3[T]{
		X: T(rand.Float64())*space.Width() + space.Min.X,
		Y: T(rand.Float64())*space.Height() + space.Min.Y,
		Z: T(rand.Float64())*space.Depth() + space.Min.Z,
	}
}

func Vec3NormPitchYaw[T Num](pitch, yaw T) Vec3[T] {
	return Vec3[T]{
		X: T(math.Cos(float64(pitch)) * math.Sin(float64(yaw))),
		Y: -T(math.Sin(float64(pitch))),
		Z: T(math.Cos(float64(pitch)) * math.Cos(float64(yaw))),
	}
}

func Vec3NormRand[T Num]() Vec3[T] {
	// using equal-area-projection on sphere
	theta := rand.Float64() * 2 * math.Pi
	long := rand.Float64()*2 - 1

	return Vec3[T]{
		X: T(math.Sqrt(1-long*long) * math.Cos(theta)),
		Y: T(math.Sqrt(1-long*long) * math.Sin(theta)),
		Z: T(long),
	}
}

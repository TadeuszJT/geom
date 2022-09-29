package geom

import (
	"math"
	"math/rand"
)

type Vec3 struct {
	X, Y, Z float32
}

func (v Vec3) Vec2() Vec2 {
	return Vec2{v.X, v.Y}
}

func (a Vec3) Dot(b Vec3) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a Vec3) Plus(b Vec3) Vec3 {
	return Vec3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (a Vec3) Minus(b Vec3) Vec3 {
	return Vec3{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func (a Vec3) Times(b Vec3) Vec3 {
	return Vec3{a.X * b.X, a.Y * b.Y, a.Z * b.Z}
}

func (v Vec3) ScaledBy(f float32) Vec3 {
	return Vec3{f * v.X, f * v.Y, f * v.Z}
}

func (v Vec3) Normal() Vec3 {
	l := v.Len()
	if l == 0 {
		return Vec3{}
	}
	return Vec3{v.X / l, v.Y / l, v.Z / l}
}

func (v Vec3) Ori2() Ori2 {
	return Ori2{v.X, v.Y, MakeAngle(v.Z)}
}

func (v Vec3) Len2() float32 {
	return v.Dot(v)
}

func (v Vec3) Len() float32 {
	return float32(math.Sqrt(float64(v.Dot(v))))
}

func (v Vec3) Pitch() Angle {
	xzNorm := math.Sqrt(float64(v.X*v.X + v.Z*v.Z))
	return Angle(math.Atan2(float64(-v.Y), xzNorm))
}

func (v Vec3) Yaw() Angle {
	return Angle(math.Atan2(float64(v.X), float64(v.Z)))
}

func (v *Vec3) PlusEquals(b Vec3) {
	v.X += b.X
	v.Y += b.Y
	v.Z += b.Z
}

func (v *Vec3) MinusEquals(b Vec3) {
	v.X -= b.X
	v.Y -= b.Y
	v.Z -= b.Z
}

func Vec3Rand(space Cuboid) Vec3 {
	return Vec3{
		X: float32(rand.Float64())*space.Width() + space.Min.X,
		Y: float32(rand.Float64())*space.Height() + space.Min.Y,
		Z: float32(rand.Float64())*space.Depth() + space.Min.Z,
	}
}

func Vec3NormPitchYaw(pitch, yaw Angle) Vec3 {
	return Vec3{
		X: float32(math.Cos(float64(pitch)) * math.Sin(float64(yaw))),
		Y: -float32(math.Sin(float64(pitch))),
		Z: float32(math.Cos(float64(pitch)) * math.Cos(float64(yaw))),
	}
}

func Vec3NormRand() Vec3 {
	// using equal-area-projection on sphere
	theta := rand.Float64() * 2 * math.Pi
	long := rand.Float64()*2 - 1

	return Vec3{
		X: float32(math.Sqrt(1-long*long) * math.Cos(theta)),
		Y: float32(math.Sqrt(1-long*long) * math.Sin(theta)),
		Z: float32(long),
	}
}

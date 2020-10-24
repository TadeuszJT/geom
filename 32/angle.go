package geom

import (
	"math"
	"math/rand"
)

type Angle float32

const (
	AnglePi     = Angle(math.Pi)
	Angle2Pi    = Angle(2 * math.Pi)
	Angle45Deg  = Angle(math.Pi / 4)
	Angle90Deg  = Angle(math.Pi / 2)
	Angle180Deg = Angle(math.Pi)
)

func MakeAngle(rad float32) Angle {
	return Angle(math.Mod(float64(rad), 2*math.Pi))
}

func AngleRand() Angle {
	return MakeAngle(rand.Float32() * 2 * float32(math.Pi))
}

func (a Angle) Plus(b Angle) Angle {
	return MakeAngle(float32(a) + float32(b))
}

func (a Angle) Minus(b Angle) Angle {
	return MakeAngle(float32(a) - float32(b))
}

func (a *Angle) PlusEquals(b Angle) {
	*a = a.Plus(b)
}

func (a *Angle) MinusEquals(b Angle) {
	*a = a.Minus(b)
}

func (a *Angle) Clamp(min, max Angle) {
	if *a > max {
		*a = max
	} else if *a < min {
		*a = min
	}
}

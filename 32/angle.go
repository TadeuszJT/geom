package geom

import (
    "math"
    "math/rand"
)

type Angle float32

func MakeAngle(rad float32) Angle {
	return Angle(math.Mod(float64(rad), 2*math.Pi))
}

func AngleRand() Angle {
    return MakeAngle(rand.Float32() * 2 * float32(math.Pi))
}

func (a Angle) Plus(rad float32) Angle {
	return MakeAngle(float32(a) + rad)
}

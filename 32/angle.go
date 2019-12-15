package geom

import "math"

type Angle float32

func MakeAngle(rad float32) Angle {
	mod := math.Mod(float64(rad), 2*math.Pi)
	if math.Signbit(mod) {
		mod += 2 * math.Pi
	}

	return Angle(mod)
}

func (a Angle) Plus(b Angle) Angle {
	return MakeAngle(float32(a + b))
}

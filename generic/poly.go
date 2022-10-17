package geom

type Poly[T Num] []Vec2[T]

func PolyCopy[T Num](poly Poly[T]) Poly[T] {
	return append(Poly[T]{}, poly...)
}

func PolyConvert[A, B Num](a Poly[A]) Poly[B] {
	b := Poly[B]{}
	for i := range a {
		b = append(b, Vec2Convert[A, B](a[i]))
	}
	return b
}

func (poly Poly[T]) Contains(v Vec2[T]) bool {
	if len(poly) < 2 {
		panic("must have at least two verts")
	}

	j := len(poly) - 1
	c := false

	for i := range poly {
		if poly[i] == v { // point is a corner
			return true
		}
		if (poly[i].Y > v.Y) != (poly[j].Y > v.Y) {
			slope := (v.X-poly[i].X)*(poly[j].Y-poly[i].Y) - (poly[j].X-poly[i].X)*(v.Y-poly[i].Y)
			if slope == 0.0 { // point is on a boundary
				return true
			}
			if (slope < 0.0) != (poly[j].Y < poly[i].Y) {
				c = !c
			}
		}
		j = i
	}

	return c
}

/* verts must be in order clockwise */
func PolyArea[T Num](poly Poly[T]) T {
	if len(poly) < 2 {
		panic("must have at least two verts")
	}

	var sum T = 0
	for i := range poly {
		if i == (len(poly) - 1) {
			sum += poly[i].X*poly[0].Y - poly[0].X*poly[i].Y
		} else {
			sum += poly[i].X*poly[i+1].Y - poly[i+1].X*poly[i].Y
		}
	}
	return sum * 0.5
}

func PolyCentroid[T Num](poly Poly[T]) Vec2[T] {
	area := PolyArea(poly) // panic if len(poly) < 2
	if area <= 0.0 {
		panic("area is 0.0")
	}

	var centroid Vec2[T]
	for i := range poly {
		pn := poly[i]
		pn1 := poly[0]
		if i < (len(poly) - 1) {
			pn1 = poly[i+1]
		}

		centroid = centroid.Plus(pn.Plus(pn1).ScaledBy(pn.Cross(pn1)))
	}

	centroid.X /= 6 * area
	centroid.Y /= 6 * area
	return centroid
}

func PolyMomentOfInertia[T Num](poly Poly[T]) T {
	mass := PolyArea(poly) // panic if len(poly) < 2
	if mass <= 0.0 {
		panic("area is 0.0")
	}

	var numerator, denominator T
	for i := range poly {
		pn := poly[i]
		pn1 := poly[0]
		if i < (len(poly) - 1) {
			pn1 = poly[i+1]
		}

		numerator += pn.Cross(pn1) * (pn.Dot(pn) + pn.Dot(pn1) + pn1.Dot(pn1))
		denominator += pn.Cross(pn1)
	}

	return (mass * numerator) / (6 * denominator)
}

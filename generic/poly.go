package geom

/* verts must be in order clockwise */
func PolyArea[T Num](verts []Vec2[T]) T {
	if len(verts) < 2 {
		panic("must have at least two verts")
	}

	var sum T = 0
	for i := range verts {
		if i == (len(verts) - 1) {
			sum += verts[i].X*verts[0].Y - verts[0].X*verts[i].Y
		} else {
			sum += verts[i].X*verts[i+1].Y - verts[i+1].X*verts[i].Y
		}
	}
	return sum * 0.5
}

func PolyCentroid[T Num](verts []Vec2[T]) Vec2[T] {
	area := PolyArea(verts) // panic if len(verts) < 2
	if area <= 0.0 {
		panic("area is 0.0")
	}

	var centroid Vec2[T]
	for i := range verts {
		pn := verts[i]
		pn1 := verts[0]
		if i < (len(verts) - 1) {
			pn1 = verts[i+1]
		}

		centroid = centroid.Plus(pn.Plus(pn1).ScaledBy(pn.Cross(pn1)))
	}

	centroid.X /= 6 * area
	centroid.Y /= 6 * area
	return centroid
}

func PolyMomentOfInertia[T Num](verts []Vec2[T]) T {
	mass := PolyArea(verts) // panic if len(verts) < 2
	if mass <= 0.0 {
		panic("area is 0.0")
	}

	var numerator, denominator T
	for i := range verts {
		pn := verts[i]
		pn1 := verts[0]
		if i < (len(verts) - 1) {
			pn1 = verts[i+1]
		}

		numerator += pn.Cross(pn1) * (pn.Dot(pn) + pn.Dot(pn1) + pn1.Dot(pn1))
		denominator += pn.Cross(pn1)
	}

	return (mass * numerator) / (6 * denominator)
}

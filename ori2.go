package geom

type Ori2 struct {
	X, Y, Theta float64
}

func (o Ori2) Vec2() Vec2 {
	return Vec2{o.X, o.Y}
}

func (a *Ori2) PlusEquals(b Ori2) {
	a.X += b.X
	a.Y += b.Y
	a.Theta += b.Theta
}

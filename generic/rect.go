package geom

type Rect[T Num] struct {
	Min, Max Vec2[T]
}

func RectOrigin[T Num](w, h T) Rect[T] {
	return Rect[T]{
		Vec2[T]{0, 0},
		Vec2[T]{w, h},
	}
}

func RectCentred[T Num](w, h T) Rect[T] {
	wh := w / 2
	hh := h / 2
	return Rect[T]{
		Vec2[T]{-wh, -hh},
		Vec2[T]{wh, hh},
	}
}

func RectCentredAt[T Num](w, h T, pos Vec2[T]) Rect[T] {
	wh := w / 2
	hh := h / 2
	return Rect[T]{
		Vec2[T]{pos.X - wh, pos.Y - hh},
		Vec2[T]{pos.X + wh, pos.Y + hh},
	}
}

func MakeRect[T Num](x, y, w, h T) Rect[T] {
	return Rect[T]{
		Vec2[T]{x, y},
		Vec2[T]{x + w, y + h},
	}
}

func RectConvert[A, B Num](a Rect[A]) Rect[B] {
	return Rect[B]{
		Vec2Convert[A, B](a.Min),
		Vec2Convert[A, B](a.Max),
	}
}

func (r Rect[T]) Width() T {
	return r.Max.X - r.Min.X
}

func (r Rect[T]) Height() T {
	return r.Max.Y - r.Min.Y
}

func (r Rect[T]) Size() Vec2[T] {
	return Vec2[T]{r.Width(), r.Height()}
}

func (r Rect[T]) Contains(v Vec2[T]) bool {
	return v.X >= r.Min.X &&
		v.X <= r.Max.X &&
		v.Y >= r.Min.Y &&
		v.Y <= r.Max.Y
}

func (r Rect[T]) Verts() [4]Vec2[T] {
	return [4]Vec2[T]{
		r.Min,
		{r.Max.X, r.Min.Y},
		r.Max,
		{r.Min.X, r.Max.Y},
	}
}

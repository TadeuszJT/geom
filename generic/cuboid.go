package geom

type Cuboid[T Num] struct {
	Min, Max Vec3[T]
}

func CuboidOrigin[T Num](width, height, depth T) Cuboid[T] {
	return Cuboid[T]{
		Min: Vec3[T]{0, 0, 0},
		Max: Vec3[T]{width, height, depth},
	}
}

func CuboidCentred[T Num](width, height, depth T) Cuboid[T] {
	return Cuboid[T]{
		Min: Vec3[T]{-width / 2, -height / 2, -depth / 2},
		Max: Vec3[T]{width / 2, height / 2, depth / 2},
	}
}

func (c Cuboid[T]) Width() T {
	return c.Max.X - c.Min.X
}

func (c Cuboid[T]) Height() T {
	return c.Max.Y - c.Min.Y
}

func (c Cuboid[T]) Depth() T {
	return c.Max.Z - c.Min.Z
}

func (c Cuboid[T]) Contains(v Vec3[T]) bool {
	return v.X >= c.Min.X &&
		v.X <= c.Max.X &&
		v.Y >= c.Min.Y &&
		v.Y <= c.Max.Y &&
		v.Z >= c.Min.Z &&
		v.Z <= c.Max.Z
}

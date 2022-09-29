package geom

type Cuboid struct {
	Min, Max Vec3
}

func CuboidOrigin(width, height, depth float32) Cuboid {
	return Cuboid{
		Min: Vec3{0, 0, 0},
		Max: Vec3{width, height, depth},
	}
}

func CuboidCentred(width, height, depth float32) Cuboid {
	return Cuboid{
		Min: Vec3{-width / 2, -height / 2, -depth / 2},
		Max: Vec3{width / 2, height / 2, depth / 2},
	}
}

func (c Cuboid) Width() float32 {
	return c.Max.X - c.Min.X
}

func (c Cuboid) Height() float32 {
	return c.Max.Y - c.Min.Y
}

func (c Cuboid) Depth() float32 {
	return c.Max.Z - c.Min.Z
}

func (c Cuboid) Contains(v Vec3) bool {
	return v.X >= c.Min.X &&
		v.X <= c.Max.X &&
		v.Y >= c.Min.Y &&
		v.Y <= c.Max.Y &&
		v.Z >= c.Min.Z &&
		v.Z <= c.Max.Z
}

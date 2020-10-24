package geom

type SliceVec3 []Vec3

func (s *SliceVec3) Len() int {
	return len(*s)
}

func (s *SliceVec3) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *SliceVec3) Delete(i int) {
	end := s.Len() - 1
	if i < end {
		s.Swap(i, end)
	}

	*s = (*s)[:end]
}

func (s *SliceVec3) Append(item interface{}) {
	i, _ := item.(Vec3)
	*s = append(*s, i)
}

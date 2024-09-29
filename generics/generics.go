package generics

type My_List[T comparable] struct {
	value []T
}

func (m *My_List[T]) Find(finding T) *T {
	for _, item := range m.value {
		if finding == item {
			return &item
		}
	}
	return nil
}

package utils

type Vector[T any] struct {
	ID       func(element T) string
	elements map[string]*T
	list     []*T
}

func (v *Vector[T]) GetAll() []*T {
	return v.list
}

func (v *Vector[T]) Get(key string) (*T, bool) {
	if v.elements == nil {
		v.elements = make(map[string]*T)
		return nil, false
	}
	element, ok := v.elements[key]
	if !ok {
		return nil, false
	}
	return element, true
}

func (v *Vector[T]) Add(element *T) bool {
	key := v.ID(*element)
	if _, ok := v.Get(key); ok {
		return false
	}
	if v.list == nil {
		v.list = make([]*T, 0)
	}
	v.list = append(v.list, element)
	v.elements[key] = element
	return true
}

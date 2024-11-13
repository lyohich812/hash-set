package hashset

type HashSet struct {
	data map[interface{}]struct{}
}

// func NewHashSet() *HashSet {
// 	return &HashSet{data: make(map[interface{}]struct{})}
// }

func NewHashSet(elements ...interface{}) *HashSet {
	set := HashSet{data: make(map[interface{}]struct{})}
	for _, el := range elements {
		set.Add(el)
	}
	return &set
}

func (h *HashSet) Add(value interface{}) {
	h.data[value] = struct{}{}
}

func (h *HashSet) Remove(value interface{}) {
	delete(h.data, value)
}

func (h *HashSet) Contains(value interface{}) bool {
	_, exists := h.data[value]
	return exists
}

func (first *HashSet) Union(second *HashSet) *HashSet {
	union := NewHashSet(*first.Iterator()...)
	for key := range second.data {
		union.Add(key)
	}
	return union
}

func (first *HashSet) Intersection(second *HashSet) *HashSet {
	intersection := NewHashSet()
	for key := range second.data {
		if first.Contains(key) {
			intersection.Add(key)
		}
	}
	return intersection
}

func (first *HashSet) Difference(second *HashSet) *HashSet {
	difference := NewHashSet(*first.Iterator()...)
	for key := range second.data {
		if difference.Contains(key) {
			difference.Remove(key)
		}
	}
	return difference
}

func (h *HashSet) Size() int {
	return len(h.data)
}

func (h *HashSet) Iterator() *[]interface{} {
	iterator := make([]interface{}, 0, h.Size())
	for key := range h.data {
		iterator = append(iterator, key)
	}
	return &iterator
}

// NewHashSet
// Add
// Remove
// Contains
// Union
// Intersection
// Difference
// Size
// Iterator (return slice)

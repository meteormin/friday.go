package domain

type CashHash map[string]*Set[uint]

var cacheHash = make(map[string]*Set[uint])

func GetCache() CashHash {
	return cacheHash
}

func (c CashHash) Put(key string, value []uint) {
	c[key] = NewSet(value)
}

func (c CashHash) Add(key string, value uint) {
	c[key].Add(value)
}

func (c CashHash) Get(key string) []uint {
	return c[key].ToSlice()
}

func (c CashHash) Contains(key string) bool {
	_, exists := c[key]
	return exists
}

func (c CashHash) Size() int {
	return len(c)
}

func (c CashHash) Keys() []string {
	keys := make([]string, 0, len(c))
	for key := range c {
		keys = append(keys, key)
	}
	return keys
}

func (c CashHash) Clear() {
	cacheHash = make(map[string]*Set[uint])
}

func (c CashHash) Remove(key string, value uint) {
	c[key].Remove(value)
}

func (c CashHash) Delete(key string) {
	delete(c, key)
}

type Set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable](values []T) *Set[T] {
	s := &Set[T]{data: make(map[T]struct{})}
	for _, value := range values {
		s.data[value] = struct{}{}
	}
	return s
}

func (s *Set[T]) Add(value T) {
	s.data[value] = struct{}{}
}

func (s *Set[T]) Remove(value T) {
	delete(s.data, value)
}

func (s *Set[T]) Contains(value T) bool {
	_, exists := s.data[value]
	return exists
}

func (s *Set[T]) Size() int {
	return len(s.data)
}

func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.data))
	for key := range s.data {
		slice = append(slice, key)
	}
	return slice
}

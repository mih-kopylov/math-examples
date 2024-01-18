package main

// Distribution is used to make sure the values are distributed randomly
type Distribution[T comparable] struct {
	store map[T]int
}

func NewDistribution[T comparable]() *Distribution[T] {
	return &Distribution[T]{
		store: make(map[T]int),
	}
}

func NewDistributionWithKnownKeys[T comparable](knownKeys []T) *Distribution[T] {
	distribution := NewDistribution[T]()
	for _, key := range knownKeys {
		distribution.Add(key)
	}
	return distribution
}

func (d *Distribution[T]) Add(value T) {
	d.store[value] = d.store[value] + 1
}

func (d *Distribution[T]) IsTooFrequent(value T) bool {
	if len(d.store) == 0 {
		return false
	}

	valueCount, found := d.store[value]
	valueCount++

	if found && len(d.store) == 1 && valueCount > 1 {
		return true
	}

	for _, count := range d.store {
		if valueCount-count > 1 {
			return true
		}
	}

	return false
}

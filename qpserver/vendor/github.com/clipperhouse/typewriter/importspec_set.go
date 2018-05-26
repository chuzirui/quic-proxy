// Generated by: main
// TypeWriter: set
// Directive: +gen on ImportSpec

package typewriter

// Set is a modification of https://github.com/deckarep/golang-set
// The MIT License (MIT)
// Copyright (c) 2013 Ralph Caraveo (deckarep@gmail.com)

// The primary type that represents a set
type ImportSpecSet map[ImportSpec]struct{}

// Creates and returns a reference to an empty set.
func NewImportSpecSet(a ...ImportSpec) ImportSpecSet {
	s := make(ImportSpecSet)
	for _, i := range a {
		s.Add(i)
	}
	return s
}

// ToSlice returns the elements of the current set as a slice
func (set ImportSpecSet) ToSlice() []ImportSpec {
	var s []ImportSpec
	for v := range set {
		s = append(s, v)
	}
	return s
}

// Adds an item to the current set if it doesn't already exist in the set.
func (set ImportSpecSet) Add(i ImportSpec) bool {
	_, found := set[i]
	set[i] = struct{}{}
	return !found //False if it existed already
}

// Determines if a given item is already in the set.
func (set ImportSpecSet) Contains(i ImportSpec) bool {
	_, found := set[i]
	return found
}

// Determines if the given items are all in the set
func (set ImportSpecSet) ContainsAll(i ...ImportSpec) bool {
	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

// Determines if every item in the other set is in this set.
func (set ImportSpecSet) IsSubset(other ImportSpecSet) bool {
	for elem := range set {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// Determines if every item of this set is in the other set.
func (set ImportSpecSet) IsSuperset(other ImportSpecSet) bool {
	return other.IsSubset(set)
}

// Returns a new set with all items in both sets.
func (set ImportSpecSet) Union(other ImportSpecSet) ImportSpecSet {
	unionedSet := NewImportSpecSet()

	for elem := range set {
		unionedSet.Add(elem)
	}
	for elem := range other {
		unionedSet.Add(elem)
	}
	return unionedSet
}

// Returns a new set with items that exist only in both sets.
func (set ImportSpecSet) Intersect(other ImportSpecSet) ImportSpecSet {
	intersection := NewImportSpecSet()
	// loop over smaller set
	if set.Cardinality() < other.Cardinality() {
		for elem := range set {
			if other.Contains(elem) {
				intersection.Add(elem)
			}
		}
	} else {
		for elem := range other {
			if set.Contains(elem) {
				intersection.Add(elem)
			}
		}
	}
	return intersection
}

// Returns a new set with items in the current set but not in the other set
func (set ImportSpecSet) Difference(other ImportSpecSet) ImportSpecSet {
	differencedSet := NewImportSpecSet()
	for elem := range set {
		if !other.Contains(elem) {
			differencedSet.Add(elem)
		}
	}
	return differencedSet
}

// Returns a new set with items in the current set or the other set but not in both.
func (set ImportSpecSet) SymmetricDifference(other ImportSpecSet) ImportSpecSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clears the entire set to be the empty set.
func (set *ImportSpecSet) Clear() {
	*set = make(ImportSpecSet)
}

// Allows the removal of a single item in the set.
func (set ImportSpecSet) Remove(i ImportSpec) {
	delete(set, i)
}

// Cardinality returns how many items are currently in the set.
func (set ImportSpecSet) Cardinality() int {
	return len(set)
}

// Iter() returns a channel of type ImportSpec that you can range over.
func (set ImportSpecSet) Iter() <-chan ImportSpec {
	ch := make(chan ImportSpec)
	go func() {
		for elem := range set {
			ch <- elem
		}
		close(ch)
	}()

	return ch
}

// Equal determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set ImportSpecSet) Equal(other ImportSpecSet) bool {
	if set.Cardinality() != other.Cardinality() {
		return false
	}
	for elem := range set {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// Returns a clone of the set.
// Does NOT clone the underlying elements.
func (set ImportSpecSet) Clone() ImportSpecSet {
	clonedSet := NewImportSpecSet()
	for elem := range set {
		clonedSet.Add(elem)
	}
	return clonedSet
}

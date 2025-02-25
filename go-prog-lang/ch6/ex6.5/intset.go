package intset

import (
	"bytes"
	"fmt"
)

const (
	ARCH = 32 << (^uint(0xFF) >> 63)
)

// An IntSet is set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/ARCH, uint(x%ARCH)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/ARCH, uint(x%ARCH)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWidth sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < ARCH; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", ARCH*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// ex6.1
func (s *IntSet) Len() int {
	var len int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < ARCH; j++ {
			if word&(1<<uint(j)) != 0 {
				len++
			}
		}
	}

	return len
}

// ex6.1
func (s *IntSet) Remove(x int) {
	word, bit := x/ARCH, uint(x%64)
	s.words[word] &= ^(1 << bit)
}

// ex6.1
func (s *IntSet) Clear() {
	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < ARCH; j++ {
			if word&(1<<uint(j)) != 0 {
				s.words[i] &= ^(1 << uint(j))
			}
		}
	}
}

// ex6.1
func (s *IntSet) Copy() *IntSet {
	t := IntSet{}
	t.words = append([]uint(nil), s.words...)
	return &t
}

// ex6.2
func (s *IntSet) AddAll(list ...int) {
	for _, x := range list {
		s.Add(x)
	}
}

// ex6.3
func (s *IntSet) IntersectWith(t *IntSet) {
	for i := range s.words {
		s.words[i] &= t.words[i]
	}
}

// ex6.3
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := range s.words {
		s.words[i] &^= t.words[i]
	}

}

// ex6.3
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i := range s.words {
		s.words[i] ^= t.words[i]
	}
}

// ex6.4
func (s *IntSet) Elems() []int {
	elems := make([]int, 0, s.Len())
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < ARCH; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, ARCH*i+j)
			}
		}
	}
	return elems
}

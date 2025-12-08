package main

import (
	"fmt"
	"math"
)

/*
	MCG — Multiplicative Coverage Generator

	This algorithm generates prime numbers not by *testing divisibility*,
	but by *continuing the global multiplicative coverage state* of the
	number space. A number is prime if (and only if) it has not yet been
	covered by the multiplicative waves of all previously known primes.

	The MCG grows unbounded — there is no upper limit and no need to
	restart computation. The generator builds and processes the number
	space segment by segment, each segment extending the global
	multiplicative coverage to higher numbers.
*/

type MCG struct {
	basePrimes []int // all primes discovered so far — used as coverage multipliers
	segment    []uint8
	low        int // lower bound of the current segment
	high       int // upper bound of the current segment
	segSize    int // segment size (performance tuning parameter)
}

// NewMCG creates a new, unbounded Multiplicative Coverage Generator.
func NewMCG() *MCG {
	m := &MCG{
		basePrimes: []int{},
		segSize:    1_000_000, // the segment length can be increased for performance
		low:        2,
	}
	m.initFirstSegment()
	return m
}

// initFirstSegment populates the initial segment starting from 2
// and establishes the first coverage from the smallest primes.
func (m *MCG) initFirstSegment() {
	m.high = m.low + m.segSize - 1
	m.segment = make([]uint8, m.segSize)

	limit := int(math.Sqrt(float64(m.high)))
	for n := 2; n <= limit; n++ {
		if m.isPrimeLocal(n) {
			// n is a prime → add to coverage multipliers
			m.basePrimes = append(m.basePrimes, n)

			// mark multiplicative coverage for all multiples of n
			for x := n * n; x <= m.high; x += n {
				m.segment[x-m.low]++
			}
		}
	}
}

// isPrimeLocal checks primality only within the current segment
// by verifying the multiplicative coverage state.
func (m *MCG) isPrimeLocal(n int) bool {
	if n < m.low || n > m.high {
		return false
	}
	return m.segment[n-m.low] == 0
}

// nextSegment advances the MCG to the next growth region of the number space,
// continuing the multiplicative coverage using all previously discovered primes.
func (m *MCG) nextSegment() {
	m.low = m.high + 1
	m.high = m.low + m.segSize - 1
	m.segment = make([]uint8, m.segSize)

	// project multiplicative coverage of all known primes into the new segment
	for _, p := range m.basePrimes {
		// first multiple of p within the new segment
		first := (m.low + p - 1) / p * p
		if first < p*p {
			first = p * p
		}
		if first > m.high {
			continue
		}
		for x := first; x <= m.high; x += p {
			m.segment[x-m.low]++
		}
	}
}

// Next returns the next prime number — the MCG runs forever.
// No upper bound is required. When a segment is exhausted, a new one is constructed.
func (m *MCG) Next() int {
	for {
		for i := 0; i < len(m.segment); i++ {
			if m.segment[i] == 0 {
				p := m.low + i
				m.segment[i] = 255 // mark consumed (prevents returning the same prime twice)
				m.basePrimes = append(m.basePrimes, p)
				return p
			}
		}
		// segment exhausted → continue coverage
		m.nextSegment()
	}
}

// example usage
func main() {
	g := NewMCG()

	// print the first 10000 primes
	for i := 0; i < 10000; i++ {
		fmt.Println(g.Next())
	}
}

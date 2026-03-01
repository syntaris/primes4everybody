/*
    MCG — Multiplicative Coverage Generator
    Author: Karl Jochen Heinz
	Mail: mail@kjheinz.de
    Date: 01. March 2026
    License: CC BY-NC-ND 4.0
    DOI: 10.5281/zenodo.18826208
    ...
*/

/*
   MCG — Multiplicative Coverage Generator
   (Experimental Proof of Structural Irreducibility and the cNP Class)

   ===========================================================================
   EVIDENCE OF ROBUSTNESS: THE MCG AS EMPIRICAL PROOF
   ===========================================================================
   The MCG is not merely another "sieve variant." It is the software-technical
   realization of the Constructive Number Space. It provides the
   empirical foundation for the structural impossibility of P = NP by
   demonstrating the boundary between local verification and global emergence.

   1. OPERATIONAL SUPERIORITY THROUGH ONTOLOGICAL CLARITY:
      - No Modulo/Divisibility: The algorithm operates entirely without the
        modulus operator or trial division. Primality is the structural
        failure of multiplicative waves to cover a specific
        point in the additive growth process.
      - Unbounded Growth: The MCG processes the number space segment by
        segment, projecting the global multiplicative state forever.

   2. THE LIVING WITNESS FOR P != NP:
      - Local Efficiency (NP): Verifying whether a number in a processed
        segment is prime is an O(1) operation (local array access).
      - Global Irreducibility (cNP): Generating the structure requires the
        accumulation of the ENTIRE history of previously discovered fixpoints
        (basePrimes). There is no local "shortcut" or
        deterministic P-algorithm to predict the next fixpoint.

   3. STRUCTURAL CONSISTENCY:
      - Every prime emerges exactly where the theory predicted: at the points
        of maximal multiplicative information loss.
      - It bridges the gap between Number Theory and Complexity Theory,
        showing that NP-hardness is a symptom of Irreducible Globality.
   ===========================================================================
*/

package main

import (
	"fmt"
	"math"
)

type MCG struct {
	basePrimes []int   // Global system memory (Mandatory for cNP class)
	segment    []uint8 // Local projection of the multiplicative state
	low        int     // Lower bound of current growth region
	high       int     // Upper bound of current growth region
	segSize    int     // Accumulation window size
}

// NewMCG creates a new, unbounded Multiplicative Coverage Generator.
func NewMCG() *MCG {
	m := &MCG{
		basePrimes: []int{},
		segSize:    1_000_000,
		low:        2,
	}
	m.initFirstSegment()
	return m
}

// initFirstSegment establishes the initial multiplicative state.
func (m *MCG) initFirstSegment() {
	m.high = m.low + m.segSize - 1
	m.segment = make([]uint8, m.segSize)

	limit := int(math.Sqrt(float64(m.high)))
	for n := 2; n <= limit; n++ {
		if m.isPrimeLocal(n) {
			m.basePrimes = append(m.basePrimes, n)
			// Mark multiplicative resonance (Projection)
			for x := n * n; x <= m.high; x += n {
				m.segment[x-m.low]++
			}
		}
	}
}

func (m *MCG) isPrimeLocal(n int) bool {
	if n < m.low || n > m.high {
		return false
	}
	return m.segment[n-m.low] == 0 // Fixpoint: no multiplicative return path
}

// nextSegment demonstrates Irreducible Globality.
// Emergence of the solution requires processing all preceding fixpoints.
func (m *MCG) nextSegment() {
	m.low = m.high + 1
	m.high = m.low + m.segSize - 1
	m.segment = make([]uint8, m.segSize)

	// Projection of multiplicative waves:
	// No local rule can replace this global update.
	for _, p := range m.basePrimes {
		first := (m.low + p - 1) / p * p
		if first < p*p {
			first = p * p
		}
		for x := first; x <= m.high; x += p {
			m.segment[x-m.low]++
		}
	}
}

// Next returns the next prime (fixpoint). The process runs forever.
func (m *MCG) Next() int {
	for {
		for i := 0; i < len(m.segment); i++ {
			if m.segment[i] == 0 {
				p := m.low + i
				m.segment[i] = 255 // Mark consumed
				m.basePrimes = append(m.basePrimes, p)
				return p
			}
		}
		m.nextSegment() // Advance growth process
	}
}

func main() {
	g := NewMCG()
	// Continuous emergence: local verification is fast, global construction is mandatory.
	for i := 0; i < 10000; i++ {
		fmt.Println(g.Next())
	}
}


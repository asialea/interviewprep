import "fmt"

// You are given  queries. Each query is of the form two integers described below:
// -  : Insert x in your data structure.
// -  : Delete one occurence of y from your data structure, if present.
// -  : Check if any integer is present whose frequency is exactly . If yes, print 1 else 0.
// The queries are given in the form of a 2-D array  of size q where queries[i][0] contains the operation, and queris
// queries [i][1] contains the data element.

// ideas:
// we need a way to easily fetch if el exists with a given count.
// way to track count of all elements
// 2 hashmaps. seen and occurenncess
// insert - if x in seen, seen[x] += 1  counts[seen[x]] +=1  counts[seen[x-1]]-=1
// delete if x in seen, seen -=1. if see[x] == 0  delete seen[x]

type Counter struct {
	freq      map[int32]int32
	valCounts map[int32]int32
}

func NewCounter() *Counter {
	return &Counter{
		freq:      make(map[int32]int32),
		valCounts: make(map[int32]int32),
	}
}

func addIfNotExists(m map[int32]int32, val int32, inc int32) int32 {
	if _, ok := m[val]; ok {
		m[val] += inc
	} else {
		if inc > 0 {
			m[val] = inc
		} else {
			return 0
		}
	}
	return m[val]
}

func (c *Counter) Insert(val int32) {
	oldCount, _ := c.valCounts[val]
	newCount := addIfNotExists(c.valCounts, val, 1)

	addIfNotExists(c.freq, oldCount, -1)
	addIfNotExists(c.freq, newCount, 1)
}

func (c *Counter) FreqExists(f int32) int32 {
	if val, ok := c.freq[f]; ok && val > 0 {
		fmt.Println(1)
		return 1
	}
	fmt.Println(0)
	return 0
}

func (c *Counter) Delete(val int32) {
	// remove el from valCounts
	oldCount, ok := c.valCounts[val]
	if !ok {
		return
	}
	newCount := addIfNotExists(c.valCounts, val, -1)
	if newCount == 0 {
		delete(c.valCounts, val)
	}

	// update frequencies
	addIfNotExists(c.freq, oldCount, -1)
	if newCount > 0 {
		addIfNotExists(c.freq, newCount, 1)
	}
}

// Complete the freqQuery function below.
func freqQuery(queries [][]int32) []int32 {
	c := NewCounter()
	results := []int32{}
	for _, q := range queries {
		switch {
		case q[0] == int32(1):
			c.Insert(q[1])
		case q[0] == int32(2):
			c.Delete(q[1])
		case q[0] == int32(3):
			results = append(results, c.FreqExists(q[1]))
		default:
		}
	}
	return results
}
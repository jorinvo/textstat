package lib

// Histogram ...
type Histogram map[string]int

// RemoveMax element with biggest value from histogram.
func (h Histogram) RemoveMax() string {
	max := h.maxByValue()
	delete(h, max)
	return max
}

func (h Histogram) maxByValue() string {
	var (
		maxKey string
		maxVal int
	)

	for key, val := range h {
		if val > maxVal {
			maxKey = key
			maxVal = val
		}
	}

	return maxKey
}

package lib

// Histogram ...
type Histogram map[string]int

// RemoveMax element with biggest value from histogram.
func (h Histogram) RemoveMax() string {
	key := h.findMax()
	delete(h, key)
	return key
}

func (h Histogram) findMax() string {
	var res string
	var maxVal int
	for key, val := range h {
		if val > maxVal {
			maxVal = val
			res = key
		}
	}
	return res
}

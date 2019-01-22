package stats

// JaccardIndex returns similarity from two sets
// https://en.wikipedia.org/wiki/Jaccard_index
func JaccardIndex(x, y []float64) float64 {
	if len(x) == 0 && len(y) == 0 {
		return 1.0
	}

	union := []float64{}
	for _, a := range x {
		for _, b := range y {
			if a == b {
				union = append(union, a)
			}
		}

	}

	return float64(len(union)) / float64((len(x) + len(y) - len(union)))
}

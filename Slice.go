package main

func Slice(a []string, nbrs ...int) []string {
	var start, end int

	switch len(nbrs) {
	case 1:
		start = normalizeIndex(nbrs[0], len(a))
		end = len(a)
	case 2:
		start = normalizeIndex(nbrs[0], len(a))
		end = normalizeIndex(nbrs[1], len(a))
	default:
		panic("Slice function requires 1 or 2 integer arguments")
	}

	if start < 0 || start >= len(a) {
		return nil
	}
	if end < 0 || end > len(a) {
		end = len(a)
	}

	return a[start:end]
}

func normalizeIndex(idx, length int) int {
	if idx < 0 {
		idx += length
		if idx < 0 {
			idx = 0
		}
	}
	if idx > length {
		idx = length
	}
	return idx
}

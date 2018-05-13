package factor

// Factor type
// FwdMap maps unique string values to an int index.
// BwdMap maps int indices to unique string values.
// Column is the encoded vector.
type Factor struct {
	FwdMap map[string]int
	BwdMap []string
	Column []int
}

// Number of unique values.
func (f Factor) ULen() int {
	return len(f.BwdMap)
}

// Number of values.
func (f Factor) Len() int {
	return len(f.Column)
}

// Get the unencoded value at index i.
func (f Factor) Value(i int) string {
	return f.BwdMap[f.Column[i]]
}

func NewFactor() Factor {
	var f Factor
	f.FwdMap = make(map[string]int)
	return f
}

func ToFactor(ss []string) Factor {
	var f Factor
	f = NewFactor()
	f.Column = make([]int, 0, len(ss))
	for _, s := range ss {
		f.Append(s)
	}
	return f
}

func (f Factor) Append(s string) {
	idx, ok := f.FwdMap[s]
	if !ok {
		idx = len(f.BwdMap)
		f.FwdMap[s] = idx
		f.BwdMap = append(f.BwdMap, s)
	}
	f.Column = append(f.Column, idx)
}

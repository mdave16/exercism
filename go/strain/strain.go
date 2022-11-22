package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) (out Ints) {
	for _, v := range i {
		if filter(v) {
			out = append(out, v)
		}
	}
	return
}

func (i Ints) Discard(filter func(int) bool) Ints {
	return i.Keep(func(n int) bool { return !filter(n) })
}

func (l Lists) Keep(filter func([]int) bool) (out Lists) {
	for _, v := range l {
		if filter(v) {
			out = append(out, v)
		}
	}
	return
}

func (s Strings) Keep(filter func(string) bool) (out Strings) {
	for _, v := range s {
		if filter(v) {
			out = append(out, v)
		}
	}
	return
}

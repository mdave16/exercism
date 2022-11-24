package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(acc, item int) int, initial int) int {
	if len(s) == 0 {
		return initial
	}

	return s[1:].Foldl(fn, fn(initial, s[0]))
}

func (s IntList) Foldr(fn func(item, acc int) int, initial int) int {
	if len(s) == 0 {
		return initial
	}

	return s[:len(s)-1].Foldr(fn, fn(s[len(s)-1], initial))
}

func (s IntList) Filter(fn func(int) bool) IntList {
	out := IntList{}
	for _, v := range s {
		if fn(v) {
			out = append(out, v)
		}
	}
	return out
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	out := IntList{}
	for _, v := range s {
		out = append(out, fn(v))
	}
	return out
}

func (s IntList) Reverse() IntList {
	out := IntList{}
	for i := len(s) - 1; i >= 0; i-- {
		out = append(out, s[i])
	}
	return out
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	if len(lists) == 0 {
		return s
	}
	return s.Append(lists[0]).Concat(lists[1:])
}

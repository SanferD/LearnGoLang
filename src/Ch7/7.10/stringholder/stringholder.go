package stringholder

type StringHolder struct {
	letters []rune
}

func MakeStringHolder(s string) StringHolder {
	return StringHolder{[]rune(s)}
}

func (s StringHolder) Len() int {
	return len(s.letters)
}

func (s StringHolder) Less(i, j int) bool {
	return s.letters[i] < s.letters[j]
}

func (s StringHolder) Swap(i, j int) {
	s.letters[i], s.letters[j] = s.letters[j], s.letters[i]
}

func (s StringHolder) String() string {
	return string(s.letters)
}
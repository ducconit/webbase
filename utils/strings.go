package utils

type String struct {
	value string
}

func StringOf(value string) *String {
	return &String{value}
}

func (s *String) Join(strs ...string) *String {
	for _, str := range strs {
		s.value += str
	}
	return s
}

func (s *String) Value() string {
	return s.value
}

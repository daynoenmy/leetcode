package main

type stack struct {
	data []int32
	p    int
}

func (s *stack) pop() (int32, bool) {
	if s.p == 0 {
		println("stack is empty")
		return ' ', false
	}
	s.p -= 1
	return s.data[s.p], true
}
func (s *stack) peek() (int32, bool) {
	if s.p == 0 {
		println("stack is empty")
		return ' ', false
	}
	return s.data[s.p-1], true
}
func (s *stack) push(val int32) {
	s.data[s.p] = val
	s.p += 1
}
func (s *stack) init(len int) {
	s.p = 0
	s.data = make([]int32, len)
}
func isVaild(s string) bool {
	st := new(stack)
	st.init(len(s))
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			st.push(v)
		} else if v == ')' {
			k, t := st.peek()
			if !t {
				return false
			}
			if k == '(' {
				st.pop()
			} else {
				return false
			}
		} else if v == ']' {
			k, t := st.peek()
			if !t {
				return false
			}
			if k == '[' {
				st.pop()
			} else {
				return false
			}
		} else if v == '}' {
			k, t := st.peek()
			if !t {
				return false
			}
			if k == '{' {
				st.pop()
			} else {
				return false
			}
		}
	}
	return st.p == 0
}
func main() {
	println(isVaild("(()[{}])"))
}

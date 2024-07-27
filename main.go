package main

import (
	"errors"
	"strconv"
)

func main() {
	d := data{
		code: []string{
			"push",
			"1",
			"push",
			"2",
			"add",
			"print",
			"halt",
		},
		stack:     []int{},
		variables: map[string]int{},
		pc:        0,
	}
	d.run()
}

type data struct {
	code      []string
	stack     []int
	variables map[string]int
	pc        int
}

func (s *data) jump(x int) {
	s.pc = x
}
func (s *data) jump_if() error {
	if len(s.stack) == 0 {
		return errors.New("stack is empty")
	}
	if s.stack[len(s.stack)-1] != 0 {
		s.pc = s.stack[len(s.stack)-1]
	}
	return nil
}

func (s *data) push(val int) {
	s.stack = append(s.stack, val)
}

func (s *data) pop() error {
	if len(s.stack) == 0 {
		return errors.New("stack does not have value")
	}
	s.stack = s.stack[:len(s.stack)-1]
	return nil
}

func (s *data) add() error {
	if len(s.stack) < 2 {
		return errors.New("stack is too short")
	}
	val1 := s.stack[len(s.stack)-1]
	val2 := s.stack[len(s.stack)-2]
	s.stack = s.stack[:len(s.stack)-2]
	s.stack = append(s.stack, val1+val2)
	return nil
}

func (s *data) sub() error {
	if len(s.stack) < 2 {
		return errors.New("stack is too short")
	}
	val1 := s.stack[len(s.stack)-1]
	val2 := s.stack[len(s.stack)-2]
	s.stack = s.stack[:len(s.stack)-2]
	s.stack = append(s.stack, val1-val2)
	return nil
}

func (s *data) mul() error {
	if len(s.stack) < 2 {
		return errors.New("stack is too short")
	}
	val1 := s.stack[len(s.stack)-1]
	val2 := s.stack[len(s.stack)-2]
	s.stack = s.stack[:len(s.stack)-2]
	s.stack = append(s.stack, val1*val2)
	return nil
}

func (s *data) set_x(x string) error {
	if len(s.stack) == 0 {
		return errors.New("stack is empty")
	}
	s.variables[x] = s.stack[len(s.stack)-1]
	return nil
}

func (s *data) get_x(x string) error {
	if _, ok := s.variables[x]; !ok {
		return errors.New("variable not found")
	}
	s.stack = append(s.stack, s.variables[x])
	return nil
}

func (s *data) print() error {
	if len(s.stack) == 0 {
		return errors.New("stack is empty")
	}
	println(s.stack[len(s.stack)-1])
	return nil
}

func (s *data) halt() {
	s.pc = -1
}

func (s *data) run() {
	for s.pc != -1 {
		switch s.code[s.pc] {
		case "push":
			i, err := strconv.Atoi(s.code[s.pc+1])
			if err != nil {
				panic(err)
			}
			s.push(i)
			s.pc += 2
		case "pop":
			err := s.pop()
			if err != nil {
				panic(err)
			}
			s.pc++
		case "add":
			if err := s.add(); err != nil {
				panic(err)
			}
			s.pc++
		case "sub":
			if err := s.sub(); err != nil {
				panic(err)
			}
			s.pc++
		case "mul":
			if err := s.mul(); err != nil {
				panic(err)
			}
			s.pc++
		case "set":
			if err := s.set_x(s.code[s.pc+1]); err != nil {
				panic(err)
			}
			s.pc += 2
		case "get":
			if err := s.get_x(s.code[s.pc+1]); err != nil {
				panic(err)
			}
			s.pc += 2
		case "print":
			if err := s.print(); err != nil {
				panic(err)
			}
			s.pc++
		case "jump":
			s.jump(s.stack[s.pc+1])
		case "jump_if":
			if err := s.jump_if(); err != nil {
				panic(err)
			}
			s.pc++
		case "halt":
			s.halt()
		}
	}
}

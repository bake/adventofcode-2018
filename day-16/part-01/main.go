package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal()
	}
	parts := bytes.Split(b, []byte("\n\n\n"))
	samples := parse(parts[0])

	var num int
	for _, s := range samples {
		if len(possibilities(s)) >= 3 {
			num++
		}
	}
	fmt.Println(num)
}

func possibilities(s sample) []string {
	var res []string
	for opcode, fn := range opcodes {
		register = make([]uint8, len(s.Before))
		copy(register, s.Before)
		fn(s.Instruction[1], s.Instruction[2], s.Instruction[3])
		if equal(s.After, register) {
			res = append(res, opcode)
		}
	}
	return res
}

type sample struct {
	Before      []uint8
	Instruction []uint8
	After       []uint8
}

func equal(a, b []uint8) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func parse(b []byte) []sample {
	var samples []sample
	bs := bytes.Split(b, []byte("\n\n"))
	for _, b := range bs {
		rows := bytes.Split(b, []byte("\n"))
		var s sample
		s.Before = make([]uint8, 4)
		s.Instruction = make([]uint8, 4)
		s.After = make([]uint8, 4)
		fmt.Sscanf(string(rows[0]), "Before: [%d, %d, %d, %d]", &s.Before[0], &s.Before[1], &s.Before[2], &s.Before[3])
		fmt.Sscanf(string(rows[1]), "%d %d %d %d", &s.Instruction[0], &s.Instruction[1], &s.Instruction[2], &s.Instruction[3])
		fmt.Sscanf(string(rows[2]), "After:  [%d, %d, %d, %d]", &s.After[0], &s.After[1], &s.After[2], &s.After[3])
		samples = append(samples, s)
	}
	return samples
}

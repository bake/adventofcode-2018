package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/bakerolls/adventofcode-2018/day-16/opcodes"
)

// Warning, reader! This might be the worst code so far.
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
	for opcode, fn := range opcodes.Opcodes {
		for i := range opcodes.Register {
			opcodes.Register[i] = s.Before[i]
		}
		fn(s.Instruction[1], s.Instruction[2], s.Instruction[3])
		if opcodes.RegisterEqual(s.After) {
			res = append(res, opcode)
		}
	}
	return res
}

type sample struct {
	Before      [4]uint8
	Instruction [4]uint8
	After       [4]uint8
}

func parse(b []byte) []sample {
	var samples []sample
	bs := bytes.Split(b, []byte("\n\n"))
	for _, b := range bs {
		rows := bytes.Split(b, []byte("\n"))
		var s sample
		fmt.Sscanf(string(rows[0]), "Before: [%d, %d, %d, %d]", &s.Before[0], &s.Before[1], &s.Before[2], &s.Before[3])
		fmt.Sscanf(string(rows[1]), "%d %d %d %d", &s.Instruction[0], &s.Instruction[1], &s.Instruction[2], &s.Instruction[3])
		fmt.Sscanf(string(rows[2]), "After:  [%d, %d, %d, %d]", &s.After[0], &s.After[1], &s.After[2], &s.After[3])
		samples = append(samples, s)
	}
	return samples
}

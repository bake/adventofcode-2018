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
	parts := bytes.Split(b, []byte("\n\n\n\n"))
	samples := parse(parts[0])

	// Create a map for every index containing the names of possible opcodes.
	opmap := map[int]map[string]bool{}
	for _, s := range samples {
		for _, o := range possibilities(s) {
			if opmap[s.Instruction[0]] == nil {
				opmap[s.Instruction[0]] = map[string]bool{}
			}
			opmap[s.Instruction[0]][o] = true
		}
	}

	// Simplyfy until each index has ony one opcode left.
	for {
		done := true
		for _, o := range opmap {
			if len(o) > 1 {
				done = false
				break
			}
		}
		if done {
			break
		}

		for i, names := range opmap {
			if len(names) == 1 {
				var name string
				for o := range names {
					name = o
					break
				}

				for j := range opmap {
					if i == j {
						continue
					}
					delete(opmap[j], name)
				}
			}
		}
	}

	// Execute the program.
	instructions := bytes.Split(parts[1], []byte("\n"))
	opcodes.Register = [4]int{}
	for _, ins := range instructions {
		var vals [4]int
		_, err := fmt.Sscanf(string(ins), "%d %d %d %d", &vals[0], &vals[1], &vals[2], &vals[3])
		if err != nil {
			continue
		}

		var name string
		for o := range opmap[vals[0]] {
			name = o
			break
		}
		opcodes.Opcodes[name](vals[1], vals[2], vals[3])
	}
	fmt.Println(opcodes.Register[0])
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
	Before      [4]int
	Instruction [4]int
	After       [4]int
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

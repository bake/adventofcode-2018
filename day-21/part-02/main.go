package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/bakerolls/adventofcode-2018/day-19/opcodes"
)

func main() {
	p, err := parse("../input.txt", 6)
	if err != nil {
		log.Fatal(err)
	}
	seen := map[int]bool{}
	var prev int
	// This is very slow. It took about two minutes. Transpiling the elfcode to
	// Go might help.
	for p.step() {
		if p.register[p.ipr] != 28 {
			continue
		}
		if _, ok := seen[p.register[2]]; ok {
			break
		}
		prev = p.register[2]
		seen[prev] = true
	}
	fmt.Println(prev)
}

type program struct {
	ipr          int
	register     []int
	instructions []instruction
}

func (p *program) step() bool {
	ins := p.instructions[p.register[p.ipr]]
	opcodes.Opcodes[ins.name](p.register, ins.a, ins.b, ins.c)
	p.register[p.ipr]++
	return p.register[p.ipr] < len(p.instructions)
}

type instruction struct {
	name    string
	a, b, c int
}

func (i instruction) String() string {
	return fmt.Sprintf("%s %d %d %d", i.name, i.a, i.b, i.c)
}

func parse(file string, register int) (*program, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	bs := bytes.Split(b, []byte("\n"))

	ipr, err := strconv.Atoi(string(bs[0][4:]))
	if err != nil {
		return nil, err
	}

	var ins []instruction
	for _, b := range bs[1:] {
		var i instruction
		_, err := fmt.Sscanf(string(b), "%s %d %d %d", &i.name, &i.a, &i.b, &i.c)
		if err != nil {
			return nil, err
		}
		ins = append(ins, i)
	}

	return &program{ipr, make([]int, register), ins}, nil
}

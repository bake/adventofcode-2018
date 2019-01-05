package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/bakerolls/adventofcode-2018/day-19/opcodes"
)

type instruction struct {
	name    string
	a, b, c int
}

func (i instruction) String() string {
	return fmt.Sprintf("%s %d %d %d", i.name, i.a, i.b, i.c)
}

func main() {
	ipr, prog, err := parse("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	reg := make([]int, 6)
	for i := 0; reg[ipr] <= len(prog); i++ {
		ins := prog[reg[ipr]]
		// fmt.Printf("[%4d] ip=%d %v %s ", i, reg[ipr], reg, ins)
		opcodes.Opcodes[ins.name](reg, ins.a, ins.b, ins.c)
		// fmt.Printf("%v\n", reg)
		reg[ipr]++
	}
	fmt.Println(reg[0])
}

func parse(file string) (int, []instruction, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return 0, nil, err
	}
	bs := bytes.Split(b, []byte("\n"))
	var prog []instruction
	for _, b := range bs[1:] {
		var ins instruction
		_, err := fmt.Sscanf(string(b), "%s %d %d %d", &ins.name, &ins.a, &ins.b, &ins.c)
		if err != nil {
			return 0, nil, err
		}
		prog = append(prog, ins)
	}
	ipr, err := strconv.Atoi(string(bs[0][4:]))
	if err != nil {
		return 0, nil, err
	}
	return ipr, prog, nil
}

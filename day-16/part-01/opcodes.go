package main

var register []uint8

// OpcodeFunc is a function taht manipulates the register.
type OpcodeFunc func(a, b, c uint8)

var opcodes = map[string]OpcodeFunc{
	"addr": func(a, b, c uint8) { register[c] = register[a] + register[b] },
	"addi": func(a, b, c uint8) { register[c] = register[a] + b },
	"mulr": func(a, b, c uint8) { register[c] = register[a] * register[b] },
	"muli": func(a, b, c uint8) { register[c] = register[a] * b },
	"banr": func(a, b, c uint8) { register[c] = register[a] & register[b] },
	"bani": func(a, b, c uint8) { register[c] = register[a] & b },
	"borr": func(a, b, c uint8) { register[c] = register[a] | register[b] },
	"bori": func(a, b, c uint8) { register[c] = register[a] | b },
	"setr": func(a, b, c uint8) { register[c] = register[a] },
	"seti": func(a, b, c uint8) { register[c] = a },
	"gtir": func(a, b, c uint8) {
		if a > register[b] {
			register[c] = 1
			return
		}
		register[c] = 0
	},
	"gtri": func(a, b, c uint8) {
		if register[a] > b {
			register[c] = 1
			return
		}
		register[c] = 0
	},
	"gtrr": func(a, b, c uint8) {
		if register[a] > register[b] {
			register[c] = 1
			return
		}
		register[c] = 0
	},
	"eqir": func(a, b, c uint8) {
		if a == register[b] {
			register[c] = 1
			return
		}
		register[c] = 0
	},
	"eqri": func(a, b, c uint8) {
		if register[a] == b {
			register[c] = 1
			return
		}
		register[c] = 0
	},
	"eqrr": func(a, b, c uint8) {
		if register[a] == register[b] {
			register[c] = 1
			return
		}
		register[c] = 0
	},
}

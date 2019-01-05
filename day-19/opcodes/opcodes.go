package opcodes

// OpcodeFunc is a function that manipulates the Register.
type OpcodeFunc func(register []int, a, b, c int)

// Opcodes manipulare the register.
var Opcodes = map[string]OpcodeFunc{
	"addr": func(register []int, a, b, c int) { register[c] = register[a] + register[b] },
	"addi": func(register []int, a, b, c int) { register[c] = register[a] + b },
	"mulr": func(register []int, a, b, c int) { register[c] = register[a] * register[b] },
	"muli": func(register []int, a, b, c int) { register[c] = register[a] * b },
	"banr": func(register []int, a, b, c int) { register[c] = register[a] & register[b] },
	"bani": func(register []int, a, b, c int) { register[c] = register[a] & b },
	"borr": func(register []int, a, b, c int) { register[c] = register[a] | register[b] },
	"bori": func(register []int, a, b, c int) { register[c] = register[a] | b },
	"setr": func(register []int, a, b, c int) { register[c] = register[a] },
	"seti": func(register []int, a, b, c int) { register[c] = a },
	"gtir": func(register []int, a, b, c int) {
		if a > register[b] {
			register[c] = 1
			return
		}
		register[c] = 0
	},
	"gtri": func(register []int, a, b, c int) {
		if register[a] > b {
			register[c] = 1
			return
		}
		register[c] = 0
	},
	"gtrr": func(register []int, a, b, c int) {
		if register[a] > register[b] {
			register[c] = 1
			return
		}
		register[c] = 0
	},
	"eqir": func(register []int, a, b, c int) {
		if a == register[b] {
			register[c] = 1
			return
		}
		register[c] = 0
	},
	"eqri": func(register []int, a, b, c int) {
		if register[a] == b {
			register[c] = 1
			return
		}
		register[c] = 0
	},
	"eqrr": func(register []int, a, b, c int) {
		if register[a] == register[b] {
			register[c] = 1
			return
		}
		register[c] = 0
	},
}

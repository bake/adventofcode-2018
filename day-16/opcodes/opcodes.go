package opcodes

// Register holds ints. And it's a exported and global variable. This isn't
// good.
var Register [4]int

// OpcodeFunc is a function taht manipulates the Register.
type OpcodeFunc func(a, b, c int)

// Opcodes manipulare the register.
var Opcodes = map[string]OpcodeFunc{
	"addr": func(a, b, c int) { Register[c] = Register[a] + Register[b] },
	"addi": func(a, b, c int) { Register[c] = Register[a] + b },
	"mulr": func(a, b, c int) { Register[c] = Register[a] * Register[b] },
	"muli": func(a, b, c int) { Register[c] = Register[a] * b },
	"banr": func(a, b, c int) { Register[c] = Register[a] & Register[b] },
	"bani": func(a, b, c int) { Register[c] = Register[a] & b },
	"borr": func(a, b, c int) { Register[c] = Register[a] | Register[b] },
	"bori": func(a, b, c int) { Register[c] = Register[a] | b },
	"setr": func(a, b, c int) { Register[c] = Register[a] },
	"seti": func(a, b, c int) { Register[c] = a },
	"gtir": func(a, b, c int) {
		if a > Register[b] {
			Register[c] = 1
			return
		}
		Register[c] = 0
	},
	"gtri": func(a, b, c int) {
		if Register[a] > b {
			Register[c] = 1
			return
		}
		Register[c] = 0
	},
	"gtrr": func(a, b, c int) {
		if Register[a] > Register[b] {
			Register[c] = 1
			return
		}
		Register[c] = 0
	},
	"eqir": func(a, b, c int) {
		if a == Register[b] {
			Register[c] = 1
			return
		}
		Register[c] = 0
	},
	"eqri": func(a, b, c int) {
		if Register[a] == b {
			Register[c] = 1
			return
		}
		Register[c] = 0
	},
	"eqrr": func(a, b, c int) {
		if Register[a] == Register[b] {
			Register[c] = 1
			return
		}
		Register[c] = 0
	},
}

// RegisterEqual compares an array with the register.
func RegisterEqual(r [4]int) bool {
	for i := range Register {
		if Register[i] != r[i] {
			return false
		}
	}
	return true
}

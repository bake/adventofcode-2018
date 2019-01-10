After evaluating the three setup-blocks, `b` and `c` are `1`, `e` is `10551293` and the rest is `0` (see `comments.txt`).

This version is pretty close to the given input. `b` and `c` are used as iterators.

```javascript
let [a, b, c, d, e, p] = [0, 1, 1, 0, 10551293, 0]
c = 1
do {
	b = 1
	do {
		if (c * b == e) {
			a += c
		}
		b++
	} while (b <= e)
	c++
} while (c <= e)
```

Written with `for` loops instead of `do..while` to make it look a little more familiar:

```javascript
const e = 10551293
for (let c = 1; c <= e; c++) {
	for (let b = 1; b <= e; b++) {
		if (c * b == e) {
			a += c
		}
	}
}
```

The goal here is to see that it's just prime factorization of `e` (10551293). The puzzles solution is the sum of all factors in (0, e].

```
1 + 53 + 199081 + 10551293 = 10750428
```

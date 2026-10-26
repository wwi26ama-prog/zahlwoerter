package exercises

import "fmt"

func ExampleBelowThousand() {
	for _, n := range []int{0, 9, 42, 100, 101, 111, 123, 200, 316, 999} {
		fmt.Println(BelowThousand(n))
	}

	// Output:
	// null
	// neun
	// zweiundvierzig
	// einhundert
	// einhunderteins
	// einhundertelf
	// einhundertdreiundzwanzig
	// zweihundert
	// dreihundertsechzehn
	// neunhundertneunundneunzig
}

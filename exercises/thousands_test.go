package exercises

import "fmt"

func ExampleBelowMillion() {
	for _, n := range []int{0, 999, 1000, 1001, 2024, 21000, 101101, 123456, 999999} {
		fmt.Println(BelowMillion(n))
	}

	// Output:
	// null
	// neunhundertneunundneunzig
	// eintausend
	// eintausendeins
	// zweitausendvierundzwanzig
	// einundzwanzigtausend
	// einhunderteintausendeinhunderteins
	// einhundertdreiundzwanzigtausendvierhundertsechsundfünfzig
	// neunhundertneunundneunzigtausendneunhundertneunundneunzig
}

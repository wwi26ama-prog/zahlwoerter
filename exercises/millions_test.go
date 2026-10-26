package exercises

import "fmt"

func ExampleBelowBillion() {
	for _, n := range []int{999999, 1000000, 1000001, 2000000, 2300000, 21000000, 123456789} {
		fmt.Println(BelowBillion(n))
	}

	// Output:
	// neunhundertneunundneunzigtausendneunhundertneunundneunzig
	// eine Million
	// eine Million eins
	// zwei Millionen
	// zwei Millionen dreihunderttausend
	// einundzwanzig Millionen
	// einhundertdreiundzwanzig Millionen vierhundertsechsundfünfzigtausendsiebenhundertneunundachtzig
}

package exercises

import "fmt"

func ExampleCapitalize() {
	fmt.Println(Capitalize("null"))
	fmt.Println(Capitalize("eine Million"))
	fmt.Println(Capitalize("Elf"))
	fmt.Printf("%q\n", Capitalize(""))

	// Output:
	// Null
	// Eine Million
	// Elf
	// ""
}

func ExampleNumberWord() {
	fmt.Println(NumberWord(0))
	fmt.Println(NumberWord(7))
	fmt.Println(NumberWord(123))
	fmt.Println(NumberWord(999))

	// Output:
	// Null
	// Sieben
	// Einhundertdreiundzwanzig
	// Neunhundertneunundneunzig
}

func ExampleNumberWord_sixDigits() {
	fmt.Println(NumberWord(1000))
	fmt.Println(NumberWord(123456))

	// Output:
	// Eintausend
	// Einhundertdreiundzwanzigtausendvierhundertsechsundfünfzig
}

func ExampleNumberWord_nineDigits() {
	fmt.Println(NumberWord(1000000))
	fmt.Println(NumberWord(123456789))

	// Output:
	// Eine Million
	// Einhundertdreiundzwanzig Millionen vierhundertsechsundfünfzigtausendsiebenhundertneunundachtzig
}

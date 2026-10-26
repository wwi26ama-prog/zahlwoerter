package exercises

import "fmt"

func ExampleOnes() {
	for n := range 10 {
		fmt.Println(Ones(n))
	}

	// Output:
	// null
	// eins
	// zwei
	// drei
	// vier
	// fünf
	// sechs
	// sieben
	// acht
	// neun
}

func ExampleCompoundForm() {
	fmt.Println(CompoundForm("eins"))
	fmt.Println(CompoundForm("einhunderteins"))
	fmt.Println(CompoundForm("einhundert"))
	fmt.Println(CompoundForm("zwei"))
	fmt.Println(CompoundForm("elf"))

	// Output:
	// ein
	// einhundertein
	// einhundert
	// zwei
	// elf
}

func ExampleBelowTwenty() {
	for _, n := range []int{0, 1, 9, 10, 11, 12, 13, 15, 16, 17, 18, 19} {
		fmt.Println(BelowTwenty(n))
	}

	// Output:
	// null
	// eins
	// neun
	// zehn
	// elf
	// zwölf
	// dreizehn
	// fünfzehn
	// sechzehn
	// siebzehn
	// achtzehn
	// neunzehn
}

func ExampleBelowHundred() {
	for _, n := range []int{0, 7, 15, 20, 21, 30, 37, 66, 71, 99} {
		fmt.Println(BelowHundred(n))
	}

	// Output:
	// null
	// sieben
	// fünfzehn
	// zwanzig
	// einundzwanzig
	// dreißig
	// siebenunddreißig
	// sechsundsechzig
	// einundsiebzig
	// neunundneunzig
}

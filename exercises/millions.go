package exercises

// BelowBillion erwartet eine Zahl von `0` bis `999 999 999` und liefert
// ihr Zahlwort. "Million" ist ein Nomen und steht als eigenes Wort:
// "eine Million", "zwei Millionen dreihunderttausend".
// Millionen, die auf `01` enden (z.B. `101` Millionen), müssen Sie
// zunächst nicht beachten (siehe README).
func BelowBillion(n int) string {
	// Hinweis:
	// Das Muster aus `BelowMillion` wiederholt sich mit der nächsten
	// Dreiergruppe, also mit `1000000`.
	// - Unterschiede: Für genau eine Million heißt es "eine Million",
	//   sonst "Millionen", und die Wörter sind durch Leerzeichen getrennt.

	if n < 1000000 {
		return BelowMillion(n)
	}
	word := "eine Million"
	if n/1000000 > 1 {
		word = BelowThousand(n/1000000) + " Millionen"
	}
	if n%1000000 > 0 {
		word += " " + BelowMillion(n%1000000)
	}
	return word
}

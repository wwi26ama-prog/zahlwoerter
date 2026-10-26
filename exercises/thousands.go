package exercises

// BelowMillion erwartet eine Zahl von `0` bis `999 999` und liefert ihr
// Zahlwort, z.B. "eintausend", "einundzwanzigtausend" oder
// "einhundertdreiundzwanzigtausendvierhundertsechsundfünfzig".
func BelowMillion(n int) string {
	// Hinweis:
	// Teilen Sie die Zahl in Dreiergruppen auf: `123456` besteht aus
	// `123` Tausendern und dem Rest `456`.
	// - Für jede Gruppe gibt es schon `BelowThousand`.
	// - Der Rest geht genauso wie in `BelowThousand`, nur mit `1000`
	//   und "tausend" statt `100` und "hundert".

	if n < 1000 {
		return BelowThousand(n)
	}
	word := CompoundForm(BelowThousand(n/1000)) + "tausend"
	if n%1000 > 0 {
		word += BelowThousand(n % 1000)
	}
	return word
}

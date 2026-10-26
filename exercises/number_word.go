package exercises

// Capitalize erwartet ein Wort, das mit einem Buchstaben von `a` bis `z`
// beginnt, und liefert es mit großem Anfangsbuchstaben.
// Ein leeres Wort bleibt leer.
//
// TODO
func Capitalize(word string) string {
	// Hinweis:
	// - `word[:1]` ist der erste Buchstabe, `word[1:]` der Rest.
	// - Im Paket `strings` gibt es eine Funktion für Großbuchstaben.
	// - Was passiert mit `word[:1]` bei einem leeren Wort?

	return ""
}

// NumberWord erwartet eine Zahl von `0` bis `999 999 999` und liefert ihr
// Zahlwort mit großem Anfangsbuchstaben, z.B. "Einhundertdreiundzwanzig".
//
// Beginnen Sie mit Zahlen bis `999` (Test `ExampleNumberWord`) und
// erweitern Sie die Funktion, sobald `BelowMillion` und `BelowBillion`
// fertig sind (Tests `ExampleNumberWord_sixDigits` und
// `ExampleNumberWord_nineDigits`).
//
// TODO
func NumberWord(n int) string {
	// Hinweis:
	// - Für Zahlen bis `999` genügen `BelowThousand` und `Capitalize`.
	// - Für größere Zahlen ersetzen Sie später nur `BelowThousand`.

	return ""
}

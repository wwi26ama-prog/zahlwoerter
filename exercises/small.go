package exercises

// Ones erwartet eine Ziffer von `0` bis `9` und liefert ihr Zahlwort,
// also "null", "eins", "zwei", ..., "neun".
//
// TODO
func Ones(n int) string {
	// Hinweis:
	// Legen Sie die zehn Wörter in einer Liste ab.
	// Die Ziffer ist dann genau der Index des passenden Worts.

	return ""
}

// CompoundForm erwartet ein Zahlwort und liefert die Form, die vor einem
// weiteren Wortteil steht: Ein "eins" am Ende wird zu "ein", wie in
// "einundzwanzig", "einhundert" oder "einhunderteintausend".
// Alle anderen Wörter bleiben unverändert.
//
// TODO
func CompoundForm(word string) string {
	// Hinweis:
	// Das Paket `strings` enthält Funktionen, die prüfen, ob ein String
	// mit einem anderen endet, und die ein Ende abschneiden.

	return ""
}

// BelowTwenty erwartet eine Zahl von `0` bis `19` und liefert ihr Zahlwort.
//
// TODO
func BelowTwenty(n int) string {
	// Hinweis:
	// - Für `0` bis `9` gibt es schon `Ones`.
	// - Legen Sie für `10` bis `19` wie in `Ones` eine Liste an.
	//   Welcher Index gehört dann zu `n`?

	return ""
}

// BelowHundred erwartet eine Zahl von `0` bis `99` und liefert ihr
// Zahlwort, z.B. "zwanzig", "einundzwanzig" oder "siebenunddreißig".
//
// TODO
func BelowHundred(n int) string {
	// Hinweis:
	// - Unter `20` hilft `BelowTwenty`.
	// - Zerlegen Sie die Zahl sonst mit `/` und `%` in Zehner- und Einerstelle.
	// - Die Zehnerwörter ("zwanzig", "dreißig", ...) passen in eine Liste,
	//   deren Index die Zehnerstelle ist.
	// - Ist die Einerstelle nicht `0`, steht sie vorne, verbunden mit "und".
	//   Denken Sie an `CompoundForm`.

	return ""
}

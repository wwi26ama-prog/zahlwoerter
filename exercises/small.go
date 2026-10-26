package exercises

import "strings"

// Ones erwartet eine Ziffer von `0` bis `9` und liefert ihr Zahlwort,
// also "null", "eins", "zwei", ..., "neun".
func Ones(n int) string {
	// Hinweis:
	// Legen Sie die zehn Wörter in einer Liste ab.
	// Die Ziffer ist dann genau der Index des passenden Worts.

	words := []string{"null", "eins", "zwei", "drei", "vier",
		"fünf", "sechs", "sieben", "acht", "neun"}
	return words[n]
}

// CompoundForm erwartet ein Zahlwort und liefert die Form, die vor einem
// weiteren Wortteil steht: Ein "eins" am Ende wird zu "ein", wie in
// "einundzwanzig", "einhundert" oder "einhunderteintausend".
// Alle anderen Wörter bleiben unverändert.
func CompoundForm(word string) string {
	// Hinweis:
	// Das Paket `strings` enthält Funktionen, die prüfen, ob ein String
	// mit einem anderen endet, und die ein Ende abschneiden.

	if strings.HasSuffix(word, "eins") {
		return strings.TrimSuffix(word, "s")
	}
	return word
}

// BelowTwenty erwartet eine Zahl von `0` bis `19` und liefert ihr Zahlwort.
func BelowTwenty(n int) string {
	// Hinweis:
	// - Für `0` bis `9` gibt es schon `Ones`.
	// - Legen Sie für `10` bis `19` wie in `Ones` eine Liste an.
	//   Welcher Index gehört dann zu `n`?

	if n < 10 {
		return Ones(n)
	}
	words := []string{"zehn", "elf", "zwölf", "dreizehn", "vierzehn",
		"fünfzehn", "sechzehn", "siebzehn", "achtzehn", "neunzehn"}
	return words[n-10]
}

// BelowHundred erwartet eine Zahl von `0` bis `99` und liefert ihr
// Zahlwort, z.B. "zwanzig", "einundzwanzig" oder "siebenunddreißig".
func BelowHundred(n int) string {
	// Hinweis:
	// - Unter `20` hilft `BelowTwenty`.
	// - Zerlegen Sie die Zahl sonst mit `/` und `%` in Zehner- und Einerstelle.
	// - Die Zehnerwörter ("zwanzig", "dreißig", ...) passen in eine Liste,
	//   deren Index die Zehnerstelle ist.
	// - Ist die Einerstelle nicht `0`, steht sie vorne, verbunden mit "und".
	//   Denken Sie an `CompoundForm`.

	if n < 20 {
		return BelowTwenty(n)
	}
	tens := []string{"", "", "zwanzig", "dreißig", "vierzig",
		"fünfzig", "sechzig", "siebzig", "achtzig", "neunzig"}
	if n%10 == 0 {
		return tens[n/10]
	}
	return CompoundForm(Ones(n%10)) + "und" + tens[n/10]
}

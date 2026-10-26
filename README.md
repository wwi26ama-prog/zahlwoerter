# Zahlen in Worten

In dieser Einheit entsteht schrittweise eine Funktion, die eine Zahl
als deutsches Zahlwort schreibt: Aus `123` wird "Einhundertdreiundzwanzig".

Die Übungsaufgaben in der Reihenfolge, in der sie aufeinander aufbauen:

- [`exercises/small.go`](exercises/small.go): Zahlwörter bis `99`.
- [`exercises/hundreds.go`](exercises/hundreds.go): Zahlwörter bis `999`.
- [`exercises/number_word.go`](exercises/number_word.go): das Ziel, zuerst
  für Zahlen bis `999`.
- [`exercises/thousands.go`](exercises/thousands.go): Zahlwörter bis `999 999`.
- [`exercises/millions.go`](exercises/millions.go): Zahlwörter bis `999 999 999`.

Erweitern Sie `NumberWord` nach den letzten beiden Schritten.

Weiterführende Aufgabe: Millionen, die auf `01` enden, z.B. `101 000 000`,
prüfen die Tests nicht. Finden Sie heraus, wie sie geschrieben werden
(z.B. "einhunderteine Million" oder "einhunderteine Millionen"),
schreiben Sie dafür einen eigenen Test und erweitern Sie `BelowBillion`.

Die Aufgaben stehen in den Kommentaren der Funktionen, die erwarteten
Ergebnisse in den zugehörigen Tests. Prüfen mit `go test ./...` in diesem
Verzeichnis, einzelne Tests z.B. mit `go test ./... -run ExampleNumberWord$`.

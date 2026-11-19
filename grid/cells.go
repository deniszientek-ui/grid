// Package grid enthält Funktionen für Spielfelder aus Zeilen und Spalten,
// wie sie z. B. Tic Tac Toe oder Vier Gewinnt brauchen.
package grid

// Diese Aufgaben sind die Grundlagen: ein Spielfeld anlegen, seine Größe
// abfragen, einzelne Zellen lesen und schreiben und das Spielfeld ausgeben.
// Ein Spielfeld ist eine Liste von Zeilen, jede Zeile eine Liste von Zellen,
// also `[][]string`. Zeilen und Spalten werden ab `0` gezählt.

// New erwartet die Anzahl der Zeilen `rows` und Spalten `cols` und einen Wert
// `value` und liefert ein neues Spielfeld, in dem jede Zelle `value` enthält.
// Mit `"."` als Wert entsteht z. B. ein leeres Spielfeld.
//
// TODO
func New(rows, cols int, value string) [][]string {
	// Ein leeres Grid erzeugen
	Spielfeld := [][]string{}

	// Die Zeilen erzeugen
	for i := 0; i < rows; i++ {

		Spielfeld = append(Spielfeld, []string{})

		for n := 0; n < cols; n++ {
			Spielfeld[i] = append(Spielfeld[i], value)
		}

	}
	return Spielfeld
}

// Rows erwartet ein Spielfeld und liefert die Anzahl seiner Zeilen.
//
// TODO
func Rows(cells [][]string) int {
	return len(cells)
}

// Cols erwartet ein Spielfeld und liefert die Anzahl seiner Spalten,
// also die Länge der ersten Zeile. Ein Spielfeld ohne Zeilen hat `0` Spalten.
//
// TODO
func Cols(cells [][]string) int {
	// len geht auch aber Rows macht das bereits in der Funktion vorher
	if Rows(cells) == 0 {
		return 0
	} else {

		return len(cells[0])
	}
}

// Get erwartet ein Spielfeld und die Zeile `row` und Spalte `col` einer Zelle
// und liefert ihren Inhalt.
//
// TODO
func Get(cells [][]string, row, col int) string {
	return cells[row][col]
}

// Set erwartet ein Spielfeld, die Zeile `row` und Spalte `col` einer Zelle
// und einen Wert `value` und schreibt den Wert in die Zelle.
// Die Änderung ist auch beim Aufrufer zu sehen.
//
// TODO
func Set(cells [][]string, row, col int, value string) {
	cells[row][col] = value
}

// Format erwartet ein Spielfeld und liefert es als Text zur Ausgabe:
// eine Textzeile pro Zeile des Spielfelds, die Zellen durch Leerzeichen
// getrennt, jede Textzeile mit einem Zeilenumbruch "\n" am Ende.
//
// TODO
func Format(cells [][]string) string {
	ausgabe := ""

	// Alle Zeilen ausgeben
	for i := range len(cells) {
		for j := range len(cells[i]) {

			ausgabe = ausgabe + cells[i][j] + " "
		}
		if len(ausgabe) > 0 {
			ausgabe = ausgabe[:len(ausgabe)-1]
		}
		ausgabe = ausgabe + "\n"
	}

	return ausgabe
}

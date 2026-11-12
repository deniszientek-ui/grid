package grid

// Diese Aufgaben prüfen eine Linie oder das ganze Spielfeld auf einen Wert,
// z. B. ob eine Zeile nur ein bestimmtes Symbol enthält
// oder ob noch ein Feld frei ist.

// AllEqual erwartet eine Linie, z. B. eine Zeile oder Spalte, und einen Wert
// und liefert `true`, wenn jede Zelle der Linie diesen Wert enthält.
// Für eine leere Linie liefert sie `true`, weil keine Zelle widerspricht.
//
// TODO
func AllEqual(line []string, value string) bool {

	richtig := true

	for i := range len(line) {
		if line[i] == value {
			richtig = true
		} else {
			return false
		}
	}

	return richtig
}

// Contains erwartet ein Spielfeld und einen Wert und liefert `true`,
// wenn mindestens eine Zelle diesen Wert enthält.
//
// TODO
func Contains(cells [][]string, value string) bool {
	return false
}

package board

// MaxCellWidth bestimmt die maximale Breite des Inhalts einer Zelle.
// Dies ist nützlich, um die Spaltenbreite für die Ausgabe zu berechnen.
func (b *Board) MaxCellWidth() int {
	maxWidth := 0
	// TODO
	return maxWidth
}

// / Width bestimmt die Breite des Spields. D.h. die Anzahl der Spalten.
func (b *Board) Width() int {
	// TODO
	return 0
}

// Height bestimmt die Höhe des Spields. D.h. die Anzahl der Zeilen.
func (b *Board) Height() int {
	// TODO
	return 0
}

// Row liefert die Zeile an der angegebenen Position zurück.
func (b *Board) Row(index int) []string {
	// TODO
	return nil
}

// Col liefert die Spalte an der angegebenen Position zurück.
func (b *Board) Col(index int) []string {
	col := make([]string, len(b.rows))
	// TODO
	return col
}

// DiagDownRight liefert eine Diagonale von oben links nach unten rechts zurück.
func (b *Board) DiagDownRight(startCol int) []string {
	diag := []string{}
	// TODO
	return diag
}

// DiagDownLeft liefert eine Diagonale von oben rechts nach unten links zurück.
func (b *Board) DiagDownLeft(startCol int) []string {
	diag := []string{}
	// TODO
	return diag
}

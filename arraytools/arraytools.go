package arraytools

// Mult erwartet ein Array und einen skalaren Faktor.
// Multipliziert jedes Element des Arrays mit dem Faktor.
func ScalarMult(a []float64, factor float64) {
	// TODO

	for pos, wert := range a {

		a[pos] = wert * factor

	}

}

// Add erwartet zwei Arrays der gleichen Länge.
// Addiert die Elemente der Arrays paarweise.
func Add(a, b []float64) {

	for pos := range a {

		a[pos] += b[pos]
	}

}

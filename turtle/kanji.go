package turtle

// A Kanji contains information about a kanji.
//
// ref: https://en.wikipedia.org/wiki/List_of_j%C5%8Dy%C5%8D_kanji
type Kanji struct {
	// The rune of the kanji.
	Rune string
	// The old rune of the kanji.
	OldRune string
	// The radical of the kanji.
	Radical string
	// Grade at which the kanji is studdied.
	Grade int
	// Meaning of the kanji.
	Meaning string
	// Pronounciation of the radical.
	Pronounciation string
}

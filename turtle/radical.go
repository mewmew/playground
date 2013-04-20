package turtle

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// A Radical contains information about a radical.
//
// ref: https://en.wikipedia.org/wiki/Kangxi_radical#Table_of_radicals
type Radical struct {
	// The radical's number.
	Num int
	// The rune of the radical.
	Rune string
	// Alternative rune representations of the radical.
	AltRune string
	// Number of strokes required to draw the radical.
	Strokes int
	// Pinyin pronunciation of the radical.
	Pinyin string
	// Hiragana pronunciation of the radical.
	Hiragana string
	// Romaji pronunciation of the radical.
	Romaji string
	// Meaning of the radical.
	Meaning string
	// Frequency of use based on the 47 035 characters listed in the Kangxi
	// dictionary.
	Freq int
	// Simplified version of the radical.
	Simplified string
	// Examples using the radical.
	Examples string
}

// GetRadical locates and returns the radical, either by rune, by radical number
// or by meaning.
func GetRadical(s string) (radical *Radical, err error) {
	// Locate by rune of radical.
	radical, ok := radicalRune[s]
	if ok {
		return radical, nil
	}

	// Locate by radical number.
	radical, ok = radicalNum[s]
	if ok {
		return radical, nil
	}

	// Locate by meaning of radical.
	radical, ok = radicalMeaning[s]
	if ok {
		return radical, nil
	}

	return nil, fmt.Errorf("turtle.GetRadical: unable to locate radical %q.", s)

}

// radicalRune maps from the rune of a radical to the struct containing
// additional information about it.
var radicalRune = make(map[string]*Radical)

// radicalNum maps from a radical's number to the struct containing additional
// information about it.
var radicalNum = make(map[string]*Radical)

// radicalMeaning maps from the meaning of a radical to the struct containing
// additional information about it. Note that some radicals have more than one
// meaning, in which case all of those map to the same radical.
var radicalMeaning = make(map[string]*Radical)

func init() {
	for _, radical := range radicals {
		// map from rune to radical.
		radicalRune[radical.Rune] = radical

		// map from num to radical.
		radicalNum[strconv.Itoa(radical.Num)] = radical

		// map from meaning(s) to radical.
		meanings := strings.Split(radical.Meaning, ",")
		for _, meaning := range meanings {
			meaning = strings.TrimSpace(meaning)
			dupe, ok := radicalMeaning[meaning]
			if ok {
				err := fmt.Errorf("turtle.init: both %q and %q have the meaning %q.", dupe.Rune, radical.Rune, meaning)
				log.Println(err)
			}
			radicalMeaning[meaning] = radical
		}
	}
}

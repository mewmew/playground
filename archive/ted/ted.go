// Package ted declares the type used to represent a TED talk.
package ted

import "time"

// Talk represents a TED talk.
type Talk struct {
	Date     time.Time     // Aug 2009
	Event    string        // TEDGlobal 2009
	Title    string        // Janine Benyus: Biomimicry in action
	Duration time.Duration // 17:42
	Download string        // http://download.ted.com/talks/JanineBenyus_2009G-480p.mp4?apikey=TEDDOWNLOAD
}

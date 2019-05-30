package wallhaven

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// An Option represents a search option.
type Option interface {
	// Key returns the key of the search parameter.
	Key() string
	// Value returns the value of the search parameter.
	Value() string
}

// --- categories

// Categories specifies the enabled wallpaper categories of a search query.
type Categories uint32

// Wallpaper categories.
const (
	CatGeneral Categories = 1 << iota
	CatAnime
	CatPeople
)

// String returns a string representation of the search option.
func (v Categories) String() string {
	catNames := map[Categories]string{
		CatGeneral: "general",
		CatAnime:   "anime",
		CatPeople:  "people",
	}
	var names []string
	for mask := CatGeneral; mask <= CatPeople; mask <<= 1 {
		if v&mask != 0 {
			name := catNames[mask]
			names = append(names, name)
		}
	}
	return strings.Join(names, ",")
}

// Set implements the flag.Value interface. A comma-separated list of the
// following options is valid:
//    * general
//    * anime
//    * people
func (v *Categories) Set(s string) error {
	cats := map[string]Categories{
		"general": CatGeneral,
		"anime":   CatAnime,
		"people":  CatPeople,
	}
	names := strings.Split(s, ",")
	for _, name := range names {
		mask, ok := cats[name]
		if !ok {
			return errors.Errorf("invalid wallpaper category %q", name)
		}
		*v |= mask
	}
	return nil
}

// Key returns the key of the search parameter.
func (v Categories) Key() string {
	return "categories"
}

// Value returns the value of the search parameter.
func (v Categories) Value() string {
	general := "0"
	if v&CatGeneral != 0 {
		general = "1"
	}
	animal := "0"
	if v&CatAnime != 0 {
		animal = "1"
	}
	people := "0"
	if v&CatPeople != 0 {
		people = "1"
	}
	return general + animal + people
}

// --- purity

// Purity specifies the enabled purity modes of a search query.
type Purity uint32

// Purity modes.
const (
	PuritySFW Purity = 1 << iota
	PuritySketchy
	PurityNSFW
)

// String returns a string representation of the search option.
func (v Purity) String() string {
	purityNames := map[Purity]string{
		PuritySFW:     "SFW",
		PuritySketchy: "sketchy",
		PurityNSFW:    "NSFW",
	}
	var names []string
	for mask := PuritySFW; mask <= PurityNSFW; mask <<= 1 {
		if v&mask != 0 {
			name := purityNames[mask]
			names = append(names, name)
		}
	}
	return strings.Join(names, ",")
}

// Set implements the flag.Value interface. A comma-separated list of the
// following options is valid:
//    * SFW
//    * sketchy
//    * NSFW
func (v *Purity) Set(s string) error {
	purity := map[string]Purity{
		"SFW":     PuritySFW,
		"sketchy": PuritySketchy,
		"NSFW":    PurityNSFW,
	}
	names := strings.Split(s, ",")
	for _, name := range names {
		mask, ok := purity[name]
		if !ok {
			return errors.Errorf("invalid purity mode %q", name)
		}
		*v |= mask
	}
	return nil
}

// Key returns the key of the search parameter.
func (v Purity) Key() string {
	return "purity"
}

// Value returns the value of the search parameter.
func (v Purity) Value() string {
	sfw := "0"
	if v&PuritySFW != 0 {
		sfw = "1"
	}
	sketchy := "0"
	if v&PuritySketchy != 0 {
		sketchy = "1"
	}
	nsfw := "0"
	if v&PurityNSFW != 0 {
		nsfw = "1"
	}
	return sfw + sketchy + nsfw
}

// --- resolutions

// Resolutions specifies the enabled screen resolution of a search query.
type Resolutions uint32

// Screen resolutions.
const (
	Res1024x768 Resolutions = 1 << iota
	Res1280x800
	Res1366x768
	Res1280x960
	Res1440x900
	Res1600x900
	Res1280x1024
	Res1600x1200
	Res1680x1050
	Res1920x1080
	Res1920x1200
	Res2560x1440
	Res2560x1600
	Res3840x1080
	Res5760x1080
	Res3840x2160
)

// String returns a string representation of the search option.
func (v Resolutions) String() string {
	resNames := map[Resolutions]string{
		Res1024x768:  "1024x768",
		Res1280x800:  "1280x800",
		Res1366x768:  "1366x768",
		Res1280x960:  "1280x960",
		Res1440x900:  "1440x900",
		Res1600x900:  "1600x900",
		Res1280x1024: "1280x1024",
		Res1600x1200: "1600x1200",
		Res1680x1050: "1680x1050",
		Res1920x1080: "1920x1080",
		Res1920x1200: "1920x1200",
		Res2560x1440: "2560x1440",
		Res2560x1600: "2560x1600",
		Res3840x1080: "3840x1080",
		Res5760x1080: "5760x1080",
		Res3840x2160: "3840x2160",
	}
	var names []string
	for mask := Res1024x768; mask <= Res3840x2160; mask <<= 1 {
		if v&mask != 0 {
			name := resNames[mask]
			names = append(names, name)
		}
	}
	return strings.Join(names, ",")
}

// Set implements the flag.Value interface. A comma-separated list of the
// following options is valid:
//    * 1024x768
//    * 1280x800
//    * 1366x768
//    * 1280x960
//    * 1440x900
//    * 1600x900
//    * 1280x1024
//    * 1600x1200
//    * 1680x1050
//    * 1920x1080
//    * 1920x1200
//    * 2560x1440
//    * 2560x1600
//    * 3840x1080
//    * 5760x1080
//    * 3840x2160
func (v *Resolutions) Set(s string) error {
	res := map[string]Resolutions{
		"1024x768":  Res1024x768,
		"1280x800":  Res1280x800,
		"1366x768":  Res1366x768,
		"1280x960":  Res1280x960,
		"1440x900":  Res1440x900,
		"1600x900":  Res1600x900,
		"1280x1024": Res1280x1024,
		"1600x1200": Res1600x1200,
		"1680x1050": Res1680x1050,
		"1920x1080": Res1920x1080,
		"1920x1200": Res1920x1200,
		"2560x1440": Res2560x1440,
		"2560x1600": Res2560x1600,
		"3840x1080": Res3840x1080,
		"5760x1080": Res5760x1080,
		"3840x2160": Res3840x2160,
	}
	names := strings.Split(s, ",")
	for _, name := range names {
		mask, ok := res[name]
		if !ok {
			return errors.Errorf("invalid screen resolution %q", name)
		}
		*v |= mask
	}
	return nil
}

// Key returns the key of the search parameter.
func (v Resolutions) Key() string {
	return "resolutions"
}

// Value returns the value of the search parameter.
func (v Resolutions) Value() string {
	return v.String()
}

// --- ratios

// Ratio specifies the enabled aspect ratios of a search query.
type Ratios uint32

// Aspect ratios.
const (
	Ratio4x3 Ratios = 1 << iota
	Ratio5x4
	Ratio16x9
	Ratio16x10
	Ratio21x9
	Ratio32x9
	Ratio48x9
)

// String returns a string representation of the search option.
func (v Ratios) String() string {
	ratioNames := map[Ratios]string{
		Ratio4x3:   "4x3",
		Ratio5x4:   "5x4",
		Ratio16x9:  "16x9",
		Ratio16x10: "16x10",
		Ratio21x9:  "21x9",
		Ratio32x9:  "32x9",
		Ratio48x9:  "48x9",
	}
	var names []string
	for mask := Ratio4x3; mask <= Ratio48x9; mask <<= 1 {
		if v&mask != 0 {
			name := ratioNames[mask]
			names = append(names, name)
		}
	}
	return strings.Join(names, ",")
}

// Set implements the flag.Value interface. A comma-separated list of the
// following options is valid:
//    * 4x3
//    * 5x4
//    * 16x9
//    * 16x10
//    * 21x9
//    * 32x9
//    * 48x9
func (v *Ratios) Set(s string) error {
	ratios := map[string]Ratios{
		"4x3":   Ratio4x3,
		"5x4":   Ratio5x4,
		"16x9":  Ratio16x9,
		"16x10": Ratio16x10,
		"21x9":  Ratio21x9,
		"32x9":  Ratio32x9,
		"48x9":  Ratio48x9,
	}
	names := strings.Split(s, ",")
	for _, name := range names {
		mask, ok := ratios[name]
		if !ok {
			return errors.Errorf("invalid aspect ratio %q", name)
		}
		*v |= mask
	}
	return nil
}

// Key returns the key of the search parameter.
func (v Ratios) Key() string {
	return "ratios"
}

// Value returns the value of the search parameter.
func (v Ratios) Value() string {
	return v.String()
}

// --- sorting

// Sorting specifies the sorting method used in search queries.
type Sorting int

// Sorting methods.
const (
	SortRelevance Sorting = 1 + iota
	SortRandom
	SortDateAdded
	SortViews
	SortFavorites
)

// String returns a string representation of the search option.
func (v Sorting) String() string {
	sortingNames := map[Sorting]string{
		SortRelevance: "relevance",
		SortRandom:    "random",
		SortDateAdded: "date_added",
		SortViews:     "views",
		SortFavorites: "favorites",
	}
	return sortingNames[v]
}

// Set implements the flag.Value interface. A comma-separated list of the
// following options is valid:
//    * relevance
//    * random
//    * date_added
//    * views
//    * favorites
func (v *Sorting) Set(s string) error {
	sortings := map[string]Sorting{
		"relevance":  SortRelevance,
		"random":     SortRandom,
		"date_added": SortDateAdded,
		"views":      SortViews,
		"favorites":  SortFavorites,
	}
	sorting, ok := sortings[s]
	if !ok {
		return errors.Errorf("invalid sorting method %q", s)
	}
	*v = sorting
	return nil
}

// Key returns the key of the search parameter.
func (v Sorting) Key() string {
	return "sorting"
}

// Value returns the value of the search parameter.
func (v Sorting) Value() string {
	return v.String()
}

// --- order

// Order specifies the sorting order used in search queries.
type Order int

// Sorting orders.
const (
	OrderAsc Order = 1 + iota
	OrderDesc
)

// String returns a string representation of the search option.
func (v Order) String() string {
	orderNames := map[Order]string{
		OrderAsc:  "asc",
		OrderDesc: "desc",
	}
	return orderNames[v]
}

// Set implements the flag.Value interface. A comma-separated list of the
// following options is valid:
//    * asc
//    * desc
func (v *Order) Set(s string) error {
	orders := map[string]Order{
		"asc":  OrderAsc,
		"desc": OrderDesc,
	}
	order, ok := orders[s]
	if !ok {
		return errors.Errorf("invalid sorting order %q", s)
	}
	*v = order
	return nil
}

// Key returns the key of the search parameter.
func (v Order) Key() string {
	return "order"
}

// Value returns the value of the search parameter.
func (v Order) Value() string {
	return v.String()
}

// --- page

// Page specifies the result page number.
type Page int

// String returns a string representation of the search option.
func (v Page) String() string {
	return strconv.Itoa(int(v))
}

// Set implements the flag.Value interface.
func (v *Page) Set(s string) error {
	x, err := strconv.Atoi(s)
	if err != nil {
		return errors.WithStack(err)
	}
	*v = Page(x)
	return nil
}

// Key returns the key of the search parameter.
func (v Page) Key() string {
	return "page"
}

// Value returns the value of the search parameter.
func (v Page) Value() string {
	return v.String()
}

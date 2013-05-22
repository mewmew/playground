package turtle

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"reflect"
	"sort"
)

// RadicalServer handles requests related to radicals.
type RadicalServer struct {
	tmpl      *template.Template
	freqOrder []*Radical
	rndOrder  []*Radical
}

// NewRadicalServer returns a server capable of handling requests related to
// radicals.
func NewRadicalServer() (srv *RadicalServer, err error) {
	// Load templates.
	tmplPath, err := data.Path("radicals/list.tmpl")
	if err != nil {
		return nil, err
	}
	srv = new(RadicalServer)
	funcs := template.FuncMap{
		"eq": eq,
	}
	srv.tmpl, err = template.New("list.tmpl").Funcs(funcs).ParseFiles(tmplPath)
	if err != nil {
		return nil, err
	}

	// Create a slice of radicals for each sort order.
	// Frequency of use.
	srv.freqOrder = make([]*Radical, len(radicals))
	copy(srv.freqOrder, radicals)
	sort.Sort(FreqOrder(srv.freqOrder))

	// Random order.
	srv.rndOrder = make([]*Radical, len(radicals))
	copy(srv.rndOrder, radicals)

	return srv, nil
}

// eq calculates the equality of a and b.
func eq(a, b interface{}) bool {
	switch a1 := a.(type) {
	case int, string:
		if a1 == b {
			return true
		}
	}
	return reflect.DeepEqual(a, b)
}

// ServeHTTP serves pages related to radicals.
//
// The page of each radical can be accessed in three ways. Either by rune, by
// radical number or by meaning.
//
// If no radical is specified a list of all radicals is displayed.
//
// Examples:
//    * page listing all radicals:
//       "/radicals/"
//    * page of radical 人:
//       "/radicals/人"
//       "/radicals/9"
//       "/radicals/person"
func (srv *RadicalServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s := req.URL.Path[len("/radicals/"):]
	if len(s) == 0 {
		// Serve a list of all radicals.
		srv.ServeList(w, req)
		return
	}

	// Serve the specified radical's page.
	srv.ServeRadical(w, req)
}

// ServeList serves a list of all radicals. The list can be ordered in three
// ways. Either by radical number (default), frequency of use [1] or in a random
// order.
//
// Examples:
//    * Order by radical number (default):
//       "/radicals/"
//       "/radicals/?order=num"
//    * Order by frequency of use:
//       "/radicals/?order=freq"
//    * Random order:
//       "/radicals/?order=rnd"
//
// [1]: https://en.wikipedia.org/wiki/Kangxi_radical#Table_of_radicals
func (srv *RadicalServer) ServeList(w http.ResponseWriter, req *http.Request) {
	t := srv.tmpl.Lookup("list.tmpl")

	// Sort order.
	m := make(map[string]interface{})
	order := req.FormValue("order")
	m["order"] = order
	switch order {
	case "freq":
		// Frequency of use.
		m["radicals"] = srv.freqOrder
	case "rnd":
		// Random order.
		sort.Sort(RndOrder(srv.rndOrder))
		m["radicals"] = srv.rndOrder
	default:
		// Radical number.
		m["radicals"] = radicals
	}
	err := t.Execute(w, m)
	if err != nil {
		log.Println(err)
		http.NotFound(w, req)
		return
	}
}

// ServeRadical serves a page containing information and mnemonic clues about
// the specified radical. On this page there will be links to the previous and
// the next radical, based on the specified order.
//
// Examples:
//    * Show links to next and prev based on radical number (default):
//       "/radicals/人"
//       "/radicals/人?order=num"
//    * Show links to next and prev based on frequency of use:
//       "/radicals/人?order=freq"
//    * Show random links to next and prev:
//       "/radicals/人?order=rnd"
func (srv *RadicalServer) ServeRadical(w http.ResponseWriter, req *http.Request) {
	s := req.URL.Path[len("/radicals/"):]
	radical, err := GetRadical(s)
	if err != nil {
		log.Println(err)
		http.NotFound(w, req)
		return
	}
	/// ### TODO ###
	_ = radical
	log.Println("RadicalServer.ServeRadical: not yet implemented.")
}

// ServeData serves files from the data directory.
func ServeData(w http.ResponseWriter, req *http.Request) {
	filePath, err := data.Path(req.URL.Path)
	if err != nil {
		log.Println(err)
		http.NotFound(w, req)
		return
	}
	if Verbose {
		fmt.Println("turtle.ServeData: serving file:", filePath)
	}
	http.ServeFile(w, req, filePath)
}

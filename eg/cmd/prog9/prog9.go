package main

import "flag"
import "log"
import "os"

import "github.com/mewmew/playground/eg"

func init() {
	flag.StringVar(&eg.FieldV4, "f", "", "Set enigmafiedV4 cookie value.")
	flag.StringVar(&eg.PhpSessid, "p", "", "Set PHPSESSID cookie value.")
	flag.Parse()
}

func main() {
	if !eg.HasSession() {
		flag.Usage()
		os.Exit(1)
	}
	err := prog9()
	if err != nil {
		log.Fatalln(err)
	}
}

// prog9 parses the game_grid, solves the equation and submits the answer.
//
// 1) Download:
//    - Download the HTML page.
// 2) Parse:
//    - Parse the #game_grid in the HTML page.
// 3) Calculate:
//    - Locate equations with only one slot it them.
//    - Propagate the information into the other equations, thereby eliminating
//      one shape.
// 4) Submit answer.
//    - Submit the answer.
func prog9() (err error) {
	buf, err := eg.Get("http://www.enigmagroup.org/missions/programming/9/")
	if err != nil {
		return err
	}
	es, err := ParseGrid(string(buf))
	if err != nil {
		return err
	}
	Print(es)
	for es.Len() > 0 {
		Calc(es)
	}
	err = submitAnswer()
	if err != nil {
		return err
	}
	return nil
}

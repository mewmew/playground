package main

import "flag"
import "fmt"
import "log"
import "os"
import "strings"

import "github.com/mewmew/foohash"
import "github.com/mewmew/foohash/google"
import "github.com/mewmew/foohash/mutation"
import "github.com/mewmew/foohash/wordlist"

var flagWordlistPath string
var flagRegular bool
var flagTitle bool
var flagUpper bool
var flagLeet bool
var flagNum bool
var flagSaltPrefix string
var flagSaltSuffix string
var flagGoogle bool

func init() {
	flag.StringVar(&flagWordlistPath, "w", "", "Wordlist path.")
	flag.BoolVar(&flagRegular, "wr", true, "Wordlist attack, using regular (unmodified) words.")
	flag.BoolVar(&flagTitle, "wt", true, "Wordlist attack, using titled words.")
	flag.BoolVar(&flagUpper, "wu", true, "Wordlist attack, using upper case words.")
	flag.BoolVar(&flagLeet, "wl", true, "Wordlist attack, using leet speak words.")
	flag.BoolVar(&flagNum, "wn", true, "Wordlist attack, using words with number suffixes.")
	flag.StringVar(&flagSaltPrefix, "sp", "", "Salt prefix.")
	flag.StringVar(&flagSaltSuffix, "ss", "", "Salt suffix.")
	flag.BoolVar(&flagGoogle, "g", false, "Google.")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: foohash [OPTION]... [HASH]...")
	fmt.Fprintln(os.Stderr, "Recover passwords from hashes.")
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		return
	}
	for _, rawHash := range flag.Args() {
		if flagGoogle {
			google.Search(rawHash)
			continue
		}
		if len(flagWordlistPath) > 0 {
			pass, found, err := CheckWordlist(rawHash, flagWordlistPath)
			if err != nil {
				log.Fatalln(err)
			}
			if found {
				fmt.Println("hash:", rawHash)
				fmt.Println("pass:", pass)
				continue
			}
		}
	}
	fmt.Println("checked:", foohash.CheckCount)
}

func CheckWordlist(rawHash, wordlistPath string) (pass string, found bool, err error) {
	hash, err := foohash.New(rawHash)
	if err != nil {
		return "", false, err
	}
	hash.SetSalt(flagSaltPrefix, flagSaltSuffix)
	words, err := wordlist.New(wordlistPath)
	if err != nil {
		return "", false, err
	}
	pass, found = CheckMutations(hash, words)
	return pass, found, nil
}

// CheckMutations checks the hash against a wordlist, using a set of
// permutations. The following permutations are performed, in the same order as
// they are listed:
//    - regular
//    - title
//    - upper case
//    - leet
//    - number suffix (for all of the above)
func CheckMutations(hash *foohash.Hash, words wordlist.Wordlist) (pass string, found bool) {
	if flagRegular {
		// regular
		pass, found = words.Check(hash)
		if found {
			return pass, true
		}
	}
	var titleWords, upperWords, leetWords wordlist.Wordlist
	if flagTitle {
		// title
		titleWords = words.Mutation(strings.Title)
		pass, found = titleWords.Check(hash)
		if found {
			return pass, true
		}
	}
	if flagUpper {
		// upper case
		upperWords = words.Mutation(strings.ToUpper)
		pass, found = upperWords.Check(hash)
		if found {
			return pass, true
		}
	}
	if flagLeet {
		// leet
		leetWords = words.Mutation(mutation.Leet)
		pass, found = leetWords.Check(hash)
		if found {
			return pass, true
		}
	}
	if flagNum {
		pass, found = CheckNumSuffixes(hash, words, titleWords, upperWords, leetWords)
		if found {
			return pass, true
		}
	}
	return "", false
}

func CheckNumSuffixes(hash *foohash.Hash, lists ...wordlist.Wordlist) (pass string, found bool) {
	for _, list := range lists {
		// number suffixes
		numWords := list.Mutations(mutation.NumSuffixes)
		pass, found = numWords.Check(hash)
		if found {
			return pass, true
		}
	}
	return "", false
}

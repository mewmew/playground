package wordlist

import "github.com/mewkiz/pkg/bufioutil"
import "github.com/mewmew/foohash"

type Wordlist []string

func New(wordlistPath string) (words Wordlist, err error) {
	words, err = bufioutil.ReadLines(wordlistPath)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func (words Wordlist) Mutation(mutateFn func(string) string) (dstWords Wordlist) {
	for _, word := range words {
		dstWords = append(dstWords, mutateFn(word))
	}
	return dstWords
}

func (words Wordlist) Mutations(mutateFn func(string) []string) (dstWords Wordlist) {
	for _, word := range words {
		dstWords = append(dstWords, mutateFn(word)...)
	}
	return dstWords
}

func (words Wordlist) Check(hash *foohash.Hash) (pass string, found bool) {
	for _, pass = range words {
		if hash.IsPlain(pass) {
			return pass, true
		}
	}
	return "", false
}

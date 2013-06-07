package foohash

import (
	"crypto/sha1"
)

func newSha1(pass string) (hash *Hash) {
	h := sha1.New()
	h.Write([]byte(pass))
	hash = new(Hash)
	hash.buf = h.Sum(nil)
	hash.newFn = newSha1
	return hash
}

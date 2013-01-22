package foohash

import "crypto/md5"

func newMd5(pass string) (hash *Hash) {
	h := md5.New()
	h.Write([]byte(pass))
	hash = new(Hash)
	hash.buf = h.Sum(nil)
	hash.newFn = newMd5
	return hash
}

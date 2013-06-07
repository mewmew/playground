package foohash

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

type Hash struct {
	buf   []byte
	newFn func(string) *Hash
	salt  Salt
}

type Salt struct {
	prefix, suffix string
}

func New(rawHash string) (hash *Hash, err error) {
	hash = new(Hash)
	hash.buf, err = hex.DecodeString(rawHash)
	if err != nil {
		return nil, err
	}
	switch len(rawHash) {
	case 32:
		hash.newFn = newMd5
	case 40:
		hash.newFn = newSha1
	default:
		return nil, fmt.Errorf("invalid hash length (%d).", len(rawHash))
	}
	return hash, nil
}

func (hash *Hash) String() string {
	return hex.EncodeToString(hash.buf)
}

func (hash *Hash) SetSalt(prefix, suffix string) {
	hash.salt.prefix = prefix
	hash.salt.suffix = suffix
}

var CheckCount int

func (hash *Hash) IsPlain(pass string) bool {
	pass = hash.salt.prefix + pass + hash.salt.suffix
	newHash := hash.newFn(pass)
	CheckCount++
	return bytes.Equal(hash.buf, newHash.buf)
}

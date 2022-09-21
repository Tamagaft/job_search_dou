package types

import (
	"crypto/md5"
	"fmt"
)

type Work struct {
	Id       int
	Title    string
	Company  string
	Cities   []string
	Link     string
	Category string
	Hash     string
}

func (w *Work) SetHash() {
	w.Hash = fmt.Sprintf("%x", md5.Sum([]byte(w.Link)))
}

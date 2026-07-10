package rns

import (
	"strings"

	chrs "github.com/fwarcis/go-chars"
)

type Pred = func(rune) bool

func Eq[C chrs.Chars](comparings C) Pred {
	pos := 0
	runes := []rune(string(comparings))
	return func(rn rune) bool {
		if pos == len(runes) {
			return false
		}
		equals := runes[pos] == rn
		pos++
		return equals
	}
}

func In[C chrs.Chars](characters C) Pred {
	return func(rn rune) bool {
		return strings.ContainsRune(
			string(characters), rn)
	}
}

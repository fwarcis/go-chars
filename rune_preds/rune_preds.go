package rune_preds

import (
	"strings"

	"github.com/fwarcis/go-chars"
	"github.com/fwarcis/go-preds/preds"
)

type Pred = preds.Pred[rune]

func Iter[C chars.Chars](comparing C) Pred {
	pos := 0
	runes := []rune(string(comparing))
	return func(rn rune) bool {
		if pos == len(runes) {
			return false
		}
		equals := runes[pos] == rn
		pos++
		return equals
	}
}

func Find[C chars.Chars](comparing C) Pred {
	return func(rn rune) bool {
		return strings.ContainsRune(
			string(comparing), rn,
		)
	}
}

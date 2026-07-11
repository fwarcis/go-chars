package rns

import (
	"iter"

	chrs "github.com/fwarcis/go-chars"
	"github.com/fwarcis/go-chars/rn_prds"
)

func Runes[C rune | chrs.Chars](iterating C) iter.Seq[rune] {
	runes := []rune(string(iterating))
	return func(yield rn_prds.Pred) {
		for _, r := range runes {
			if !yield(r) {
				return
			}
		}
	}
}

func All[C rune | chrs.Chars](iterating C) iter.Seq2[int, rune] {
	runes := []rune(string(iterating))
	return func(yield func(int, rune) bool) {
		for i, r := range runes {
			if !yield(i, r) {
				return
			}
		}
	}
}

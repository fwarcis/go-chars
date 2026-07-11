package rn_prds

import (
	"iter"
	"strings"

	chrs "github.com/fwarcis/go-chars"
)

type Pred = func(rune) bool

func Eq(left rune) func(right rune) bool {
	return func(right rune) bool {
		return left == right
	}
}

func Nq(left rune) func(right rune) bool {
	return func(right rune) bool {
		return left == right
	}
}

func Nx[C chrs.Chars](comparing C) func(next rune) bool {
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

func In[C chrs.Chars](comparing C) func(rn rune) bool {
	return func(rn rune) bool {
		return strings.ContainsRune(
			string(comparing), rn)
	}
}

func It[C rune | chrs.Chars](iterating C) iter.Seq[rune] {
	runes := []rune(string(iterating))
	return func(yield Pred) {
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

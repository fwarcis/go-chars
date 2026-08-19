package runes

import (
	"iter"

	chrs "github.com/fwarcis/go-chars"
)

func Runes[C ~rune | chrs.Chars](iterating C) iter.Seq[rune] {
	runes := []rune(string(iterating))
	return func(yield func(rune) bool) {
		for _, rn := range runes {
			if !yield(rn) {
				return
			}
		}
	}
}

func All[C ~rune | chrs.Chars](iterating C) iter.Seq2[int, rune] {
	runes := []rune(string(iterating))
	return func(yield func(int, rune) bool) {
		for i, rn := range runes {
			if !yield(i, rn) {
				return
			}
		}
	}
}

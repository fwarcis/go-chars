package rns

import (
	"iter"

	chrs "github.com/fwarcis/go-chars"
)

func Runes[R ~rune, C ~rune | chrs.Chars](iterating C) iter.Seq[R] {
	runes := []rune(string(iterating))
	return func(yield func(R) bool) {
		for _, r := range runes {
			if !yield(R(r)) {
				return
			}
		}
	}
}

func All[R ~rune, C ~rune | chrs.Chars](iterating C) iter.Seq2[int, R] {
	runes := []rune(string(iterating))
	return func(yield func(int, R) bool) {
		for i, r := range runes {
			if !yield(i, R(r)) {
				return
			}
		}
	}
}

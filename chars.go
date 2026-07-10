package chrs

type Chars interface {
	~string | ~[]byte | ~[]rune
}

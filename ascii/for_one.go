package ascii

func IsLatin(rn rune) bool {
	return (rn >= 'a' && rn <= 'z') ||
		(rn >= 'A' && rn <= 'Z')
}

func IsDigit(rn rune) bool {
	return rn >= '0' && rn <= '9'
}

func IsAlphaNumeric(rn rune) bool {
	return (rn >= 'a' && rn <= 'z') ||
		(rn >= 'A' && rn <= 'Z') ||
		(rn >= '0' && rn <= '9')
}

func IsUpper(rn rune) bool {
	return rn >= 'A' && rn <= 'Z'
}

func IsLower(rn rune) bool {
	return rn >= 'a' && rn <= 'z'
}

func IsCapitalized(rn rune) bool {
	return rn >= 'A' && rn <= 'Z'
}

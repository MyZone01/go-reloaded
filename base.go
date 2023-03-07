package goreloaded

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	for i := 0; i <= len(s)-len(toFind); i++ {
		subs := s[i : i+len(toFind)]
		if subs == toFind {
			return i
		}
	}
	return -1
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return NbrBase(AtoiBase(nbr, baseFrom), baseTo)
}

func AtoiBase(s string, base string) int {
	baseLen := len(base)
	number := 0
	factor := 1
	isNegative := false
	if s[0] == '-' {
		s = s[1:]
		isNegative = true
	}
	numberLen := len(s)
	for i := numberLen - 1; i >= 0; i-- {
		number += Index(base, string(s[i])) * factor
		if !isNegative && number < 0 {
			number = -(number + 1)
		}
		factor *= baseLen
	}
	if isNegative {
		number = -number
	}
	return number
}

func NbrBase(nbr int, base string) string {
	baseLen := len(base)
	number := ""
	isNegative := false
	if nbr < 0 {
		isNegative = true
	}
	if nbr != 0 {
		for nbr != 0 {
			mod := nbr % baseLen
			if mod < 0 {
				mod = -mod
			}
			number += string(base[mod])
			nbr /= baseLen
		}
	} else {
		number = string(base[0])
	}
	if isNegative {
		number += "-"
	}
	number = StrRev(number)
	return number
}

func StrRev(s string) string {
	reverse := []rune(s)
	for i := 0; i < len(s)/2; i++ {
		reverse[i], reverse[len(s)-1-i] = reverse[len(s)-1-i], reverse[i]
	}
	return string(reverse)
}

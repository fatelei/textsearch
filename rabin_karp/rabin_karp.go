package RabinKarp

import (
	"encoding/binary"
	"math"
	"unicode/utf8"
)

// RabinKarp struct
type RabinKarp struct {
	base         float64
	module       float64
	curHashValue float64
}

// EPSILON is used to compare two float value.
const EPSILON float64 = 0.00000001

// Compare float value.
func floatEquals(a, b float64) bool {
	if diff := math.Abs(a - b); diff < EPSILON {
		return true
	}

	return false
}

// Convert rune to float.
func (p *RabinKarp) runeToFloat(text string, r rune, index int) float64 {
	var buf [utf8.UTFMax]byte
	rl := utf8.RuneLen(r)

	copy(buf[:], text[index:index+rl])
	rst, _ := binary.Uvarint(buf[:])
	return float64(rst)
}

// Hash function.
func (p *RabinKarp) hash(text string, length int) float64 {
	var value float64

	for index, s := range text {
		rst := p.runeToFloat(text, s, index)
		value += rst * math.Pow(p.base, float64(length-1-index))
	}

	return math.Mod(value, p.module)
}

// Search.
func (p *RabinKarp) search(text string, pattern string) bool {
	length := len(pattern)
	textLength := len(text)
	patternHashValue := p.hash(pattern, length)

	for index := range text {
		if index+length > textLength {
			return false
		}

		if index == 0 {
			p.curHashValue = p.hash(text[index:index+length], length)
		} else {
			beforeRuneValue := p.runeToFloat(text, rune(text[index-1]), index-1)
			tmp := p.base * (p.curHashValue - beforeRuneValue*math.Pow(p.base, float64(length-1)))
			p.curHashValue = tmp + p.runeToFloat(text, rune(text[index+length-1]), index+length-1)
		}

		if floatEquals(patternHashValue, p.curHashValue) && text[index:index+length] == pattern {
			return true
		}
	}
	return false
}

package RabinKarp

import "testing"

func TestRuneToFloat(t *testing.T) {
	p := RabinKarp{base: 101, module: 1 << 64}
	var rst float64
	rst = p.runeToFloat("世界 means world", '世', 0)

	if rst != 367716 {
		t.Fail()
	}

	rst = p.runeToFloat("世界 means world", '界', 3)
	if rst != 199399 {
		t.Fail()
	}
}

func TestHash(t *testing.T) {
	p := RabinKarp{base: 101, module: 1 << 64}
	var rst float64
	rst = p.hash("abc", 3)

	if rst != 999494 {
		t.Fail()
	}
}

func TestSearch(t *testing.T) {
	var rst bool
	p := RabinKarp{base: 101, module: 1 << 64}
	rst = p.search("abcd", "d")
	if !rst {
		t.Fail()
	}

	rst = p.search("abcd", "bcd")
	if !rst {
		t.Fail()
	}

	rst = p.search("世界 means world", "世界")
	if !rst {
		t.Fail()
	}
}

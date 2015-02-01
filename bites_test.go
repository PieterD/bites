package bites

import (
	"testing"
)

func TestCapacity(t *testing.T) {
	b := Empty()[0:0:0].Capacity(16)
	if len(b) != 0 {
		t.Fatalf("FAIL! Expected len 0, got %d", len(b))
	}
	if cap(b) != 16 {
		t.Fatalf("FAIL! Expected cap 16, got cap %d", cap(b))
	}
	b = b.Extend(16, true)
	if cap(b) != 16 || len(b) != 16 {
		t.Fatalf("FAIL! Invalid len(%d) or cap(%d), expected 16 for both", len(b), cap(b))
	}
	b = b.Reuse()
	if cap(b) != 16 {
		t.Fatalf("FAIL! Reuse should not change cap")
	}
	if len(b) != 0 {
		t.Fatalf("FAIL! Reuse should set len to 0")
	}
	b = b.Capacity(10)
	if cap(b) != 16 {
		t.Fatalf("FAIL! Short capacity should not change cap")
	}
}

func TestExtendShort(t *testing.T) {
	b := make(Bites, 50).Set(1)
	for i := range b {
		if b[i] != 1 {
			t.Fatalf("FAIL! Set did not set everything")
		}
	}

	origcap := cap(b)
	b = b.Reuse().Extend(20, true)
	if cap(b) != origcap {
		t.Fatalf("FAIL! Short extend changed cap")
	}
	if len(b) != 20 {
		t.Fatalf("FAIL! Short extend did not set len properly, got %d expect %d", len(b), 20)
	}
	for i := range b {
		if b[i] != 0 {
			t.Fatalf("FAIL! Short extend did not set everything to 0")
		}
	}
}

func TestExtendMid(t *testing.T) {
	b := make(Bites, 50).Set(1)
	origcap := cap(b)
	b = b.Reuse().Extend(500, true)
	if cap(b) == origcap {
		t.Fatalf("FAIL! Mid extend did not change cap")
	}
	if len(b) != 500 {
		t.Fatalf("FAIL! Mid extend did not set len properly, got %d expect %d", len(b), 500)
	}
	for i := range b {
		if b[i] != 0 {
			t.Fatalf("FAIL! Mid extend did not set everything to 0")
		}
	}
}

func TestExtendLong(t *testing.T) {
	b := make(Bites, 50).Set(1)
	origcap := cap(b)
	b = b.Reuse().Extend(5000, true)
	if cap(b) == origcap {
		t.Fatalf("FAIL! Long extend did not change cap")
	}
	if len(b) != 5000 {
		t.Fatalf("FAIL! Long extend did not set len properly, got %d expect %d", len(b), 5000)
	}
	for i := range b {
		if b[i] != 0 {
			t.Fatalf("FAIL! Long extend did not set everything to 0")
		}
	}
}

func TestSlicing(t *testing.T) {
	b := Empty().PutString("0123456789")
	if b.Last(3).String() != "789" {
		t.Fatalf("FAIL! Last 3 bytes should be 789, were %s", b.Last(3))
	}
	if b.Snip(4).String() != "012345" {
		t.Fatalf("FAIL! Snipping the last 4 bytes should result in 012345, but became %s", b.Snip(4))
	}
	if b.Skip(6).String() != "6789" {
		t.Fatalf("FAIL! Skipping the first 6 bytes should result in 6789, but became %s", b.Skip(6))
	}
}

func TestSimplePutString(t *testing.T) {
	b := Empty().PutString("Hello").PutByte(',').PutRune(' ', nil).PutSlice([]byte("world!"))
	if b.String() != "Hello, world!" {
		t.Fatalf("FAIL! Expected 'Hello, world!', got '%s'", b.String())
	}
	if !b.Sequal("Hello, world!") {
		t.Fatalf("FAIL! Not equal to 'Hello, world!', got '%s'", b.String())
	}
}

func TestRune(t *testing.T) {
	b := Empty().PutString("世界!")
	var r1, r2, r3 rune
	var s1, s3 int
	b.GetRune(&r1, &s1).GetRune(&r2, nil).GetRune(&r3, &s3)
	if r1 != '世' {
		t.Fatalf("First rune mismatch: expected %d, got %d", '世', r1)
	}
	if r2 != '界' {
		t.Fatalf("Second rune mismatch: expected %d, got %d", '界', r2)
	}
	if r3 != '!' {
		t.Fatalf("Third rune mismatch: expected %d, got %d", '!', r3)
	}
	if s1 != 3 {
		t.Fatalf("Invalid rune size: %d expected %d", s1, 3)
	}
	if s3 != 1 {
		t.Fatalf("Invalid rune size: %d expected %d", s3, 1)
	}
}

func TestPanicOnBadRune(t *testing.T) {
	err := func() (err *ErrorInvalidRune) {
		defer func() {
			p := recover()
			if p == nil {
				err = nil
				return
			}
			if ir, ok := p.(ErrorInvalidRune); ok {
				err = &ir
				return
			}
			err = nil
			return
		}()
		Empty().PutRune(2000000000, nil)
		return nil
	}()
	if err == nil {
		t.Fatalf("FAIL! Expected ErrorInvalidRune")
	}
	if *err != 2000000000 {
		t.Fatalf("FAIL! Expected rune 2000000000, got %d", err)
	}
	if err.Error() != "Invalid rune: 2000000000" {
		t.Fatalf("FAIL! Invalid error message: '%s'", err.Error())
	}
}

func TestCloneZero(t *testing.T) {
	b1 := Empty().PutString("Hello")
	b2 := b1.Clone()
	b1 = b1.Reuse().PutString("moo")
	if b2.String() != "Hello" {
		t.Fatalf("FAIL! Write after clone should not overwrite")
	}

	b2.Zero()
	if b1.String() != "moo" {
		t.Fatalf("FAIL! Write after clone should not overwrite (%s)", b1.String())
	}

	for i := range b2 {
		if b2[i] != 0 {
			t.Fatalf("FAIL! Zero did not zero slice")
		}
	}
}

func TestSplit(t *testing.T) {
	b1, b2 := Empty().PutString("123987654").Split(3)
	if b1.String() != "123" {
		t.Fatalf("FAIL! First part of split is wrong, expected '123' got %d", b1.String())
	}
	if b2.String() != "987654" {
		t.Fatalf("FAIL! Second part of split is wrong, expected '987654' got %d", b2.String())
	}
}

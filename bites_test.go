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
	b = b.Extend(16)
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
	b = b.Reuse().Extend(20)
	if cap(b) != origcap {
		t.Fatalf("FAIL! Short extend changed cap")
	}
	if len(b) != 20 {
		t.Fatalf("FAIL! Short extend did not set len properly, got %d expect %d", len(b), 20)
	}
}

func TestExtendMid(t *testing.T) {
	b := make(Bites, 50).Set(1)
	origcap := cap(b)
	b = b.Reuse().Extend(500)
	if cap(b) == origcap {
		t.Fatalf("FAIL! Mid extend did not change cap")
	}
	if len(b) != 500 {
		t.Fatalf("FAIL! Mid extend did not set len properly, got %d expect %d", len(b), 500)
	}
}

func TestExtendLong(t *testing.T) {
	b := make(Bites, 50).Set(1)
	origcap := cap(b)
	b = b.Reuse().Extend(5000)
	if cap(b) == origcap {
		t.Fatalf("FAIL! Long extend did not change cap")
	}
	if len(b) != 5000 {
		t.Fatalf("FAIL! Long extend did not set len properly, got %d expect %d", len(b), 5000)
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

func TestSimpleString(t *testing.T) {
	b := Empty().PutString("Hello").PutByte(',').PutRune(' ', nil).PutSlice([]byte("world!"))
	if b.String() != "Hello, world!" {
		t.Fatalf("FAIL! Expected 'Hello, world!', got '%s'", b.String())
	}
	if !b.Sequal("Hello, world!") {
		t.Fatalf("FAIL! Not equal to 'Hello, world!', got '%s'", b.String())
	}
	var s string
	b.GetString(&s, 5)
	if s != "Hello" {
		t.Fatalf("FAIL! GetString not equal to 'Hello'")
	}
}

func TestRune(t *testing.T) {
	var os1 int
	b := Empty().PutRune('世', &os1).PutString("界!")
	if os1 != 3 {
		t.Fatalf("First rune put mismatch: expected %d, got %d", 3, os1)
	}
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

func TestPanics(t *testing.T) {
	func() {
		var r rune
		var s int
		defer catch(t, ErrInvalidRune)
		Empty().PutByte(0xff).GetRune(&r, &s)
	}()
	func() {
		defer catch(t, ErrorInvalidRune(2000000000))
		Empty().PutRune(2000000000, nil)
	}()
	func() {
		defer catch(t, ErrSliceEOF)
		slice := make([]byte, 20)
		Empty().Extend(10).GetSliceCopy(slice)
	}()
	func() {
		defer catch(t, ErrSliceEOF)
		var str string
		Empty().Extend(10).GetString(&str, 20)
	}()
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

func TestBool(t *testing.T) {
	b := Empty().PutBool(true, false, false, true, true, false, false, true, true, false, true)
	if len(b) != 2 {
		t.Fatalf("FAIL! Expected size 2, got %d", len(b))
	}
	if b[0] != 153 || b[1] != 160 {
		t.Fatalf("FAIL! Expected 153 and 160, got %d and %d", b[0], b[1])
	}
	var b1, b2, b3, b6, b7, b8, b9, b10, b11 bool
	b = b.GetBool(&b1, &b2, &b3, nil, nil, &b6, &b7, &b8, &b9, &b10, &b11)
	if len(b) != 0 {
		t.Fatalf("FAIL! GetBool didn't snip")
	}
	if b1 != true || b2 != false || b3 != false || b6 != false || b7 != false || b8 != true || b9 != true || b10 != false || b11 != true {
		t.Fatalf("FAIL! Bad booleans returned")
	}
}

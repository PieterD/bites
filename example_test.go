package bites_test

import (
	"fmt"

	"github.com/PieterD/bites"
)

func ExampleStringOps() {
	// Chain operations to construct your buffer.
	b := bites.New().PutByte('H').PutRune('e', nil).PutString("llo ").PutSlice([]byte("world!"))
	fmt.Println(b.String())
	// Output: Hello world!
}

func ExampleExpect() {
	// Chain operations to construct your buffer.
	var hash [32]byte
	b := bites.New().PutString("some string").PutUint32(513).PutVar(912345)
	b = b.PutByte('Q').PutSlice(hash[:]).PutString("another one")

	// Chain operations to read your buffer, using expect and get.
	var qbyte byte
	var hashcopy [32]byte
	var str string
	r := b.Skip(11).ExpectUint32(513).ExpectVarInt(912345, nil).GetByte(&qbyte)
	r = r.GetSliceCopy(hashcopy[:]).GetString(&str, 7).ExpectString(" one")
	fmt.Printf("%c %s\n", qbyte, str)
	// Output: Q another
}

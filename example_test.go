package bites_test

import (
	"fmt"

	"github.com/PieterD/bites"
)

func ExampleBites() {
	// Chain operations to construct your buffer.
	var hash [32]byte
	b := bites.New().PutString("some string").PutUint32(513).PutVar(912345)
	b = b.PutByte('Q').PutSlice(hash[:]).PutString("another one")

	// Chain operations to read your buffer, using expect and get.
	var qbyte byte
	var hashcopy [32]byte
	var str string
	r := b.Get().Skip(11).ExpectUint32(513).ExpectVarInt(912345, nil).GetByte(&qbyte)
	r = r.GetSliceCopy(hashcopy[:]).GetString(&str, 7).ExpectString(" one")

	fmt.Printf("%c %s\n", qbyte, str)
	// Output: Q another
}

func ExampleWriter() {
	b := bites.New().PutString("Hello")
	w := b.NewWriter()
	fmt.Fprintf(w, ", world!")
	fmt.Println(w.Bites().String())
	// Output: Hello, world!
}

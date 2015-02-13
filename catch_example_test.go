package bites_test

import (
	"fmt"

	"github.com/PieterD/bites"
)

func ExampleCatch() {
	err := PutSomething()
	fmt.Println(err.Error())
	// Output: Unexpectedly reached the end of the bite slice
}

func PutSomething() (err error) {
	defer bites.Catch(&err)
	bites.New().ExpectString("something")
	return nil
}

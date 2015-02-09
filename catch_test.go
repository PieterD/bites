package bites

import (
	"reflect"
	"strings"
	"testing"
)

func catch(t *testing.T, expect error) {
	e := recover()
	if e == nil {
		t.Fatalf("Expected panic: %#v", expect)
		return
	}
	got, ok := e.(error)
	if !ok {
		panic(e)
	}

	if reflect.TypeOf(got) != reflect.TypeOf(expect) {
		t.Fatalf("Expected panic %#v, got %#v", expect, got)
	}
	if !reflect.DeepEqual(got, expect) {
		t.Fatalf("Expected panic %#v, got %#v", expect, got)
	}
	if expect.Error() != got.Error() {
		t.Fatalf("Expected string %s, got %s", expect.Error(), got.Error())
	}
	if strings.ContainsAny(got.Error(), "%!") {
		t.Fatalf("Error string contains %% or !: %s", got.Error())
	}
}

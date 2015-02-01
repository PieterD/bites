// Package bites deals with byte slices.
// Its purpose is to make it easier to marshal and unmarshal the various basic types to and from byte slices.
//
// Most of these methods do not allocate, and many of them are inlined.
// When a method does allocate, it is noted in the docs.
package bites

// Package cast provides save casting functions for converting between types without panicking.
package cast

import (
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
)

// Array65 casts a slice to an array of length 65.
func Array65[A any](slice []A) ([65]A, error) {
	if len(slice) == 65 {
		return [65]A(slice), nil
	}

	return [65]A{}, errors.New("slice length not 65", "len", len(slice))
}

// Must32 casts a slice to an array of length 32.
func Must32[A any](slice []A) [32]A {
	arr, err := Array32(slice)
	if err != nil {
		panic(err)
	}

	return arr
}

// Must20 casts a slice to an array of length 20.
func Must20[A any](slice []A) [20]A {
	arr, err := Array20(slice)
	if err != nil {
		panic(err)
	}

	return arr
}

// EthHash casts a byte slice to an Ethereum hash.
func EthHash(b []byte) (common.Hash, error) {
	resp, err := Array32(b)
	if err != nil {
		return common.Hash{}, errors.New("invalid hash length", "len", len(b))
	}

	return resp, nil
}

// Array32 casts a slice to an array of length 32.
func Array32[A any](slice []A) ([32]A, error) {
	if len(slice) == 32 {
		return [32]A(slice), nil
	}

	return [32]A{}, errors.New("slice length not 32", "len", len(slice))
}

// EthAddress casts a byte slice to an Ethereum address.
func EthAddress(b []byte) (common.Address, error) {
	resp, err := Array20(b)
	if err != nil {
		return common.Address{}, errors.New("invalid address length", "len", len(b))
	}

	return resp, nil
}

// Array20 casts a slice to an array of length 32.
func Array20[A any](slice []A) ([20]A, error) {
	if len(slice) == 20 {
		return [20]A(slice), nil
	}

	return [20]A{}, errors.New("slice length not 20", "len", len(slice))
}

// Array8 casts a slice to an array of length 2.
func Array8[A any](slice []A) ([8]A, error) {
	if len(slice) == 8 {
		return [8]A(slice), nil
	}

	return [8]A{}, errors.New("slice length not 8", "len", len(slice))
}
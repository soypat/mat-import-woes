//go:build matty
// +build matty

package main

import "gonum.org/v1/gonum/mat"

func (m *M) T() mat.Matrix {
	return mat.Transpose{m}
}

package belajar_golang_generics

import (
	"fmt"
	"testing"
)

type Data[T any] struct {
	First  T
	Second T
}

func TestData(t *testing.T) {
	data := Data[string]{
		First: "Ucok",
		Second: "Baba",
	}
	fmt.Println(data)
}

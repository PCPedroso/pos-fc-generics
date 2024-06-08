package main

import (
	"fmt"
	"generics/calculos"
)

type Meu int

type Number interface {
	~int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, val := range m {
		soma += val
	}
	return soma
}

func Soma2[T Number](a, b T) T {
	return a + b
}

func main() {
	m := map[string]Meu{"Valr 1": 100, "Val2": 200}
	m2 := map[string]int{"Valr 1": 100, "Val2": 200}
	fmt.Println(Soma(m))
	fmt.Println(Soma(m2))

	fmt.Println(calculos.Adicao(1, 2))
}

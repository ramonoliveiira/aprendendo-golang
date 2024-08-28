package main

import "fmt"

// Declarando variáveis
func sum(n1 int, n2 int) int {
	return n1 + n2
}

func sum2(n1 int32, n2 int32) (total int32) {
	total = n1 + n2
	return
}

func main() {
	// Declaração de variáveis
	hello := "Hello World!"
	var (
		s  string
		n1 int32
		n2 int32
		f  float64
	)
	s = "Ramon"
	n1 = 10
	n2 = 20
	f = 33.33

	// Cast do tipo int32 para int
	sumResult := sum(int(n1), int(n2))
	sumResult2 := sum2(n1, int32(f))

	fmt.Println("Nome:", s)
	fmt.Println("Soma:", sumResult)
	fmt.Println("Float Number:", sumResult2)
	fmt.Printf("Teste %s", hello)
}

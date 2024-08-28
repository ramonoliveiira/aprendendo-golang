package main

import "fmt"

func simpleLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func simpleLoop2(nomes []string) {
	for i := 0; i < len(nomes); i++ {
		fmt.Println(nomes[i])
	}
}

func rangeLoop(nomes []string) {
	for k, nome := range nomes {
		fmt.Println(k, nome)
	}

	// for _, nome := range nomes {
	// 	fmt.Println(nome)
	// }
}

func continuedLoop(nomes []string) {
	var i int
	for i < len(nomes) {
		fmt.Println(nomes[i])
		i++
	}
}

// Útil caso precise que o código seja executado de forma ininterrupta, como aguardando alguma ação ou ouvindo uma nova conexão.
// É equivalente ao while (true) {}
func infiniteLoop(nomes []string) {
	var i int
	for {
		fmt.Println(nomes[i])
		i++
	}
}

func main() {
	simpleLoop()
	nomes := []string{"Tiago", "Daniel", "Maria", "Ramon"}
	simpleLoop2(nomes)
	rangeLoop(nomes)
	continuedLoop(nomes)
}

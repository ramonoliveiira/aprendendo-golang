package main

import "fmt"

func slices() {
	// slice = São dinâmicos, eles aumentam a capacidade conforme vai recebendo valores
	nomes := []string{"Tiago", "Dani", "Marcos", "Maria"}
	fmt.Println(nomes, len(nomes), cap(nomes))

	// Quando aumentamos um slice em GO, caso o tamanho e a capacidade do array esteja no limite
	// O GO cria um novo array com capacidade maior, realocando o array. Para não ter que fazer isso sempre ao adicionar um novo valor,
	// ele cria um array com uma capacidade maior para evitar muitas realocações.
	nomes = append(nomes, "Rafael")
	fmt.Println(nomes, len(nomes), cap(nomes))

	// Slices iniciados sem valor inicial
	// Inicia o slice com tamanho 10 e capacidade 10
	// nomes2 := make([]string, 10)

	// Aloca um slice de tamanho 10 e capacidade 20
	// nomes3 := make([]string, 10, 20)
}

func maps() {
	idades := make(map[string]uint8)
	idades["Tiago"] = 31
	idades["Dani"] = 36
	idades["Maria"] = 23

	// Map não tem ordenação. Ao ser chamado, ele não exibe os valores na mesma sequência em que foi salvo
	fmt.Println(idades)

	// Validando indices
	// Caso não encontra o índice, retorna o "valor zero" daquele tipo de dado
	fmt.Println(idades["Ramon"])

	val, ok := idades["Tiago"]
	fmt.Println(val, ok)

	val2, ok2 := idades["Ramon"]
	fmt.Println(val2, ok2)

	// Importante saber!
	// Não é possível passar o endereço de memória do valor do índice
	// val3 := &idades["Ramon"]

}

// Atributos públicos e privados
// Público: São atributos ou funções que começam com uma letra maiúscula
// Privado: São atributos ou funções que começam com uma letra minúscula

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     uint8
	Status    bool
	cpf       string // Atributo privado
}

// Structs não podem se referenciar diretamente, mas podem se referenciar indiretamente por ponteiro
type Categoria struct {
	Nome string
	Pai  *Categoria // Ponteiro de Categoria
}

func structs() {
	// var p Pessoa
	// p.Nome = "Ramon"
	// p.Sobrenome = "Oliveira"

	// Atribui os valores conforme a ordem de aparecimento
	// Não é recomendável, pois irá quebrar caso alguém adicione um novo atributo em qualquer ordem
	// p1 := Pessoa{"Ramon", "Oliveira", 26, true}

	// Forma mais recomendada
	p := Pessoa{
		Nome:      "Ramon",
		Sobrenome: "Oliveira",
		Idade:     26,
		Status:    true,
		cpf:       "000.000.000-00",
	}
	fmt.Println(p)
	// Obs: Atributos privados não são acessíveis por packages diferentes do package que detém a struct.
	// Exceto caso seja o mesmo package detentor da struct, como é o caso desse exemplo com "p.cpf".
	fmt.Println(p.Nome, p.Idade, p.cpf)
}

func main() {
	slices()
	maps()
	structs()
}

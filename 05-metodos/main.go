package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     uint8
	Status    bool
	cpf       string // Atributo privado
}

// Método "Mágico" - Parece que esse não é o nome correto
func (p Pessoa) String() string {
	// Formata uma string e a retorna
	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos", p.Nome, p.Idade)
}

type Categoria struct {
	Nome string
	Pai  *Categoria
}

// Geralmente a declaração do nome da struct em um método tem apenas uma letra, sendo ela a primeira letra da struct
func (c Categoria) HasParent() bool {
	return c.Pai != nil
}

// Passando struct como ponteiro
// Quando utiliza o ponteiro a struct vai como referência para o método e é possível alterar os valores
// Caso não utilize ponteiro, ele vai como valor e só é possível ler/recuperar os valores
func (c *Categoria) SetPai(pai *Categoria) {
	c.Pai = pai
}

func main() {
	p := Pessoa{Nome: "Ramon", Sobrenome: "Oliveira", Idade: 26, Status: true, cpf: "000.000.000-00"}
	p.cpf = "1"
	fmt.Println(p)

	cat := Categoria{Nome: "Categoria 1"}
	cat.SetPai(&Categoria{Nome: "Pai"})
	if !cat.HasParent() {
		fmt.Println("Não tem pai")
	}

}

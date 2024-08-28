package main

import "fmt"

type Pessoa struct {
	Nome   string
	Idade  uint8
	Status bool
}

// O método abaixo será herdado pela struct PessoaFisica
func (p Pessoa) String() string {
	return fmt.Sprintf("[Pessoa] Olá, meu nome é %s e eu tenho %d anos", p.Nome, p.Idade)
}

type PessoaFisica struct {
	Pessoa // Herança em GO = Embedding
	// Nome      string // Exemplo campo duplicado na struct PessoaFisica
	Sobrenome string
	cpf       string
}

// func (p PessoaFisica) String() string {
// 	// É possível informar de onde está vindo a informação, usando p.Pessoa.Nome e p.Pessoa.Idade
// 	// return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos", p.Pessoa.Nome, p.Pessoa.Idade)

// 	// Caso o atributo Nome exista na struct Pessoa e também em PessoaFisica, o valor da struct pai será utilizado (tem prioridade).
// 	return fmt.Sprintf("[Pessoa Física] Olá, meu nome é %s e eu tenho %d anos", p.Nome, p.Idade)
// }

type PessoaJuridica struct {
	RazaoSocial string
	cnpj        string
}

func main() {
	p := PessoaFisica{
		Pessoa{Nome: "Ramon", Idade: 26, Status: true},
		// "",
		"Oliveira",
		"000.000.000-00",
	}
	fmt.Println(p)
}

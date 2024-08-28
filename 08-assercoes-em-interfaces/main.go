package main

import "fmt"

type Document interface {
	Doc() string
}

type Pessoa struct {
	Nome   string
	Idade  uint8
	Status bool
}

func (p Pessoa) String() string {
	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos", p.Nome, p.Idade)
}

type PessoaFisica struct {
	Pessoa
	Sobrenome string
	cpf       string
}

func (pf PessoaFisica) Doc() string {
	return fmt.Sprintf("meu cpf é: %s", pf.cpf)
}

type PessoaJuridica struct {
	Pessoa
	RazaoSocial string
	cnpj        string
}

func (pj PessoaJuridica) Doc() string {
	return fmt.Sprintf("meu cnpj é: %s", pj.cnpj)
}

func show(d Document) {
	// type assertion
	// O "d" abaixo é válido apenas no escopo do if e é do tipo PessoaFisica
	// O nome "ok" é uma convenção do Golang ao

	// Essa forma funciona muito bem para um asserção de apenas 1 tipo
	// if d, ok := d.(PessoaFisica); ok {
	// 	fmt.Println(d.Sobrenome)
	// } else if d, ok := d.(PessoaJuridica); ok {
	// 	// Não funciona no else, pois estou tentando fazer asserção novamente no d... eu precisaria ter mais de uma variável T.T
	// 	fmt.Println(d.RazaoSocial)
	// }

	// Valida o tipo do valor d
	// Porém não faz a asserção do tipo, logo a "inferência" deve ser feita dentro do case
	// switch d.(type) {
	// case PessoaFisica:
	// 	fmt.Println(d.(PessoaFisica).Sobrenome)
	// case PessoaJuridica:
	// 	fmt.Println(d.(PessoaJuridica).RazaoSocial)
	// default:
	// 	fmt.Println("tipo desconhecido")
	// }

	// Versão aprimorada
	switch t := d.(type) {
	case PessoaFisica:
		fmt.Println(t.Sobrenome)
	case PessoaJuridica:
		fmt.Println(t.RazaoSocial)
	default:
		fmt.Println("tipo desconhecido")
	}

	fmt.Println(d)
	fmt.Println(d.Doc())
}

func main() {
	pf := PessoaFisica{
		Pessoa{Nome: "Ramon", Idade: 26, Status: true},
		// "",
		"Oliveira",
		"000.000.000-00",
	}
	show(pf)
	fmt.Println()

	pj := PessoaJuridica{
		Pessoa{Nome: "Aprenda Golang", Idade: 0, Status: true},
		"Temporin LTDA",
		"00.000.000/0000-00",
	}
	show(pj)
}

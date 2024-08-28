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
	return fmt.Sprintf("[Pessoa] Olá, meu nome é %s e eu tenho %d anos", p.Nome, p.Idade)
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

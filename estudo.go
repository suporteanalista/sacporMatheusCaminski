package main

import "fmt"

// PRATICA DE USO DE ESTRUTURA E HERANÇA

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint
	tamanho   float32
}

type estudante struct {
	pessoa
	curso        string
	universidade string
}

// AQUI EONDE ESTÃO LOCALIADAS AS VARIAVEIS

func main() {
	fmt.Println("Ola mundo!")

	fmt.Println("-----------------------------------")

	p1 := pessoa{"Matheus", "Caminski", 21, 1.85}
	fmt.Println("Nome: ", (p1.nome), (p1.sobrenome), "Idade:", (p1.idade), "Tamanho:", (p1.tamanho))

	fmt.Println("-----------------------------------")

	p2 := pessoa{"Michael", "Jhoni", 25, 1.78}
	fmt.Println("Nome: ", (p2.nome), (p2.sobrenome), "Idade:", (p2.idade), "Tamanho:", (p2.tamanho))

	fmt.Println("-----------------------------------")

	e1 := estudante{p1, "Sistemas de informação", "UNEMAT"}
	fmt.Println("Nome do aluno: ", p1.nome, "Curso: ", e1.curso, "Qual univercidade: ", e1.universidade)

	fmt.Println("-----------------------------------")

	e2 := estudante{p2, "Radiologia", "FASIP"}
	fmt.Println("Nome do aluno: ", p2.nome, "Curso: ", e2.curso, "Qual univercidade: ", e2.universidade)

	fmt.Println("-----------------------------------")

	// AQUI UTILIZANDO AS VARIAVEIS ACIMA PARA PRATICAR O IF E ELSE

	if p1.idade > p2.idade {
		fmt.Println(p1.nome, "é mais velho que ", p2.nome)
	} else {
		fmt.Println(p2.nome, "é mais velho")
	}

	fmt.Println("-----------------------------------")

	if p1.tamanho > p2.tamanho {
		fmt.Println(p1.nome, "é maior que", p2.nome)
	} else {
		fmt.Println(p2.nome, "é maior")
	}

}

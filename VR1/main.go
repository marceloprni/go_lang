package main

import "fmt"

type Apresentavel interface {
	Apresentar() string
}

type Pessoa struct {
	Nome string
}

func (p Pessoa) Apresentar() string {
	return "Oi, eu sou " + p.Nome
}

func Mostrar(a Apresentavel) {
	fmt.Println(a.Apresentar())
}

func main() {
	p := Pessoa{Nome: "Marcel"}
	Mostrar(p) // Pessoa implementa Apresentavel sem declarar explicitamente
}

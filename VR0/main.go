package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/meet"
)

type Cliente struct {
	Nome     string
	Idade    int
	Endereco Endereco
	Email    string
}

type Endereco struct {
	Rua    string
	Numero int
	Cep    string
	Estado string
}

type Pessoa struct {
	Nome  string
	Idade int
}

func (p *Pessoa) Apresentar() {
	p.Nome = "Joana"
	fmt.Printf("Olá, meu nome é %s e tenho %d anos. \n", p.Nome, p.Idade)
}

func main() {
	var idade int = 25
	var contador int32 = 5
	var indice int64 = 1012312312

	var floarNumber float32 = 10.99
	var doubleNumber float64 = 100.99999

	var booleanVar bool = true

	var stringVar string = "Hello, VR0!"
	var stringVarInferred string = "Inferred String"
	var stringTogether string = stringVar + " " + stringVarInferred

	var gavetas [2]string

	gavetas[0] = "Gaveta 1"
	gavetas[1] = "Gaveta 2"

	var gavetasSlice []string
	gavetasSlice = append(gavetasSlice, "Slice Gaveta 1", "Gaveta Extra")
	gavetasSlice = append(gavetasSlice, "Slice Gaveta 2")

	fmt.Println("Slice Gavetas:", gavetasSlice[1:])

	fmt.Println("Slice Gavetas:", gavetasSlice[0])
	fmt.Println("Slice Gavetas:", gavetasSlice[1])
	fmt.Println("Slice Gavetas:", gavetasSlice[2])

	fmt.Println("Array Gavetas:", gavetas[0])
	fmt.Println("Array Gavetas:", gavetas[1])

	fmt.Println("String Together:", strings.ToUpper(stringTogether))
	fmt.Println(strings.Contains(stringVar, "VR0"))

	fmt.Println("Inferred String Value:", stringVarInferred)

	fmt.Println("String Value:", stringVar)

	fmt.Println("Boolean Value:", booleanVar)

	fmt.Println("Idade:", idade)
	fmt.Println("Contador:", contador)
	fmt.Println("Indice:", indice)

	fmt.Println("Float Number:", floarNumber)
	fmt.Println("Double Number:", doubleNumber)

	meet.Meet()
	meet.SayHello("Hello from VR0 Meet!")

	var pessoas = map[string]int{}

	pessoas["Marcel"] = 25
	pessoas["Ana"] = 30
	fmt.Println("Pessoas Map:", pessoas["Marcel"])
	fmt.Println("Pessoas Map:", pessoas["Ana"])

	// ok gera um booleano
	if idade, ok := pessoas["Marcel"]; ok {
		fmt.Println("Idade de Marcel no if:", idade)
	} else {
		fmt.Println("Marcel não encontrado")
	}

	delete(pessoas, "Ana")
	fmt.Println("Pessoas Map after delete:", pessoas)

	var nota int = 7

	if nota >= 7 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}

	if err := thisIsAnError(); err != nil {
		fmt.Println("Erro encontrado:", err)
	}

	player := map[string]int{
		"score": 100,
	}

	if value, ok := player["score"]; ok {
		fmt.Println("Level found:", value)
	}

	switch nota {
	case 10:
		fmt.Println("Excelente")
	case 7, 8, 9:
		fmt.Println("Bom")
	case 5, 6:
		fmt.Println("Regular")
	default:
		fmt.Println("Reprovado")
	}

	sum := 0
	for i := 1; i <= 10; i++ {
		fmt.Println("Iteration:", i)
		sum += i
	}
	//fmt.Println("Sum from 1 to 10:", sum)

	sum1 := 0
	for sum1 < 20 {
		fmt.Println("Sum in for while style:", sum1)
		sum1 += 2
	}

	nums := []int{1, 2, 3, 4, 5}
	for index, value := range nums {
		fmt.Println("Index:", index, "Value:", value)
	}

	users := map[string]string{
		"nome":      "Marcel",
		"sobreNome": "Silva",
	}
	for index, value := range users {
		fmt.Println("Index:", index, "Value:", value)
	}

	// len(nums) - identificar o tamanho do array
	for b := 0; b < len(nums); b++ {
		fmt.Println(len(nums))
		fmt.Println("Index with traditional for:", b, "Value:", nums[b])
	}

	cliente1 := Cliente{
		Nome:  "lais",
		Idade: 26,
		Endereco: Endereco{
			Rua:    "Rua das Flores",
			Numero: 123,
			Estado: "SP",
		},
	}

	cliente2 := Cliente{
		Nome:  "juan",
		Idade: 39,
	}

	cliente2.Email = "juan@google.com"

	fmt.Println(cliente1.Endereco.Numero)
	cliente1.Endereco.Numero = 124
	fmt.Println(cliente1.Endereco.Numero)

	fmt.Println(cliente2)

	var resultado int = Soma(10, 20)
	fmt.Println("Resultado da Soma:", resultado)

	var fixo = 5
	multiplica := func(x int) int {
		return x * fixo
	}

	var resultado2 int = multiplica(5)
	fmt.Println("Resultado da Multiplicação:", resultado2)

	//p1 := Pessoa{Nome: "Lais", Idade: 26}
	//p1.Apresentar()
	//fmt.Println(p1.Nome)

	var p1 Pessoa = Pessoa{Nome: "lais"}
	var p3 *Pessoa = &p1

	p3.Nome = "Vanessa"

	fmt.Println(p1)
	fmt.Println(p3)

}

// a saida e inteiro
func Soma(a, b int) int {
	return a + b
}

func thisIsAnError() error {
	return errors.New("This is an error")
}

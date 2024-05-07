package main

import (
	"fmt"
	"math"
)

func main() {

	//IMPORTANTE,métodos ou variaveis (PRIVADOS OU PUBLICAS) É DISTINGUIDO POR Minusculo ou Maiusculo, por exemplo

	//Soma = PUBLICA
	//soma = PRIVADA

	// das variaveis

	nome := "Murilo"
	testenumber := 10
	testefloat := 3.15
	testebool := true
	fmt.Println(nome)
	fmt.Printf("%v \n", testenumber)
	fmt.Printf("%v \n", testefloat)
	fmt.Printf("%v \n", testebool)

	//Funções exportadas
	resultado := math.Soma(1, 2)
	fmt.Printf("%v \n", resultado)
	resultadoX := math.SomaX(20)
	fmt.Printf("%v", resultadoX)
}

package main

import "fmt"

//Método em Go (é baseado em estrutura)
type Carro struct {
	Nome string
}

func (c Carro) andar() {
	fmt.Println(c.Nome, "andou")
}

func main() {

	//Para acessar a função do método "Carro"
	carro := Carro{
		Nome: "BMW",
	}

	carro.andar()

	//Variavel que pode substituir uma função fora do escopo da main
	result := func(x ...int) func() int {
		res := 0

		//Loop para somar todos os valores de x
		for _, v := range x {
			res += v
		}
		return func() int {
			return res * res
		}
	}
	//

	resultado1 := soma(6, 2)
	resultado2 := somaTudo(10, 20, 3, 4, 5, 6, 1, 2)

	fmt.Println(resultado1)
	fmt.Println(resultado2)
	/*Para esse método de função tem q se passar um segundo parametro pois o primeiro faz a função de calculo e o segundo retorna o valor desse calculo	*/
	fmt.Println(result(10, 20, 30)())
}

//Função simples
func soma(a int, b int) (result int) {
	result = a + b

	return
}

//Função para passar mais de 1 valor, como se fosse um array
func somaTudo(x ...int) int {
	resultado := 0

	//Loop para somar todos os valores de x
	for _, v := range x {
		resultado += v
	}
	return resultado
}

package main

import "fmt"

//Sintaxe de um struct em Go
type Carro struct {
	Name string
}

//Como acessar e alterar valor de um struct
func (c *Carro) andou() {
	c.Name = "BMW"
	fmt.Println(c.Name, "andou")
}

func main() {
	variavel := 10
	//Aqui o valor da 'variavel' vai ser 200 pois a função abc o parametro esta apontando o valor de 200 na memoria '*a = 200'
	abc(&variavel)
	fmt.Println(variavel)

	carro := Carro{
		Name: "Celtinha",
	}

	carro.andou()
}

func abc(a *int) {
	*a = 200
}

//Função somente para rascunhar anotação
func Rascunho(x int) {
	//'a' tem o valor de 10, onde essa variavel com esse valor é apontado em um endereço da memoria
	//memoria-0xc00000a098(10)   <---- a < ----- 10

	a := 10

	var ponteiro *int = &a

	//O asterisco reconhece(acessa) o valor da memória
	fmt.Println(*ponteiro)

	//Forma de atribuição direto no ponteiro da memória
	*ponteiro = 50
	fmt.Println(*ponteiro)
}

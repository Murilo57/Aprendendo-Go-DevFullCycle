package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cliente struct {
	Nome  string
	Email string
	CPF   int
}

func (c Cliente) Exibe() {
	fmt.Println("Exibindo cliente pelo nome", c.Nome)
}

type ClienteInternacional struct {
	Cliente        //Compondo uma struct com outra, funciona como uma herança
	Pais    string `json:"pais"` //tag para tratamento em caso de JSON, mudar o nome do método
}

func main() {

	cliente := Cliente{
		Nome:  "Murilo",
		Email: "murilo.tt@email.com",
		CPF:   123456789,
	}

	fmt.Println(cliente)

	cliente2 := Cliente{"Marcos", "m@email.com", 987654321}
	fmt.Printf("Nome: %s. Email: %s. CPF: %d \n", cliente2.Nome, cliente2.Email, cliente2.CPF)

	cliente3 := ClienteInternacional{
		Cliente: Cliente{
			Nome:  "Davi",
			Email: "sl@email.com",
			CPF:   43252353,
		},
		Pais: "Inglaterra",
	}

	fmt.Printf("Nome: %s. Email: %s. CPF: %d. Pais: %s \n", cliente3.Nome, cliente3.Email, cliente3.CPF, cliente3.Pais)
	cliente.Exibe()
	cliente2.Exibe()
	cliente3.Exibe()

	//Convertendo uma struct para JSON, a função 'Marshal' da biblioteca json converte uma struct para um JSON
	//Porém esse Marshal converte para um slice em byte
	clienteJson, err := json.Marshal(cliente3)
	if err != nil {
		log.Fatal(err.Error())
	}

	//a função string() converte o JSON de byte para uma string, assim exibindo o JSON corretamente
	fmt.Println(string(clienteJson))

	jsonCliente4 := `{"Nome":"Davi","Email":"sl@email.com","CPF":43252353,"pais":"Inglaterra"}`
	cliente4 := ClienteInternacional{}

	json.Unmarshal([]byte(jsonCliente4), &cliente4)
	fmt.Println(cliente4)
}

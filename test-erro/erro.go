package main

import (
	"fmt"
	"net/http"
)

func main() {

	//Teste de erros, "err" geralmente é a variavel q se caso houver erro na requisição
	res, err := http.Get("http://google.com.br")
	if err != nil {
		//função panic para todo o sistema caso caia nela
		panic(err)
	} else {
		fmt.Println(res.Header)
	}
}

package main

import "fmt"
import "reflect"

func main() {
	nome := "Leo"
	idade := 26
	versao := 1.1
	fmt.Println("Olá,", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("O tipo da variável versao é", reflect.TypeOf(versao))
}
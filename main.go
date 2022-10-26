package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

// Escrever registra uma mensagem na tela
func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("juliana@tempocerto.inf.br")
	fmt.Println(erro)
}

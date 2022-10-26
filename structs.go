package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Juliana"
	u.idade = 34
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua Monaco", 0}

	usuario2 := usuario{"Juliana", 34, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 34}
	fmt.Println(usuario3)

}

package main

import (
	"fmt"
	enderecos "simple-test/src/endereco"
)

func main() {
	rua := enderecos.TipoDeEndereco("avenida Paulista")
	fmt.Println(rua)
}

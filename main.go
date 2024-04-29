package main

import (
	"fmt"
	endereco "simple-test/src"
)

func main() {
	rua := endereco.TipoDeEndereco("avenida Paulista")
	fmt.Println(rua)
}

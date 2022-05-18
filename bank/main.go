package main

import (
	// "estudos_golang/bank/clientes"
	"estudos_golang/bank/contas"
	"fmt"
)

func main() {
	// 	clienteLeonardo := clientes.Titular{"Leonardo", "123.456.789.12", "Desenvolvedor"}
	contaDoLeonardo := contas.ContaCorrente{}

	contaDoLeonardo.Depositar(100)

	fmt.Println(contaDoLeonardo)

}

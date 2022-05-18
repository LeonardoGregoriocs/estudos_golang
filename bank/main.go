package main

import (
	"estudos_golang/bank/contas"
	"fmt"
)

func main() {

	contaLeo := contas.ContaCorrente{}
	fmt.Println(contaLeo)

	contaLeoPoupe := contas.ContaPoupanca2{}
	fmt.Println(contaLeoPoupe.ObterSaldo)

}

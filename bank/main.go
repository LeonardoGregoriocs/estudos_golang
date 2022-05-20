package main

import (
	"estudos_golang/bank/clientes"
	"estudos_golang/bank/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	Leonardo := clientes.Titular{
		Nome:      "Leonardo",
		CPF:       "741.852.963.36",
		Profissao: "Desenvolvedor",
	}

	contaLeo := contas.ContaCorrente{
		Titular:       Leonardo,
		NumeroAgencia: 123,
		NumemoConta:   654321,
	}

	contaLeo.Depositar(100)
	fmt.Println("Saldo contaLeo:", contaLeo.ObterSaldo())

	PagarBoleto(&contaLeo, 60)
	fmt.Println("Saldo contaLeo:", contaLeo.ObterSaldo())

}

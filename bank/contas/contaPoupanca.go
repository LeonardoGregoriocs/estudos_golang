package contas

import (
	"estudos_golang/bank/clientes"
)

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumemoConta   int
	Operacao      int
	saldo         float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente!"
	}
}

func (c *ContaPoupanca) Depositar(valordoDeposito float64) (string, float64) {
	if valordoDeposito > 0 {
		c.saldo += valordoDeposito
		return "Deposito feito com sucesso!", c.saldo
	} else {
		return "Valor do deposito menor que zero!", c.saldo
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}

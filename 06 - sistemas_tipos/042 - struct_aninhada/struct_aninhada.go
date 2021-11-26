package main

import "fmt"

type item struct {
	produtoID  int
	quantidade int
	preco      float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0

	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}

	return total
}

func main() {

	pedido := pedido{
		userID: 1,
		itens: []item{
			// 1 = produtoID | 2 = quantidade | 12.10 = preço
			item{produtoID: 1, quantidade: 2, preco: 12.10},
			// 2 = produtoID | 1 = quantidade | 23.58 = preço
			item{2, 1, 23.58},
			// 11 = produtoID | 110 = quantidade | 0.89 = preço
			item{11, 110, 0.89},
		},
	}

	fmt.Printf("valor total do Pedido é de R$ %.2f", pedido.valorTotal())
}

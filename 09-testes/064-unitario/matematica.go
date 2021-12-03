package matematica

import (
	"fmt"
	"strconv"
)

// Função reponsável por calcular a média entre um grupo de Valores
func Media(numeros ...float64) float64 {
	total := 0.0

	for _, num := range numeros {
		total += num
	}
	media := total / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)

	return mediaArredondada
}
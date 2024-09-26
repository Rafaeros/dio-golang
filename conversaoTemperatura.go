package conversaoTemperatura

import (
	"fmt"
)

func converterEmCelsius(kelvin int64) int64{
	var celsius int64
	celsius = kelvin - 273
	return celsius
}

func main() {
	var pontoEbulicao int64 = 373
	c := converterEmCelsius(pontoEbulicao)
	fmt.Printf("O ponto de ebulição da agua em kelvin é: %d e a conversão para celsius é: %d", pontoEbulicao, c)
}
	
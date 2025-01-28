package main

import (
	"fmt"
	"math"
)

type formas interface {
	calculaArea() float64
}

func mostrarAreaForma(f formas) {
	fmt.Printf("A forma possui area = %0.2f \n", f.calculaArea())
}
type retangulo struct {
	largura int
	altura  int
}
func (r retangulo) calculaArea() float64 {
	return float64(r.altura * r.largura)
}

type circulo struct {
	raio int
}

func (c circulo) calculaArea() float64 {
	return math.Pi * math.Pow(float64(c.raio), 2)
}

func main() {
	fmt.Println("CALCULO RESOLVIDO")
	r := retangulo{60,40}
	mostrarAreaForma(r)

	circle := circulo{23}
	mostrarAreaForma(circle)
}
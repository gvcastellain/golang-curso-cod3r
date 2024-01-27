package main

import "math"

//se comeca com maiuscula eh publico
// minuscula ou _Ponto priado

type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cX, cY float64) {
	cX = math.Abs(b.x - a.x)
	cY = math.Abs(b.y - a.y)
	return
}

func distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}

package main

import (
	"fmt"
	"math"
	"time"
)

/*
Task
A rectangle with sides equal to even integers a and b is drawn on the Cartesian plane. Its center (the intersection point of its diagonals) coincides with the point (0, 0), but the sides of the rectangle are not parallel to the axes; instead, they are forming 45 degree angles with the axes.

How many points with integer coordinates are located inside the given rectangle (including on its sides)?

Example
For a = 6 and b = 4, the output should be 23

The following picture illustrates the example, and the 23 points are marked green.

OBS: Ver imagem blob.png

Input/Output
[input] integer a

A positive even integer.

Constraints: 2 ≤ a ≤ 10000.

[input] integer b

A positive even integer.

Constraints: 2 ≤ b ≤ 10000.

[output] an integer

The number of inner points with integer coordinates.
*/

func main() {
	start := time.Now()
	fmt.Println(RectangleRotation(30000, 20000))
	fmt.Println(time.Since(start).Milliseconds())
	startV2 := time.Now()
	fmt.Println(RectangleRotationV2(30000, 20000))
	fmt.Println(time.Since(startV2).Milliseconds())
	// Note que o tempo de execução da versão 2 é muito menor que a versão 1, mesmo para valores grandes de a e b
}

/*
Parametrização do retângulo, com base na equação da reta, para determinar os pontos internos com coordenadas inteiras.
x
*/

func RectangleRotation(a, b int) int {
	var count int
	var xMin, xMax float64 = -(float64(a) / 2), float64(a) / 2
	var yMin, yMax float64 = -(float64(b) / 2), float64(b) / 2

	var sin45 float64 = math.Sin(math.Pi / 4)
	var cos45 float64 = math.Cos(math.Pi / 4)

	var xRotMin, xRotMax float64 = -(float64(a+b) * cos45 / 2), float64(a+b) * cos45 / 2
	var yRotMin, yRotMax float64 = -(float64(a+b) * sin45 / 2), float64(a+b) * sin45 / 2

	if a <= 0 || b <= 0 {
		return 0
	}
	if a%2 != 0 || b%2 != 0 {
		return 0
	}

	// Delimitação do retângulo rotacionado, com base na equação da reta, para determinar os pontos internos com coordenadas inteiras.
	// y = sqrt(2)/2 * (yRot - xRot)
	// x = sqrt(2)/2 * (yRot + xRot)
	// VERIFICAR IMAGEM COM RASCUNHOS
	for xRot := int(math.Ceil(xRotMin)); xRot <= int(math.Floor(xRotMax)); xRot++ {
		for yRot := int(math.Ceil(yRotMin)); yRot <= int(math.Floor(yRotMax)); yRot++ {
			if (float64(yRot-xRot)*sin45 <= yMax && float64(yRot-xRot)*sin45 >= yMin) && (float64(yRot+xRot)*cos45 <= xMax && float64(yRot+xRot)*cos45 >= xMin) {
				count++
			}
		}
	}

	return count
}

/*
Solução devin
Dá pra sair de O(a·b) para O(1) (só uma raiz quadrada inteira), sem float.

Ideia: girar o ponto em vez do retângulo. (x,y) está dentro do retângulo girado 45° sse

|x+y| ≤ a/√2 e |y−x| ≤ b/√2.

Com u = x+y, v = y−x, o mapa (x,y) ↔ (u,v) é bijetivo exatamente quando u e v têm a mesma paridade. Então basta contar:

U = maior inteiro com 2U² ≤ a², V = maior inteiro com 2V² ≤ b²
pares = (ímpares em [-U,U] × ímpares em [-V,V]) + (pares × pares)
*/
func RectangleRotationV2(a, b int) int {
	if a <= 0 || b <= 0 {
		return 0
	}
	u := maxHalf(a) // max u >= 0 com 2u^2 <= a^2
	v := maxHalf(b)

	oddU, evenU := u-u/2, u/2 // contagens em [1..n]
	oddV, evenV := v-v/2, v/2

	// [-n,n]: ímpares = 2*odd, pares = 2*even + 1 (inclui 0)
	return (2*oddU)*(2*oddV) + (2*evenU+1)*(2*evenV+1)
}

func maxHalf(a int) int {
	n := int(math.Sqrt(float64(a) * float64(a) / 2))
	for 2*(n+1)*(n+1) <= a*a {
		n++
	}
	for n > 0 && 2*n*n > a*a {
		n--
	}
	return n
}

/*
Dois pontos sobre a versão original:

os if a <= 0 || b <= 0 e a%2 != 0 || b%2 != 0 estavam depois dos cálculos (desperdício) e o teste de paridade retorna 0 para lados ímpares — no problema clássico (CodeSignal) lados ímpares são válidos. A versão acima trata qualquer a, b > 0; se você precisa mesmo do comportamento antigo, é só manter o guard de paridade no topo.
comparações com float64 na borda (<= xMax) são sensíveis a erro de ponto flutuante; a versão inteira elimina isso.
*/

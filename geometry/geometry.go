package geometry

/*
	TODO:
	- O projeto para o qual vc quer gerar os testes será sempre um projeto externo
	- Testar com projetos do benchmark do artigo do nxt_unit
	- Verificar a ferramenta nxt_unit
	- Verificar o código de obtenção de função
*/
import (
	"fmt"
	"math"
)

type Point struct {
	X	float64;
	Y   float64;
}

type Line struct {
	m 	float64;
	b 	float64;
}

func Add(a, b int) int {
	return a + b
}

func InitPoint(x, y float64) (*Point) {
	point := &Point{
		x, 
		y,
	}

	return point
}

func (pointA *Point) ManhattanDistance(pointB *Point) float64 {
	return math.Abs(pointA.X - pointB.X) + math.Abs(pointA.Y - pointB.Y)
}

func (pointA *Point) EuclideanDistance(pointB *Point) float64 {
	return math.Sqrt(math.Pow((pointA.X - pointB.X), 2) + math.Pow((pointA.Y - pointA.Y), 2))
}

func PointToLineDistance(point *Point, line *Line) float64 {
	num := math.Abs(line.m * point.X + point.Y + line.b)
	dem := math.Sqrt(math.Pow(line.m, 2) + 1)

	return num / dem
}

func (point *Point) ToString() string {
	return fmt.Sprintf("(%v, %v)", point.X, point.Y)
}

func GetLineFromPoints(pointA *Point, pointB *Point) *Line {
	m := (pointB.Y - pointA.Y) / (pointB.X - pointA.X)
	b := -(m * pointA.X) + pointA.Y

	return &Line{m, b}
}

func (line *Line) ToString() string {
	var sign string
	if line.b > 0 {
		sign = "+"
	} else {
		sign = "-"
	}

	return fmt.Sprintf("y = %.2v %v %.2v", line.m, sign, math.Abs(line.b))
}


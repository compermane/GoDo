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

type All struct {
	A string;
	B bool;
	C int;
	D int32;
	E int64;
	F float32;
	G float64
}

type Point struct {
	X	float64;
	Y   float64;
}

type Line struct {
	m 	float64;
	b 	float64;
}

func Example(all Point) float64 {
	return all.X + all.Y
}

func (all *All) Example() int {
	return 1
}

func (all *All) ExampleWithArgs(a int, b string, c bool) int {
	if c {
		fmt.Println(b)
	}
	return a
}

func Add(a, b int) int {
	return a + b
}

func MinusOne(a int) int {
	return a - 1
}

func SumString(a, b string) string {
	return a + b
}

func BoolFunc(a, b, c bool) bool {
	return a
}

func InitPoint(x, y float64) (*Point) {
	point := &Point{
		x, 
		y,
	}

	return point
}

func AlwaysPanic() {
	panic("This is a panic")
}

func (pointA Point) Bruh() {
	fmt.Println("BRuh")
}
func (pointA *Point) ManhattanDistance(pointB *Point) float64 {
	return math.Abs(pointA.X - pointB.X) + math.Abs(pointA.Y - pointB.Y)
}

func (pointA *Point) EuclideanDistance(pointB *Point, a int) float64 {
	return math.Sqrt(math.Pow((pointA.X - pointB.X), 2) + math.Pow((pointA.Y - pointA.Y), 2))
}

func PointToLineDistance(point *Point, line *Line) float64 {
	num := math.Abs(line.m * point.X + point.Y + line.b)
	dem := math.Sqrt(math.Pow(line.m, 2) + 1)

	return num / dem
}

func ToString(point *Point) string {
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


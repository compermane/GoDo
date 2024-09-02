package coordinates

import (
	"fmt"
	"math"
)

type Point struct {
	x	float64;
	y   float64;
}

type Line struct {
	m 	float64;
	b 	float64;
}

func InitPoint(x, y float64) (*Point) {
	point := &Point{
		x, 
		y,
	}

	return point
}

func (pointA *Point) ManhattanDistance(pointB *Point) float64 {
	return math.Abs(pointA.x - pointB.x) + math.Abs(pointA.y - pointB.y)
}

func (pointA *Point) EuclideanDistance(pointB *Point) float64 {
	return math.Sqrt(math.Pow((pointA.x - pointB.x), 2) + math.Pow((pointA.y - pointA.y), 2))
}

func PointToLineDistance(point *Point, line *Line) float64 {
	num := math.Abs(line.m * point.x + point.y + line.b)
	dem := math.Sqrt(math.Pow(line.m, 2) + 1)

	return num / dem
}

func (point *Point) ToString() string {
	return fmt.Sprintf("(%v, %v)", point.x, point.y)
}

func GetLineFromPoints(pointA, pointB *Point) *Line {
	m := (pointB.y - pointA.y) / (pointB.x - pointA.x)
	b := -(m * pointA.x) + pointA.y

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


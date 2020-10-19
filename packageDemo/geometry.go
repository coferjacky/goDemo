package packageDemo

import "math"

type Point struct {
	X, Y float64
}

//P Q 两点的矢量距离，其中p就是一个方法的接收器
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-q.Y)
}

//坐标放大几倍
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

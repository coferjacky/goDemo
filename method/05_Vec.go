package main

import "math"

//二维矢量结构
type Vec2 struct {
	X, Y float32
}

//两个矢量相加
func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		v.X + other.X,
		v.Y + other.Y,
	}
}

//两个矢量相减
func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{
		v.X - other.X,
		v.Y - other.Y,
	}
}

//两个矢量相乘
func (v Vec2) Scale(s float32) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

//两个矢量距离
func (v Vec2) Distance_To(other Vec2) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

//返回当前矢量的标准化矢量
func (v Vec2) Normalize() Vec2 {
	mag := v.X*v.X + v.Y*v.Y
	if mag > 0 {
		one_Over_Mag := 1 / float32(math.Sqrt(float64(mag)))
		return Vec2{v.X * one_Over_Mag, v.Y * one_Over_Mag}
	}
	return Vec2{0, 0}
}

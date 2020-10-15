package main

import "fmt"

//处理移动的逻辑 入-进阶 6.5
type Player struct {
	curr_Pos, target_Pos Vec2
	speed                float32
}

//设置玩家的移动的目标位置
func (p *Player) Move_To(v Vec2) {
	p.target_Pos = v
}

//获取当前玩家位置
func (p *Player) Pos() Vec2 {
	return p.curr_Pos
}

//判断是否到达目的地
func (p *Player) Is_Arrived() bool {
	return p.curr_Pos.Distance_To(p.target_Pos) < p.speed
}

//更新玩家位置
func (p *Player) Update() {
	if !p.Is_Arrived() { //Is_Arrived=false执行
		//计算当前位置指向目标的朝向
		dir := p.target_Pos.Sub(p.curr_Pos).Normalize()
		//添加速度矢量生成新的位置
		new_Pos := p.curr_Pos.Add(dir.Scale(p.speed))

		//移动完成后更新当前的位置
		p.curr_Pos = new_Pos
	}
}
func New_Player(speed float32) *Player {
	return &Player{

		speed: speed,
	}
}

func main() {
	//实例化玩家
	p := New_Player(0.5)
	//玩家移动到 3,1点
	p.Move_To(Vec2{3, 1})

	//如果没到达就一直循环
	for !p.Is_Arrived() {
		p.Update()
		fmt.Println(p.Pos())
	}
}

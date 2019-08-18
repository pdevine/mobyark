package main

import (
	"math"
	sprite "github.com/pdevine/go-asciisprite"
	vec2d "github.com/pdevine/vector2d"
)

const whale_c1 = `xxxxxxxxxxxxxxxxxxxxxxxxx.xxxxx
xxxxxxxxxxxxxxxxxxxxxxxx==xxxxx
xxxxxxxxxxxxxxxxxxxxxxx===xxxxx
xx/""""""""""""""""\___/x===xxx
x{                      /xx===x
xx\______ o          __/xxxxxxx
xxxx\    \        __/xxxxxxxxxx`

const whale_c1_rev = `xxxxx.xxxxxxxxxxxxxxxxxxxxxxxxx
xxxxx==xxxxxxxxxxxxxxxxxxxxxxxx
xxxxx===xxxxxxxxxxxxxxxxxxxxxxx
xxx===x\___/""""""""""""""""\xx
x===xx\                      }x
xxxxxxx\__          o ______/xx
xxxxxxxxxx\__        /    /xxxx`


type Moby struct {
	sprite.BaseSprite
	Vel   vec2d.Vec2D
        Accel vec2d.Vec2D
}

func NewMoby() *Moby {
	s := &Moby{BaseSprite: sprite.BaseSprite{
		X: 20,
		Y: 40,
		Visible: true,
		},
	}
	s.AddCostume(sprite.NewCostume(whale_c1_rev, 'x'))
	s.AddCostume(sprite.NewCostume(whale_c1, 'x'))
	return s
}

func (s *Moby) Update() {
	s.Vel = s.Vel.Add(s.Accel)
	s.Accel = vec2d.NewVec2D(0, 0)
	s.Vel = s.Vel.Multiply(0.91)
	
	s.X += int(math.Round(s.Vel.X))
}

func (s *Moby) MoveLeft() {
	s.CurrentCostume = 1
	s.Accel = s.Accel.Add(vec2d.NewVec2D(-2, 0))
	s.Vel = vec2d.NewVec2D(0, 0)
}

func (s *Moby) MoveRight() {
	s.CurrentCostume = 0
	s.Accel = s.Accel.Add(vec2d.NewVec2D(2, 0))
	s.Vel = vec2d.NewVec2D(0, 0)
}

type Ball struct {
	sprite.BaseSprite
	Pos vec2d.Vec2D
	Vel vec2d.Vec2D
}

func NewBall(pos, vel vec2d.Vec2D) *Ball {
	s := &Ball{BaseSprite: sprite.BaseSprite{
		X: int(pos.X),
		Y: int(pos.Y),
		Visible: true,
		},
	}
	s.Pos = pos
	s.Vel = vel
	s.AddCostume(sprite.NewCostume("o", ' '))
	return s
}

func (s *Ball) Update() {
	s.Pos = s.Pos.Add(s.Vel)
	s.X = int(s.Pos.X)
	s.Y = int(s.Pos.Y)
	if s.Y <= 0 {
		s.Vel = vec2d.NewVec2D(s.Vel.X, s.Vel.Y * -1)
		s.Pos = vec2d.NewVec2D(s.Pos.X, 1) 
		s.Y = 0
	}
	if s.Y >= 55 {
		s.Dead = true
	}
	if s.Y >= moby.Y+3 && s.Y < moby.Y+9 && s.X >= moby.X+2 && s.X < moby.X+moby.Width {
		s.Vel = vec2d.NewVec2D(s.Vel.X, s.Vel.Y * -1)
		s.Pos = vec2d.NewVec2D(s.Pos.X, float64(moby.Y))
		s.Y = moby.Y
	}
	if s.X <= 0 {
		s.Vel = vec2d.NewVec2D(s.Vel.X * -1, s.Vel.Y)
		s.Pos = vec2d.NewVec2D(1, s.Pos.Y) 
		s.X = 1
	}
	if s.X >= 103 {
		s.Vel = vec2d.NewVec2D(s.Vel.X * -1, s.Vel.Y)
		s.Pos = vec2d.NewVec2D(102, s.Pos.Y) 
		s.X = 102
	}
}

func RemoveBall(idx int, b *Ball) {
	allSprites.Remove(b)
	copy(gs.Balls[idx:], gs.Balls[idx+1:])
	gs.Balls[len(gs.Balls)-1] = nil
	gs.Balls = gs.Balls[:len(gs.Balls)-1]
}

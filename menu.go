package main

import (
	"math"
	sprite "github.com/pdevine/go-asciisprite"
)

const iso_a = `
xxxx___
xxx/\  \
xx/::\  \
x/::\:\__\
x\/\::/  /
xxx/:/  /
xxx\/__/`

const iso_e = `
xxxx___
xxx/\  \
xx/::\  \
x/::\:\__\
x\:\:\/  /
xx\:\/  /
xxx\/__/`

const iso_g = `
xxxx___
xxx/\  \
xx/::\  \
x/:/\:\__\
x\:\:\/__/
xx\::/  /
xxx\/__/` 

const iso_m = `
xxxx___
xxx/\__\
xx/::L_L_
x/:/L:\__\
x\/_/:/  /
xxx/:/  /
xxx\/__/` 

const iso_o = `
xxxx___
xxx/\  \
xx/::\  \
x/:/\:\__\
x\:\/:/  /
xx\::/  /
xxx\/__/`

const iso_r = `
xxxx___
xxx/\  \
xx/::\  \
x/::\:\__\
x\;:::/  /
xx|:\/__/
xxx\|__|`

const iso_v = `
xxxx___
xxx/\__\
xx/:/ _/_
x|::L/\__\
x|::::/  /
xxL;;/__/`

var LetterMap = map[rune]string{
	'a': iso_a,
	'e': iso_e,
	'g': iso_g,
	'm': iso_m,
	'o': iso_o,
	'r': iso_r,
	'v': iso_v,
}

type Letter struct {
        sprite.BaseSprite
	Timer   int
	TimeOut int
}

type GameOver struct {
	Letters []*Letter
}

func NewLetter(letter rune) *Letter {
	l := &Letter{BaseSprite: sprite.BaseSprite{
                X: 20,
                Y: 20,
                Visible: true},
	}
	l.AddCostume(sprite.NewCostume(LetterMap[letter], 'x'))
	return l
}

func (l *Letter) Update() {
	l.Timer++
	l.Y = int(math.Sin(float64(l.Timer/10))*3) + 20
}

func NewGameOver() *GameOver {
	g := &GameOver{}

	for n, c := range []rune{'g', 'a', 'm', 'e', 'o', 'v', 'e', 'r'} {
		l := NewLetter(c)
		l.Timer = n
		if n < 4 {
			l.X = n*10 + 20
		} else {
			l.X = n*10 + 30
		}
		g.Letters = append(g.Letters, l)
		allSprites.Sprites = append(allSprites.Sprites, l)
	}

	return g
}

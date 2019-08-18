package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	sprite "github.com/pdevine/go-asciisprite"
	tm "github.com/pdevine/go-asciisprite/termbox"
	vec2d "github.com/pdevine/vector2d"
)

type GameState struct {
	CurrLevel    *Level
	Balls        []*Ball
	LevelNames   []string
	NextLevelIdx int
}

type Tile struct {
	sprite.BaseSprite
	TileType rune
	Dead     bool
}

type Background struct {
	sprite.BaseSprite
}

type Level struct {
	Tiles []*Tile
}

var ColorMap = map[rune]tm.Attribute{
  'R': tm.ColorRed,
  'b': tm.Attribute(53),
  't': tm.Attribute(180),
  'Y': tm.ColorYellow,
  'N': tm.ColorBlack,
  'B': tm.ColorBlue,
  'o': tm.Attribute(209),
  'O': tm.Attribute(167),
  'w': tm.ColorWhite,
  'g': tm.ColorGreen,
  'G': tm.Attribute(35),
}

func findLevels() []string {
	var levelNames []string
	filepath.Walk("levels", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			levelNames = append(levelNames, filepath.Join("levels", info.Name()))
		}
		return nil
	})
	return levelNames
}
	

func NewGameState() *GameState{
        ball := NewBall(vec2d.NewVec2D(10, 30), vec2d.NewVec2D(-0.5, -0.5))
	allSprites.Sprites = append(allSprites.Sprites, ball)

	gs = &GameState{
		Balls: []*Ball{ball},
	}
	gs.LevelNames = findLevels()
	gs.NextLevel()
	return gs
}

func (g *GameState) NextLevel() {
	g.CurrLevel = LoadLevel(g.LevelNames[g.NextLevelIdx])	
	g.NextLevelIdx++
}

func (g *GameState) Update() {
	g.CurrLevel.Update()
	// XXX - do something fancier here
	if len(g.CurrLevel.Tiles) == 0 {
		g.NextLevel()
	}
}

func NewTile(x, y int, tileType rune) *Tile {
	s := &Tile{BaseSprite: sprite.BaseSprite{
		X:        x,
		Y:        y,
		Visible:  true,
		},
		Dead:     false,
		TileType: tileType,
	}

	var blocks []*sprite.Block
	var b *sprite.Block
	for cnt := 0; cnt < 6; cnt++ {
		if cnt % 2 == 0 {
			b = &sprite.Block{'[', ColorMap[tileType], ColorMap['N'], cnt, 0}
		} else {
			b = &sprite.Block{']', ColorMap[tileType], ColorMap['N'], cnt, 0}
		}
		blocks = append(blocks, b)
	}

	s.AddCostume(sprite.Costume{Blocks: blocks})
	s.Width = 6
	
	return s
}

func (s *Tile) Update() {
	for _, ball := range gs.Balls {
		if ball.Dead == true {
			continue
		}
		if !s.Dead && ball.Y == s.Y && ball.X >= s.X && ball.X <= s.X+s.Width {
			ball.Vel = vec2d.NewVec2D(ball.Vel.X, ball.Vel.Y * -1)
			s.Dead = true
			if s.TileType == 'w' {
				if rand.Intn(10) == 0 {
        				ball := NewBall(vec2d.NewVec2D(float64(s.X+3), float64(s.Y)), vec2d.NewVec2D(-0.5, -0.5))
					allSprites.Sprites = append(allSprites.Sprites, ball)
					gs.Balls = append(gs.Balls, ball)
				}
			}
		}
	}
}

func NewBackground() *Background {
	s := &Background{BaseSprite: sprite.BaseSprite{
		Visible: true,
		},
	}

	var blocks []*sprite.Block
	var b *sprite.Block
	for cnt := 0; cnt < 40; cnt++ {
		b = &sprite.Block{Char: '!', X: 0, Y:cnt}
		blocks = append(blocks, b)
		b = &sprite.Block{Char: '!', X: 103, Y:cnt}
		blocks = append(blocks, b)
	}
	s.AddCostume(sprite.Costume{Blocks: blocks})
	
	return s
}

func LoadLevel(filename string) *Level {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	l := &Level{}
	for y, line := range strings.Split(string(dat), "\n") {
		for x, c := range line {
			if c != ' ' {
				t := NewTile(x*6+1, y, c)
				allSprites.Sprites = append(allSprites.Sprites, t)
				l.Tiles = append(l.Tiles, t)
			}
		}
	}
	return l
}

func (l *Level) Update() {
	// iterate through the tiles in reverse to be able to properly remove them
	for cnt := len(l.Tiles)-1; cnt >= 0; cnt-- {
		if l.Tiles[cnt].Dead {
			l.RemoveTile(cnt, l.Tiles[cnt])
		}
	}
	for cnt := len(gs.Balls)-1; cnt >= 0; cnt-- {
		if gs.Balls[cnt].Dead {
			RemoveBall(cnt, gs.Balls[cnt])
		}
	}
}

func (l *Level) RemoveTile(idx int, t *Tile) {
	allSprites.Remove(t)
	copy(l.Tiles[idx:], l.Tiles[idx+1:])
	l.Tiles[len(l.Tiles)-1] = nil
	l.Tiles = l.Tiles[:len(l.Tiles)-1]
}


package main

import (
	//"flag"
	"fmt"
	"time"
	"math/rand"

	sprite "github.com/pdevine/go-asciisprite"
	tm "github.com/pdevine/go-asciisprite/termbox"
)

var allSprites sprite.SpriteGroup
var Width int
var Height int

const (
	title = iota
	play
	paused
	gameover
)

var gamemode = play
var moby *Moby
var gs *GameState

func main() {
	// XXX - hack to make this work inside of a Docker container
	time.Sleep(1000 * time.Millisecond)

	//debug := flag.Bool("debug", false, "")
	//flag.Parse()

	err := tm.Init()
	if err != nil {
		fmt.Printf("Could not initialize the terminal. Try running as `docker run -it --rm mobyark`.")
		panic(err)
	}
	defer tm.Close()

	rand.Seed(time.Now().UnixNano())
	Width, Height = tm.Size()

	moby = NewMoby()
	bg := NewBackground()

	gs = NewGameState()

	//NewGameOver()

	allSprites.Sprites = append(allSprites.Sprites, bg)
	allSprites.Sprites = append(allSprites.Sprites, moby)

	event_queue := make(chan tm.Event)
	go func() {
		for {
			event_queue <- tm.PollEvent()
		}
	}()

mainloop:
	for {
		start := time.Now()
		tm.Clear(tm.ColorDefault, tm.ColorDefault)

		select {
		case ev := <-event_queue:
			if ev.Type == tm.EventKey {
				if ev.Key == tm.KeyEsc {
					break mainloop
				}
				if gamemode == title {
					//
				} else if ev.Key == tm.KeyEnter {
					//
				} else if ev.Key == tm.KeySpace {
					//
				} else if ev.Key == tm.KeyArrowLeft {
					moby.MoveLeft()
				} else if ev.Key == tm.KeyArrowRight {
					moby.MoveRight()
				} else if ev.Ch == 'p' || ev.Ch == 'P' {
					if gamemode == paused {
						gamemode = play
					} else if gamemode == play {
						gamemode = paused
					}
				}
			} else if ev.Type == tm.EventResize {
				Width = ev.Width
				Height = ev.Height
			}
		default:
			allSprites.Update()
			gs.Update()
			allSprites.Render()
			elapsed := time.Since(start)
			time.Sleep(time.Second/60 - elapsed)
		}
	}
}

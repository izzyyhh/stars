package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Star struct {
	x,y,z float32
	vx, vy float32
	radius float32
}

func (c *Star) Update() {
    c.x += c.vx
    c.y += c.vy

    if c.x-c.radius < 0 || c.x+c.radius > width {
        c.vx = -c.vx
    }
    if c.y-c.radius < 0 || c.y+c.radius > height {
        c.vy = -c.vy
    }
}

const (
	width = 400
	height = 400
)

type Window struct{
	circles []*Star
}

func (w *Window) Update() error {
	for _, circle := range w.circles {
        circle.Update()
    }
	return nil
}

func (w *Window) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	for i := range w.circles {
		vector.DrawFilledCircle(screen, w.circles[i].x, w.circles[i].y, w.circles[i].radius, color.White, true)
	}
}

func (w *Window) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Stars")

	circles := []*Star{
        {x: 320, y: 240, vx: 4, vy: 4, radius: 30},
    }

	if err := ebiten.RunGame(&Window{circles: circles}); err != nil {
		log.Fatal(err)
	}
}
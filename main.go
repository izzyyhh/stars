package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Star struct {
	x,y,z float32
	radius float32
}

func (c *Star) Update() {

}

func (s *Star) Show(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, s.x, s.y, s.radius, color.White, true)
}

const (
	width = 400
	height = 400
)

type Window struct{
	stars []*Star
}

func (w *Window) Update() error {
	for _, circle := range w.stars {
        circle.Update()
    }
	return nil
}

func (w *Window) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	for i := range w.stars {
		w.stars[i].Show(screen)
	}
}

func (w *Window) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Stars")

	stars := []*Star{}

	for i := 0; i < 100; i++ {
		x := float32(rand.Intn(width))
        y := float32(rand.Intn(height))
		star := Star{x: float32(x), y: float32(y), radius: 4}
		stars = append(stars, &star)
	}

	if err := ebiten.RunGame(&Window{stars: stars}); err != nil {
		log.Fatal(err)
	}
}
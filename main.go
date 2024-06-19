package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var speed float64 = 3

func Map(value, start1, stop1, start2, stop2 float64) float64 {
    return start2 + (value-start1)*(stop2-start2)/(stop1-start1)
}

type Star struct {
	x, sx, y, sy,z, pz float64
	radius float64
}

func NewStar() *Star{
	x := float64(randRange(-width/2, width/2))
	y := float64(randRange(-height/2, height/2))
	z := float64(randRange(0, width))
	sx := Map(x / z, 0, 1, 0, width) / 2
	sy := Map(y / z, 0, 1, 0, height) / 2
	radius :=  Map(z, 0, width, 16, 0)

	return &Star{
		x: x, sx: sx, y: y, sy: sy, z:z, radius: radius, pz: z,
	}
}

func (s *Star) Update() {
	s.sx = Map(s.x / s.z, 0, 1, 0, width) / 2
	s.sy = Map(s.y / s.z, 0, 1, 0, height) / 2
	s.z =  s.z - speed
	s.radius = Map(s.z, 0, width, 8, 0)

	if s.z < 1 {
		s.z = width
		s.sx = Map(s.x / s.z, 0, 1, 0, width) / 2
		s.sy = Map(s.y / s.z, 0, 1, 0, height) / 2
		s.radius = Map(s.z, 0, width, 16, 0)
	}
}

func (s *Star) Show(screen *ebiten.Image) {
	vector.DrawFilledCircle(
		screen, float32(s.sx + width/2), float32(s.sy + height/2), float32(s.radius), color.White, true,
	)
}

const (
	width = 400
	height = 400
)

type Window struct{
	stars []*Star
}

func randRange(min, max int) int {
    return rand.Intn(max-min) + min
}

func (w *Window) Update() error {
	for _, star := range w.stars {
        star.Update()
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
	ebiten.SetWindowTitle("Sternenfeld")

	stars := []*Star{}

	for i := 0; i < 100; i++ {
		stars = append(stars, NewStar())
	}

	if err := ebiten.RunGame(&Window{stars: stars}); err != nil {
		log.Fatal(err)
	}
}
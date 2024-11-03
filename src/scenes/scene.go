package scenes

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type Scene struct {
	scene fyne.Window
	container *fyne.Container
}

func NewScene(stage fyne.Window) *Scene {
	return &Scene{scene: stage, container: nil}
}

func (s *Scene) Init () {
	rect := canvas.NewRectangle(color.RGBA{R: 50, G: 50, B: 50, A: 255})
	rect.Resize(fyne.NewSize(650, 300))
	rect.Move(fyne.NewPos(0, 0))

	s.container = container.NewWithoutLayout(rect)
	s.scene.SetContent(s.container)

	sidewalk := canvas.NewRectangle(color.RGBA{R: 128, G: 128, B: 128, A: 255} )
	sidewalk.Resize(fyne.NewSize(650, 50))
	sidewalk.Move(fyne.NewPos(0, 0))

	s.container.Add(sidewalk)

	sidewalkLeft := canvas.NewRectangle(color.RGBA{R: 128, G: 128, B: 128, A: 255} )
	sidewalkLeft.Resize(fyne.NewSize(280, 25))
	sidewalkLeft.Move(fyne.NewPos(0, 205))

	s.container.Add(sidewalkLeft) 

	sidewalkRight := canvas.NewRectangle(color.RGBA{R: 128, G: 128, B: 128, A: 255} )
	sidewalkRight.Resize(fyne.NewSize(320, 25))
	sidewalkRight.Move(fyne.NewPos(340, 205))

	s.container.Add(sidewalkRight) 

	initialX := 25
	
	for i := 0; i < 21; i++ {
		line := canvas.NewRectangle(color.White)
		line.Resize(fyne.NewSize(2, 55))
		line.Move(fyne.NewPos(float32(initialX), 50))
		s.container.Add(line) 
		initialX = initialX + 30
	}
	
}


func (s *Scene) AddWidget (widget fyne.Widget) {
	s.container.Add(widget)
	s.container.Refresh()
}

func (s *Scene) AddImage (image *canvas.Image) {
	s.container.Add(image)
	s.container.Refresh()
}

func (s *Scene) DeleteImage(image *canvas.Image) {
	s.container.Remove(image)
	s.container.Refresh()
}
package views

import (
	"fmt"
	"math/rand"
	"project/src/models"
	"project/src/scenes"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

type Car struct {
	Car *canvas.Image
	Scene *scenes.Scene
}

func NewCar(s *scenes.Scene) *Car {
	return &Car{Car: nil, Scene: s}
}

func (c *Car) AddCar() {
	rand.Seed(time.Now().UnixNano()) 
	idImage := rand.Intn(5) + 1
	car := canvas.NewImageFromURI(storage.NewFileURI(fmt.Sprintf("./assets/car_%d.png", idImage)))
	car.Resize(fyne.NewSize(20,50))
	car.Move(fyne.NewPos(285, 355))
	c.Car = car
	c.Scene.AddImage(car)
}

func (c *Car) DeleteCar() {
	c.Scene.DeleteImage(c.Car)
}

func (c *Car) Update(pos models.Pos) {
	c.Car.Move(fyne.NewPos(float32(pos.X), float32(pos.Y)))
}

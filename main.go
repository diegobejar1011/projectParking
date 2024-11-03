package main

import (
	"fmt"
	"math/rand"
	"project/src/models"
	"project/src/resources"
	"project/src/scenes"
	"project/src/views"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	stage := a.NewWindow("Simulador")
	stage.CenterOnScreen()
	stage.Resize(fyne.NewSize(650, 300))
	stage.SetFixedSize(true)

	scene := scenes.NewScene(stage)
	scene.Init()

	var wg sync.WaitGroup

	button := widget.NewButton("Click", func() {
		p := resources.NewParking(make(chan int, 20), &sync.Mutex{})

		for i := 0; i < 100; i++ {
			wg.Add(i)
			go func(id int){
				// Creamos un notificador
				c1 := models.NewCar(0)

				// Creamos un suscriptor
				car := views.NewCar(scene) 
				car.AddCar()

				// Registramos el suscriptor al notificador
				fmt.Println(p)
				c1.Register(car)

				go c1.Run(p, &wg) 
			}(i)

			time.Sleep((time.Duration(rand.Intn(2)+1) * time.Second))
		}

	})

	button.Move(fyne.NewPos(300, 0))
	button.Resize(fyne.NewSize(50, 50))
	scene.AddWidget(button)


	stage.ShowAndRun()

	wg.Wait()
}
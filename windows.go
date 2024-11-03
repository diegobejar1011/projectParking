package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func windows() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hola mundo")

	// Definir las dimensiones de la ventana
	myWindow.Resize(fyne.NewSize(600, 400))

	// Definir las etiquetas 
	label := canvas.NewText("Usuario", nil)

	label.Move(fyne.NewPos(50, 50))

	// Declarar el contenido 
	content := container.NewWithoutLayout(label)

	// Asigna el contenido a la ventana
	myWindow.SetContent(content)

	// Correr la ventana
	myWindow.ShowAndRun()
}
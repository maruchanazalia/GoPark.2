package main

import (
    "estacionamiento/screens"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
)

func main() {
    a := app.New()
    w := a.NewWindow("EstacionamientoGo")

    w.CenterOnScreen()
    w.SetFixedSize(true)
    w.Resize(fyne.NewSize(1000, 800))
    screens.NuevaEscena(w) // Llama a la función de inicialización de pantalla
    w.ShowAndRun()
}

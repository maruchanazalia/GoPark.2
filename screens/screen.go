package screens

import (
    "estacionamiento/models"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/storage"
)

type EscenaJuego struct {
    ventana fyne.Window
    contenido *fyne.Container
}

func NuevaEscena(ventana fyne.Window) *EscenaJuego {
    escena := &EscenaJuego{ventana: ventana}
    escena.Renderizar()
    return escena
}

func (e *EscenaJuego) IniciarJuego() {
    estacionamiento := models.NuevoEstacionamiento(20)
    go models.GenerarCarros(100, estacionamiento)
    go e.DibujarCarros(estacionamiento)
}

func (e *EscenaJuego) Renderizar() {
    imagenFondo := canvas.NewImageFromURI(storage.NewFileURI("./assets/estacionamiento.png"))
    imagenFondo.Resize(fyne.NewSize(1000, 800))
    imagenFondo.Move(fyne.NewPos(0, 0))

    e.contenido = container.NewWithoutLayout(imagenFondo)
    e.ventana.SetContent(e.contenido)
    e.IniciarJuego()
}

func (e *EscenaJuego) DibujarCarros(estacionamiento *models.Estacionamiento) {
    for {
        imagen := <-estacionamiento.CanalDibujoCarro
        e.contenido.Add(imagen)
        e.ventana.Canvas().Refresh(e.contenido)
    }
}

package models

import (
    "fmt"
    "math/rand"
    "time"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/storage"
)

type Carro struct {
    estacionamiento *Estacionamiento
    indice          int
    lugarOcupado    int
    imagen          *canvas.Image
}

func NuevoCarro(e *Estacionamiento, img *canvas.Image) *Carro {
    return &Carro{
        estacionamiento: e,
        imagen: img,
    }
}

func (c *Carro) MoverCarro() {
    c.estacionamiento.Espacios <- true
    c.estacionamiento.mutex.Lock()

    for i := 0; i < len(c.estacionamiento.LugaresEstacionamiento); i++ {
        if !c.estacionamiento.LugaresEstacionamiento[i].ocupado {
            c.imagen.Move(fyne.NewPos(c.estacionamiento.LugaresEstacionamiento[i].x, c.estacionamiento.LugaresEstacionamiento[i].y))
            c.imagen.Refresh()
            c.lugarOcupado = i
            c.estacionamiento.LugaresEstacionamiento[i].ocupado = true
            break
        }
    }

    fmt.Println("Carro ", c.indice, " entra")
    c.estacionamiento.mutex.Unlock()

    tiempoAleatorio := rand.Intn(30) + 5
    time.Sleep(time.Duration(tiempoAleatorio) * time.Second)

    c.estacionamiento.mutex.Lock()
    <-c.estacionamiento.Espacios
    c.estacionamiento.LugaresEstacionamiento[c.lugarOcupado].ocupado = false
    c.imagen.Move(fyne.NewPos(460, 45))
    fmt.Println("Carro ", c.indice, " sale")
    time.Sleep(300 * time.Millisecond)
    c.estacionamiento.mutex.Unlock()
    c.imagen.Move(fyne.NewPos(460000, 45000))
}

func GenerarCarros(n int, e *Estacionamiento) {
    e.Espacios <- true
    for i := 0; i < n; i++ {
        numCarro := rand.Intn(8) + 1
        nombreImagen := fmt.Sprintf("./assets/auto%d.png", numCarro)

        imagenCarro := canvas.NewImageFromURI(storage.NewFileURI(nombreImagen))
        imagenCarro.Resize(fyne.NewSize(70, 120))
        imagenCarro.Move(fyne.NewPos(460, 650))

        carro := NuevoCarro(e, imagenCarro)
        carro.indice = i + 1

        e.CanalDibujoCarro <- imagenCarro
        time.Sleep(time.Millisecond * 400)
        go carro.MoverCarro()
        time.Sleep(time.Duration(rand.ExpFloat64() * float64(time.Second)))
    }
}

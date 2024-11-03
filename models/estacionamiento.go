package models

import (
    "sync"
    "fyne.io/fyne/v2/canvas"
)

type Espacio struct {
    x        float32
    y        float32
    ocupado  bool
}

type Estacionamiento struct {
    Espacios        chan bool
    CanalDibujoCarro chan *canvas.Image
    mutex           sync.Mutex
    LugaresEstacionamiento []Espacio
}

func NuevoEstacionamiento(numeroEspacios int) *Estacionamiento {
    return &Estacionamiento{
        Espacios: make(chan bool, numeroEspacios+1),
        CanalDibujoCarro: make(chan *canvas.Image, 1),
        LugaresEstacionamiento: []Espacio{
            // Columna izquierda
            {x: 100, y: 150, ocupado: false},
            {x: 100, y: 260, ocupado: false},
            {x: 100, y: 350, ocupado: false},
            {x: 100, y: 440, ocupado: false},
            {x: 100, y: 540, ocupado: false},
            // Columna derecha
            {x: 300, y: 150, ocupado: false},
            {x: 300, y: 260, ocupado: false},
            {x: 300, y: 350, ocupado: false},
            {x: 300, y: 440, ocupado: false},
            {x: 300, y: 540, ocupado: false},
            // Columna izquierda inferior
            {x: 600, y: 150, ocupado: false},
            {x: 600, y: 260, ocupado: false},
            {x: 600, y: 350, ocupado: false},
            {x: 600, y: 440, ocupado: false},
            {x: 600, y: 540, ocupado: false},
            // Columna derecha inferior
            {x: 800, y: 150, ocupado: false},
            {x: 800, y: 260, ocupado: false},
            {x: 800, y: 350, ocupado: false},
            {x: 800, y: 440, ocupado: false},
            {x: 800, y: 540, ocupado: false},
        },
    }
}

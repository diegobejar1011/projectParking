package models

type Pos struct {
	X int32
	Y int32
}

type Suscriber interface {
	Update(pos Pos)
}

type Notificator interface {
	Register(suscriber Suscriber)
	Unregister(suscriber Suscriber)
	NotifyAll()
}
package models

import (
	"math/rand"
	"project/src/resources"
	"sync"
	"time"
)

type Car struct {
	id int
	parkingTime time.Duration
	space int
	posX, posY int32
	suscribers []Suscriber
}

func NewCar(id int) *Car {
	return &Car{
		id: id,
		posX: 285, 
		posY: 355, 
		parkingTime: time.Duration(rand.Intn(10)+10) * time.Second,
		space: 0,
	}
}

func (c *Car) Run (p *resources.Parking, wg *sync.WaitGroup) {

	defer wg.Done()

	
	 for i := 0; i < 5; i++ {
	 	c.posY = c.posY - 20
	 	c.NotifyAll()
	 	time.Sleep(time.Millisecond * 200)
	 }

	// c.posY = 255
	// c.NotifyAll()

	c.Enter(p)

	time.Sleep(c.parkingTime)

	c.Leave(p)
	c.Unregister()
}

func (c *Car) Enter (p *resources.Parking) {
	p.GetSpaces() <- c.id
	p.GetEntrance().Lock()

	spaces := p.GetSpacesArray()

	 for i := 0; i < 3; i++ {
	 	c.posY = c.posY - 20
	 	c.NotifyAll()
	 	time.Sleep(time.Millisecond * 100)
	 }

	// c.posY = 215
	// c.NotifyAll()
	// time.Sleep(time.Millisecond * 200)
	
	p.GetEntrance().Unlock()

	for i := 0; i < len(spaces); i++ {
		if !spaces[i] {
			spaces[i] = true
			c.space = i
			c.posY = 55
			c.posX = 25 + int32(5 +(i*30))
			c.NotifyAll()
			break
		}
	}

	p.SetSpacesArray(spaces)
}

func (c *Car) Leave (p *resources.Parking) {
	p.GetEntrance().Lock()
	<- p.GetSpaces()

	spaces := p.GetSpacesArray()
	spaces[c.space] = false
	p.SetSpacesArray(spaces)

	// c.posY = 230
	// c.NotifyAll()
	// time.Sleep(time.Millisecond * 100)
	// c.posY = 500
	// c.NotifyAll()

	c.posY = 120

	for i := 0; i < 10; i++ {
		c.posX = 315
		c.posY = c.posY + 20
		c.NotifyAll()
		time.Sleep(time.Millisecond * 75)
	}

	p.GetEntrance().Unlock()



}

func (c *Car) Register(suscriber Suscriber) {
	c.suscribers = append(c.suscribers, suscriber)
}

func (c *Car) Unregister() {
	c.suscribers = []Suscriber{}
}

func (c *Car) NotifyAll() {

	for _, o := range c.suscribers {
		o.Update(Pos{X: c.posX, Y: c.posY})
	}
}
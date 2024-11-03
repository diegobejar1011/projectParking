package resources

import "sync"

type Parking struct {
	spaces  chan int
	entrace *sync.Mutex
	spacesArray [20]bool
}

func NewParking(spaces chan int, entrance *sync.Mutex) *Parking {
	return &Parking{
		spaces: spaces,
		entrace: entrance,
		spacesArray: [20]bool{},
	}
}

func (p *Parking) GetSpaces () chan int {
	return p.spaces
}

func (p *Parking) GetEntrance () *sync.Mutex {
	return p.entrace
}

func (p *Parking) GetSpacesArray () [20]bool {
	return p.spacesArray
}

func (p *Parking) SetSpacesArray(spacesArray [20]bool) {
	p.spacesArray = spacesArray
}


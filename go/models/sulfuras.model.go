package models

type Sulfuras struct {
	Name    string
	SellIn  int
	Quality int
}

func (ab *Sulfuras) UpdateQuality() {
}

func (ab *Sulfuras) GetItem() *Item {
	return &Item{
		name:    ab.Name,
		sellIn:  ab.SellIn,
		quality: ab.Quality,
	}
}

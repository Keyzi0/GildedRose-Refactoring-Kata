package models

type Sulfuras struct {
	*Item
}

func NewSulfuras(item *Item) *Sulfuras {
	return &Sulfuras{item}
}

func (ab *Sulfuras) UpdateQuality() {
}

func (ab *Sulfuras) GetItem() *Item {
	return &Item{
		name:    ab.name,
		sellIn:  ab.sellIn,
		quality: ab.quality,
	}
}

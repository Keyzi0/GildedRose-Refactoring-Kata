package models

type Conjured struct {
	*Item
}

func NewConjured(item *Item) *Conjured {
	return &Conjured{item}
}

func (ab *Conjured) UpdateQuality() {
	if ab.quality > 0 {
		ab.quality -= 2
		if ab.quality < 0 {
			ab.quality = 0
		}
	}
	ab.sellIn -= 1
	if ab.sellIn < 0 && ab.quality > 0 {
		ab.quality -= 2
		if ab.quality < 0 {
			ab.quality = 0
		}
	}
}

func (ab *Conjured) GetItem() *Item {
	return &Item{
		name:    ab.name,
		sellIn:  ab.sellIn,
		quality: ab.quality,
	}
}

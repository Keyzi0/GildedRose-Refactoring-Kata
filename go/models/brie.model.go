package models

type AgedBrie struct {
	*Item
}

func NewAgedBrie(item *Item) *AgedBrie {
	return &AgedBrie{item}
}

func (ab *AgedBrie) UpdateQuality() {
	if ab.quality < 50 {
		ab.quality += 1
	}
	ab.sellIn -= 1
	if ab.sellIn < 0 && ab.quality < 50 {
		ab.quality += 1
	}
}

func (ab *AgedBrie) GetItem() *Item {
	return &Item{
		name:    ab.name,
		sellIn:  ab.sellIn,
		quality: ab.quality,
	}
}

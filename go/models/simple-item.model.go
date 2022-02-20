package models

type SimpleItem struct {
	*Item
}

func NewSimpleItem(item *Item) *SimpleItem {
	return &SimpleItem{item}
}

func (ab *SimpleItem) UpdateQuality() {
	if ab.quality > 0 {
		ab.quality -= 1
	}
	ab.sellIn -= 1
	if ab.sellIn < 0 && ab.quality > 0 {
		ab.quality -= 1
	}
}

func (ab *SimpleItem) GetItem() *Item {
	return &Item{
		name:    ab.name,
		sellIn:  ab.sellIn,
		quality: ab.quality,
	}
}

package models

type SpecifiedItemInterface interface {
	UpdateQuality()
	GetItem() *Item
}

func NewSpecifiedItem(item *Item) SpecifiedItemInterface {
	return getTypedItem(item)
}

func getTypedItem(item *Item) SpecifiedItemInterface {
	switch item.name {
	case "Aged Brie":
		return NewAgedBrie(item)
	case "Sulfuras, Hand of Ragnaros":
		return NewSulfuras(item)
	case "Backstage passes to a TAFKAL80ETC concert":
		return NewBackstagePass(item)
	default:
		return NewSimpleItem(item)
	}
}

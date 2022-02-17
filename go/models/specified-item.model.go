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
		return &AgedBrie{
			Name:    item.name,
			SellIn:  item.sellIn,
			Quality: item.quality,
		}
	case "Sulfuras, Hand of Ragnaros":
		return &Sulfuras{
			Name:    item.name,
			SellIn:  item.sellIn,
			Quality: item.quality,
		}
	case "Backstage passes to a TAFKAL80ETC concert":
		return &BackstagePass{
			Name:    item.name,
			SellIn:  item.sellIn,
			Quality: item.quality,
		}
	default:
		return &SimpleItem{
			Name:    item.name,
			SellIn:  item.sellIn,
			Quality: item.quality,
		}
	}
}

package models

import (
	"fmt"
	"strings"
)

type SpecifiedItemInterface interface {
	UpdateQuality()
	GetItem() *Item
}

func NewSpecifiedItem(item *Item) SpecifiedItemInterface {
	itm := getTypedItem(item)
	fmt.Println(fmt.Sprintf("item %v, %v, %v", item, *item, &item))
	fmt.Println(fmt.Sprintf("item %v, %v, %v\n\n", itm, itm, &itm))
	return itm //getTypedItem(item)
}

func getTypedItem(item *Item) SpecifiedItemInterface {
	switch true {
	case item.name == "Aged Brie":
		return NewAgedBrie(item)
	case item.name == "Sulfuras, Hand of Ragnaros":
		return NewSulfuras(item)
	case strings.HasPrefix(item.name, "Backstage passes"):
		return NewBackstagePass(item)
	case strings.HasPrefix(item.name, "Conjured"):
		return NewConjured(item)
	default:
		return NewSimpleItem(item)
	}
}

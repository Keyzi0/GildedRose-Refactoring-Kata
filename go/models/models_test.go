package models

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func Test_Aged_Brie(t *testing.T) {
	name := "Aged Brie"

	var updatedItems = []*Item{
		{name, -1, 2},
		{name, 4, 6},
		{name, -11, 12},
		{name, 9, 11},
		{name, 14, 16},
		{name, -1, -3},
		{name, -1, 50},
		{name, -1, 51},
	}

	doUpdateQuality(t, name, updatedItems)
}

func Test_Simple_Item(t *testing.T) {
	name := "Simple Item"

	var updatedItems = []*Item{
		{name, -1, 0},
		{name, 4, 4},
		{name, -11, 8},
		{name, 9, 9},
		{name, 14, 14},
		{name, -1, -5},
		{name, -1, 47},
		{name, -1, 49},
	}

	doUpdateQuality(t, name, updatedItems)
}

func Test_Sulfuras(t *testing.T) {
	name := "Sulfuras, Hand of Ragnaros"

	var updatedItems = []*Item{
		{name, 0, 0},
		{name, 5, 5},
		{name, -10, 10},
		{name, 10, 10},
		{name, 15, 15},
		{name, 0, -5},
		{name, 0, 49},
		{name, 0, 51},
	}

	doUpdateQuality(t, name, updatedItems)
}

func Test_Backstage_Passes(t *testing.T) {
	name := "Backstage passes to a TAFKAL80ETC concert"

	var updatedItems = []*Item{
		{name, -1, 0},
		{name, 4, 8},
		{name, -11, 0},
		{name, 9, 12},
		{name, 14, 16},
		{name, -1, 0},
		{name, -1, 0},
		{name, -1, 0},
	}

	doUpdateQuality(t, name, updatedItems)
}

// =========================================================
func printItems(name string, items []*Item) {
	fmt.Printf("\n%v\n", name)
	for i, item := range items {
		fmt.Println(fmt.Sprintf("%v. s:%v q:%v", i, item.sellIn, item.quality))
	}
}

func doUpdateQuality(t *testing.T, name string, expectedItems []*Item) {
	var items = []*Item{
		{name, 0, 0},
		{name, 5, 5},
		{name, -10, 10},
		{name, 10, 10},
		{name, 15, 15},
		{name, 0, -5},
		{name, 0, 49},
		{name, 0, 51},
	}

	printItems(name, items)

	for i := range items {
		si := NewSpecifiedItem(items[i])
		si.UpdateQuality()
		items[i] = si.GetItem()
	}

	printItems(name, items)

	for i, item := range items {
		assert.Equal(t, *item, *expectedItems[i])
	}
}

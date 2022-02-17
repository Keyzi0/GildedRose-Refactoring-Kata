package models

type SimpleItem struct {
	Name    string
	SellIn  int
	Quality int
}

func (ab *SimpleItem) UpdateQuality() {
	if ab.Quality > 0 {
		ab.Quality -= 1
	}
	ab.SellIn -= 1
	if ab.SellIn < 0 && ab.Quality > 0 {
		ab.Quality -= 1
	}
}

func (ab *SimpleItem) GetItem() *Item {
	return &Item{
		name:    ab.Name,
		sellIn:  ab.SellIn,
		quality: ab.Quality,
	}
}

package models

type AgedBrie struct {
	Name    string
	SellIn  int
	Quality int
}

func (ab *AgedBrie) UpdateQuality() {
	if ab.Quality < 50 {
		ab.Quality += 1
	}
	ab.SellIn -= 1
	if ab.SellIn < 0 && ab.Quality < 50 {
		ab.Quality += 1
	}
}

func (ab *AgedBrie) GetItem() *Item {
	return &Item{
		name:    ab.Name,
		sellIn:  ab.SellIn,
		quality: ab.Quality,
	}
}

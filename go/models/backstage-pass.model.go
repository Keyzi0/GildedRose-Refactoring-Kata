package models

type BackstagePass struct {
	Name    string
	SellIn  int
	Quality int
}

func (ab *BackstagePass) UpdateQuality() {
	if ab.Quality < 50 {
		ab.Quality += 1
		if ab.SellIn < 11 {
			if ab.Quality < 50 {
				ab.Quality += 1
			}
		}
		if ab.SellIn < 6 {
			if ab.Quality < 50 {
				ab.Quality += 1
			}
		}
	}
	ab.SellIn -= 1
	if ab.SellIn < 0 {
		ab.Quality = 0
	}
}

func (ab *BackstagePass) GetItem() *Item {
	return &Item{
		name:    ab.Name,
		sellIn:  ab.SellIn,
		quality: ab.Quality,
	}
}

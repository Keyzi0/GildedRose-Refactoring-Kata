package models

type BackstagePass struct {
	*Item
}

func NewBackstagePass(item *Item) *BackstagePass {
	return &BackstagePass{item}
}

func (ab *BackstagePass) UpdateQuality() {
	if ab.quality < 50 {
		ab.quality += 1
		if ab.sellIn < 11 {
			if ab.quality < 50 {
				ab.quality += 1
			}
		}
		if ab.sellIn < 6 {
			if ab.quality < 50 {
				ab.quality += 1
			}
		}
	}
	ab.sellIn -= 1
	if ab.sellIn < 0 {
		ab.quality = 0
	}
}

func (ab *BackstagePass) GetItem() *Item {
	return &Item{
		name:    ab.name,
		sellIn:  ab.sellIn,
		quality: ab.quality,
	}
}

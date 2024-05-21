package gildedrose

const (
	defaultQualityDropRate  = 1
	conjuredQualityDropRate = 2
)

type ItemHandler func(Item) Item

type Item struct {
	Name            string
	SellIn, Quality int
}

type GuildRoseShop struct {
	itemHandler map[string]ItemHandler
}

func NewGuildRoseShop(m map[string]ItemHandler) GuildRoseShop {
	grs := GuildRoseShop{itemHandler: m}
	return grs
}

func (grs *GuildRoseShop) UpdateQuality(items []Item) []Item {
	var retItem []Item
	for _, e := range items {
		handler, ok := grs.itemHandler[e.Name]
		if ok {
			retItem = append(retItem, handler(e))
		} else {
			sellIn := e.SellIn - 1
			retItem = append(retItem, Item{
				Name:    e.Name,
				SellIn:  sellIn,
				Quality: calculateQuality(sellIn, e.Quality, defaultQualityDropRate)})
		}
	}
	return retItem
}

func agedBrieHandler(i Item) Item {
	return Item{Name: i.Name, SellIn: i.SellIn - 1, Quality: qualityCheck(i.Quality + 1)}
}

func sulfurasHandler(i Item) Item {
	return Item{Name: i.Name, SellIn: i.SellIn, Quality: i.Quality}
}

func backstagepassesHandler(i Item) Item {
	retItem := Item{Name: i.Name, SellIn: i.SellIn - 1, Quality: qualityCheck(i.Quality + 1)}
	if retItem.SellIn < 0 {
		retItem.Quality = 0
		return retItem
	}
	if retItem.SellIn < 11 && retItem.SellIn > 5 {
		retItem.Quality = qualityCheck(retItem.Quality + 1)
	}
	if retItem.SellIn < 6 {
		retItem.Quality = qualityCheck(retItem.Quality + 2)
	}
	return retItem
}

func conjuredHandler(i Item) Item {
	sellIn := i.SellIn - 1
	return Item{Name: i.Name, SellIn: sellIn, Quality: calculateQuality(sellIn, i.Quality, conjuredQualityDropRate)}
}

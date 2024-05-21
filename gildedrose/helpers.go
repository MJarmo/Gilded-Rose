package gildedrose

func MakeItemHandlerMap() map[string]ItemHandler {
	m := make(map[string]ItemHandler)
	m["Aged Brie"] = agedBrieHandler
	m["Sulfuras, Hand of Ragnaros"] = sulfurasHandler
	m["Backstage passes to a TAFKAL80ETC concert"] = backstagepassesHandler
	m["Conjured Mana Cake"] = conjuredHandler
	return m
}

func qualityCheck(quality int) int {
	if quality > 50 {
		return 50
	} else if quality <= 0 {
		return 0
	}
	return quality
}

func calculateQuality(sellIn, quality, rate int) int {
	if quality > 50 {
		return 50
	}
	if sellIn < 0 {
		return qualityCheck(quality - 2*rate)
	}
	return qualityCheck(quality - rate)

}

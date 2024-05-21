package main

import (
	"fmt"
	"main/gildedrose"
)

type GuildRoseShop interface {
	UpdateQuality(items []gildedrose.Item) []gildedrose.Item
}

func main() {
	var items = []gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 30},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Conjured Mana Cake", 3, 6}, // <-- :O
	}

	days := 4

	grs := gildedrose.NewGuildRoseShop(gildedrose.MakeItemHandlerMap())
	for day := 0; day < days; day++ {
		fmt.Printf("-------- day %d --------\n", day)
		fmt.Println("Name, SellIn, Quality")
		for i := 0; i < len(items); i++ {
			fmt.Println(items[i])
		}
		fmt.Println("")
		items = grs.UpdateQuality(items)
	}
}

package gildedrose

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldRunHandlersWithSuccess(t *testing.T) {
	var test = []struct {
		input    Item
		expected Item
		handler  ItemHandler
	}{
		{Item{"brie", 10, 10}, Item{"brie", 9, 11}, agedBrieHandler},
		{Item{"brie", 0, 10}, Item{"brie", -1, 11}, agedBrieHandler},
		{Item{"brie", 3, 50}, Item{"brie", 2, 50}, agedBrieHandler},
		{Item{"backstage", 12, 30}, Item{"backstage", 11, 31}, backstagepassesHandler},
		{Item{"backstage", 10, 30}, Item{"backstage", 9, 32}, backstagepassesHandler},
		{Item{"backstage", 4, 30}, Item{"backstage", 3, 33}, backstagepassesHandler},
		{Item{"backstage", 0, 30}, Item{"backstage", -1, 0}, backstagepassesHandler},
		{Item{"conjured", 0, 0}, Item{"conjured", -1, 0}, conjuredHandler},
		{Item{"conjured", 21, 35}, Item{"conjured", 20, 33}, conjuredHandler},
		{Item{"conjured", 0, 30}, Item{"conjured", -1, 26}, conjuredHandler},
		{Item{"conjured", 0, 30}, Item{"conjured", 0, 30}, sulfurasHandler},
	}

	for _, tt := range test {
		strs := (runtime.FuncForPC(reflect.ValueOf(tt.handler).Pointer()).Name())
		t.Run(strs, func(t *testing.T) {
			result := tt.handler(tt.input)
			assert.Equal(t, tt.expected, result)
		})

	}
}

func Test_ShouldRunAgedBrieHandlerWithSuccess(t *testing.T) {
	//given:
	m := make(map[string]ItemHandler)
	m["Juan Pablito"] = func(i Item) Item { return Item{Name: "Juan Pablito", SellIn: 21, Quality: 37} }
	items := []Item{
		{Name: "non given item", SellIn: 2, Quality: 5},
		{Name: "Juan Pablito", SellIn: 1, Quality: 1}}
	expectedItems := []Item{
		{Name: "non given item", SellIn: 1, Quality: 4},
		{Name: "Juan Pablito", SellIn: 21, Quality: 37}}
	grs := NewGuildRoseShop(m)
	//when:
	response := grs.UpdateQuality(items)
	//then:
	assert.ElementsMatch(t, response, expectedItems)
}

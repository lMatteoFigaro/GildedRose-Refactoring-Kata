package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func TestUpdateQuality(t *testing.T) {
	tests := []struct {
		name   string
		items  []*gildedrose.Item
		expect []*gildedrose.Item
	}{
		{
			name: "Aged Brie",
			items: []*gildedrose.Item{
				{Name: "Aged Brie", SellIn: 2, Quality: 0},
			},
			expect: []*gildedrose.Item{
				{Name: "Aged Brie", SellIn: 1, Quality: 1},
			},
		},
		{
			name: "Aged Brie expired",
			items: []*gildedrose.Item{
				{Name: "Aged Brie", SellIn: 0, Quality: 0},
			},
			expect: []*gildedrose.Item{
				{Name: "Aged Brie", SellIn: -1, Quality: 2},
			},
		},
		{
			name: "Aged Brie maximum quality",
			items: []*gildedrose.Item{
				{Name: "Aged Brie", SellIn: 30, Quality: 50},
			},
			expect: []*gildedrose.Item{
				{Name: "Aged Brie", SellIn: 29, Quality: 50},
			},
		},
		{
			name: "Backstage passes 0 day",
			items: []*gildedrose.Item{
				{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 0, Quality: 20},
			},
			expect: []*gildedrose.Item{
				{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: -1, Quality: 0},
			},
		},
		{
			name: "Backstage passes 9 days",
			items: []*gildedrose.Item{
				{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 9, Quality: 20},
			},
			expect: []*gildedrose.Item{
				{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 8, Quality: 22},
			},
		},
		{
			name: "Backstage passes",
			items: []*gildedrose.Item{
				{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 5, Quality: 20},
			},
			expect: []*gildedrose.Item{
				{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 4, Quality: 23},
			},
		},
		{
			name: "regular item",
			items: []*gildedrose.Item{
				{Name: "Regular Item", SellIn: 5, Quality: 20},
			},
			expect: []*gildedrose.Item{
				{Name: "Regular Item", SellIn: 4, Quality: 19},
			},
		},
		{
			name: "regular item expired",
			items: []*gildedrose.Item{
				{Name: "Regular Item", SellIn: 0, Quality: 20},
			},
			expect: []*gildedrose.Item{
				{Name: "Regular Item", SellIn: -1, Quality: 18},
			},
		},
		{
			name: "regular item no quality",
			items: []*gildedrose.Item{
				{Name: "Regular Item", SellIn: 30, Quality: 0},
			},
			expect: []*gildedrose.Item{
				{Name: "Regular Item", SellIn: 29, Quality: 0},
			},
		},
		// more test cases
		{
			name: "sulfuras regular",
			items: []*gildedrose.Item{
				{Name: "Sulfuras, Hand of Ragnaros", SellIn: 30, Quality: 40},
			},
			expect: []*gildedrose.Item{
				{Name: "Sulfuras, Hand of Ragnaros", SellIn: 30, Quality: 40},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gildedrose.UpdateQuality(tt.items)
			for i, item := range tt.items {
				if item.Name != tt.expect[i].Name || item.SellIn != tt.expect[i].SellIn || item.Quality != tt.expect[i].Quality {
					t.Errorf("UpdateQuality() = %v, want %v", item, tt.expect[i])
				}
			}
		})
	}
}

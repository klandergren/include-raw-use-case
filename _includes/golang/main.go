package main

import (
	"sort"
)

func main() {
	fruit := []string{"durian", "canteloupe", "apple", "banana"}

	// sort strings. sort string slice
	sort.Strings(fruit)                              // ascending
	sort.Sort(sort.Reverse(sort.StringSlice(fruit))) // descending

	// note: sort.StringSlice attaches the methods of Interface to []string,
	// sorting in increasing order.

	counts := []int{3, 5, 4, 2, 1}
	// sort ints. sort int slice
	sort.Ints(counts)                              // ascending
	sort.Sort(sort.Reverse(sort.IntSlice(counts))) // descending

	// note: sort.IntSlice attaches the methods of Interface to []int, sorting in
	// increasing order.

	// sort structs. sort struct slice
	type coord struct {
		x int
		y int
	}
	coords := []coord{{3, 2}, {1, 2}, {1, 1}}
	sort.Slice(coords, func(a int, b int) bool {
		if coords[a].y == coords[b].y {
			return coords[a].x < coords[b].x
		}

		return coords[a].y < coords[b].y
	})
}

package ui_test

import (
	"reflect"
	"testing"

	"github.com/mbriggs/gesso/ui"
)

func TestPaginationSeries(t *testing.T) {
	t.Parallel()
	num := func(n int) ui.PaginationEntry { return ui.PaginationEntry{Number: n} }
	gap := ui.PaginationEntry{Gap: true}

	cases := []struct {
		name             string
		page, pages, win int
		want             []ui.PaginationEntry
	}{
		{"single page returns nil", 1, 1, 1, nil},
		{"zero pages returns nil", 1, 0, 1, nil},
		{
			"first page small set",
			1, 3, 1,
			[]ui.PaginationEntry{num(1), num(2), num(3)},
		},
		{
			"middle page introduces gaps on both sides",
			5, 10, 1,
			[]ui.PaginationEntry{num(1), gap, num(4), num(5), num(6), gap, num(10)},
		},
		{
			"current page near end leaves a left gap",
			9, 10, 1,
			[]ui.PaginationEntry{num(1), gap, num(8), num(9), num(10)},
		},
		{
			"current page near start leaves a right gap",
			2, 10, 1,
			[]ui.PaginationEntry{num(1), num(2), num(3), gap, num(10)},
		},
		{
			"window=0 only first/current/last",
			5, 10, 0,
			[]ui.PaginationEntry{num(1), gap, num(5), gap, num(10)},
		},
		{
			"window=2 widens the visible band",
			5, 10, 2,
			[]ui.PaginationEntry{num(1), gap, num(3), num(4), num(5), num(6), num(7), gap, num(10)},
		},
		{
			"out-of-range page clips to bounds",
			15, 10, 1,
			[]ui.PaginationEntry{num(1), gap, num(10)},
		},
		{
			"negative window coerces to 0",
			5, 10, -1,
			[]ui.PaginationEntry{num(1), gap, num(5), gap, num(10)},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ui.PaginationSeries(tc.page, tc.pages, tc.win)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("PaginationSeries(%d, %d, %d) = %+v, want %+v", tc.page, tc.pages, tc.win, got, tc.want)
			}
		})
	}
}

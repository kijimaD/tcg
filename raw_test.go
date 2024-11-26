package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	str := `
[[place]]
Name = "example1"
Title = "title1"
PlaceCategory = "HISTORY"
BgPath = "bgpath1"
KeyPath = "keypath1"
Descs = ["desc1"]
Location = "location1"

[[place]]
Name = "example2"
Title = "title2"
PlaceCategory = "SCENIC"
BgPath = "bgpath2"
KeyPath = "keypath2"
Descs = ["desc2"]
Location = "location2"
`
	raw, err := Load(str)
	assert.NoError(t, err)

	expect := RawMaster{
		Raws: Raws{
			Places: []Place{
				Place{
					Name:          "example1",
					Title:         "title1",
					PlaceCategory: placeCategoryHistory,
					BgPath:        "bgpath1",
					KeyPath:       "keypath1",
					Descs:         []string{"desc1"},
					Location:      "location1",
				},
				Place{
					Name:          "example2",
					Title:         "title2",
					PlaceCategory: placeCategoryScenic,
					BgPath:        "bgpath2",
					KeyPath:       "keypath2",
					Descs:         []string{"desc2"},
					Location:      "location2",
				},
			},
		},
		PlaceIndex: map[string]int{
			"example1": 0,
			"example2": 1,
		},
	}
	assert.Equal(t, expect, raw)
}

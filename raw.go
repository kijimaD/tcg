package main

import "github.com/BurntSushi/toml"

type RawMaster struct {
	Raws       Raws
	PlaceIndex map[string]int
}

type Raws struct {
	Places []Place `toml:"place"`
}

func Load(content string) (RawMaster, error) {
	rw := RawMaster{}
	rw.PlaceIndex = map[string]int{}

	_, err := toml.Decode(string(content), &rw.Raws)
	if err != nil {
		return rw, err
	}

	for i, place := range rw.Raws.Places {
		rw.PlaceIndex[place.Name] = i
	}

	return rw, nil
}

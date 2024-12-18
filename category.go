package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrInvalidEnumType = errors.New("rawファイルでenumに無効な値が指定された")

type placeCategory string

const (
	// 歴史
	placeCategoryHistory = placeCategory("HISTORY")
	// 景勝
	placeCategoryScenic = placeCategory("SCENIC")
	// レア
	placeCategoryRare = placeCategory("RARE")
)

func (pc placeCategory) Valid() error {
	switch pc {
	case placeCategoryHistory, placeCategoryScenic, placeCategoryRare:
		return nil
	}

	return fmt.Errorf("%w: %s", ErrInvalidEnumType, pc)
}

func (pc placeCategory) String() string {
	var result string
	switch pc {
	case placeCategoryHistory:
		result = "歴"
	case placeCategoryScenic:
		result = "景"
	case placeCategoryRare:
		result = "珍"
	default:
		log.Fatal("invalid place category")
	}

	return result
}

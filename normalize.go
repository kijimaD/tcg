package main

import (
	"image"
	"image/png"
	"os"

	"golang.org/x/image/draw"
)

func SquareTrimImage(img image.Image, size int) image.Image {
	// 画像のサイズを取得する
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	// 短辺の長さを取得する
	shorter := width
	if height < shorter {
		shorter = height
	}

	// 左上の座標を計算する
	top := (height - shorter) / 2
	left := (width - shorter) / 2

	// 新しい画像を用意する
	newImage := image.NewRGBA(image.Rect(0, 0, size, size))

	// 画像の中心を切り抜きつつ、最終的なサイズ(size × size)になるようにリサイズする
	draw.BiLinear.Scale(newImage, newImage.Bounds(), img, image.Rect(left, top, width-left, height-top), draw.Over, nil)

	return newImage
}

func loadImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func saveImage(img image.Image, filePath string) error {
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	return png.Encode(out, img)
}

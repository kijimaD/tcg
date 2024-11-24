package main

import (
	"image"
	"image/png"
	"os"

	"github.com/fogleman/gg"
	"golang.org/x/image/draw"
)

// キービジュアル用の画像を用意する
func normalizeKey(inputPath string, outputPath string) error {
	img, err := loadImage(inputPath)
	if err != nil {
		return err
	}
	croppedImg := trimImage(img, keyVisualWidth, keyVisualHeight)
	saveImage(croppedImg, outputPath)

	return nil
}

// カード背景用の画像を用意する
func normalizeBg(inputPath string, outputPath string) error {
	img, err := loadImage(inputPath)
	if err != nil {
		return err
	}
	newImg := trimImage(img, 250, 400)
	newImg = round(newImg)
	saveImage(newImg, outputPath)

	return nil
}

func trimImage(img image.Image, w int, h int) image.Image {
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
	newImage := image.NewRGBA(image.Rect(0, 0, w, h))

	// 画像の中心を切り抜きつつ、最終的なサイズになるようにリサイズする
	draw.BiLinear.Scale(newImage, newImage.Bounds(), img, image.Rect(left, top, width-left, height-top), draw.Over, nil)

	return newImage
}

// よく理解していない
func roundCornersWithAntialias(img image.Image, radius int) *image.RGBA {
	// 元画像のサイズを取得
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	// 描画コンテキストを作成
	dc := gg.NewContext(width, height)

	// 四隅を丸くするためのパスを作成
	dc.NewSubPath()
	dc.MoveTo(float64(radius), 0)                                                           // 左上から始める
	dc.LineTo(float64(width-radius), 0)                                                     // 上辺
	dc.QuadraticTo(float64(width), 0, float64(width), float64(radius))                      // 右上の曲線
	dc.LineTo(float64(width), float64(height-radius))                                       // 右辺
	dc.QuadraticTo(float64(width), float64(height), float64(width-radius), float64(height)) // 右下の曲線
	dc.LineTo(float64(radius), float64(height))                                             // 下辺
	dc.QuadraticTo(0, float64(height), 0, float64(height-radius))                           // 左下の曲線
	dc.LineTo(0, float64(radius))                                                           // 左辺
	dc.QuadraticTo(0, 0, float64(radius), 0)                                                // 左上の曲線
	dc.ClosePath()

	// パスをクリップ領域として設定
	dc.Clip()

	// 元画像を描画
	dc.DrawImage(img, 0, 0)

	// RGBA画像として返す
	return dc.Image().(*image.RGBA)
}

func round(img image.Image) image.Image {
	// 角を丸くする
	radius := 16
	roundedImg := roundCornersWithAntialias(img, radius)

	return roundedImg
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

package main

import (
	"image"
	"image/color"
	"image/png"
	"os"

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

func roundCornersOnly(img image.Image, radius int) *image.RGBA {
	// 画像サイズを取得
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	// 出力用のRGBA画像を作成
	dst := image.NewRGBA(bounds)

	// 全体をコピー
	draw.Draw(dst, bounds, img, bounds.Min, draw.Src)

	// 各角に異なるマスクを適用
	applyCornerMask(dst, radius, 0, 0, "tl")                        // 左上 (top-left)
	applyCornerMask(dst, radius, width-radius, 0, "tr")             // 右上 (top-right)
	applyCornerMask(dst, radius, 0, height-radius, "bl")            // 左下 (bottom-left)
	applyCornerMask(dst, radius, width-radius, height-radius, "br") // 右下 (bottom-right)

	return dst
}

func applyCornerMask(dst *image.RGBA, radius, xOffset, yOffset int, corner string) {
	// 角ごとの丸みを作成
	for y := 0; y < radius; y++ {
		for x := 0; x < radius; x++ {
			// 円の内外を判定
			distance := (x-radius)*(x-radius) + (y-radius)*(y-radius)
			if distance > radius*radius {
				switch corner {
				case "tl": // 左上 (top-left)
					dst.Set(x+xOffset, y+yOffset, color.RGBA{0, 0, 0, 0})
				case "tr": // 右上 (top-right)
					dst.Set(xOffset+radius-x-1, y+yOffset, color.RGBA{0, 0, 0, 0})
				case "bl": // 左下 (bottom-left)
					dst.Set(x+xOffset, yOffset+radius-y-1, color.RGBA{0, 0, 0, 0})
				case "br": // 右下 (bottom-right)
					dst.Set(xOffset+radius-x-1, yOffset+radius-y-1, color.RGBA{0, 0, 0, 0})
				}
			}
		}
	}
}

func round(img image.Image) image.Image {
	radius := 14
	roundedImg := roundCornersOnly(img, radius)

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

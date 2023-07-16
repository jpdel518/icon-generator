package service

import (
	"bytes"
	"crypto/sha1"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

var _ Service = (*IconService)(nil)

type IconService struct {
}

func NewIconService() *IconService {
	return &IconService{}
}

func (s IconService) Generate(letter string, size int) ([]byte, error) {
	// pngの文字アイコンの生成

	// colorの生成
	_a := sha1.Sum([]byte(letter))
	hash := _a[:sha1.Size]
	rgba := color.RGBA{hash[0], hash[1], hash[2], 0xff}

	// ベースイメージ（矩形）の作成
	img := image.NewNRGBA(image.Rect(0, 0, size, size))

	// 円形の描写
	r := float64(size / 2)
	var c color.RGBA
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			dx, dy := r-float64(x), r-float64(y)
			d := math.Sqrt(dx*dx+dy*dy) / r
			if d > 1 {
				c = color.RGBA{0, 0, 0, 0}
			} else {
				c = rgba
			}
			img.Set(x, y, c)
		}
	}

	// 文字の描写
	// フォントの読み込み
	// ft, err := truetype.Parse(gobold.TTF)
	// 日本語フォントに対応するためにフォントのロード
	ftBinary, err := os.ReadFile("BIZUDGothic-Bold.ttf")
	ft, err := truetype.Parse(ftBinary)
	if err != nil {
		log.Printf("failed to load font: %v", err)
		return nil, err
	}
	opt := truetype.Options{
		Size:              r,
		DPI:               0,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}
	face := truetype.NewFace(ft, &opt)
	dr := &font.Drawer{
		Dst:  img,
		Src:  image.White,
		Face: face,
	}
	dr.Dot.X = (fixed.I(size) - dr.MeasureString(letter)) / 2
	dr.Dot.Y = (fixed.I(size)+fixed.I(size/2))/2 - fixed.I(4)

	dr.DrawString(letter)

	// 出力
	// file, err := os.Create("icon.png")
	// if err != nil {
	// 	log.Printf("failed to create file: %v", err)
	// 	return nil, err
	// }
	// defer file.Close()
	// err = png.Encode(file, img)
	// if err != nil {
	// 	log.Printf("failed to encode png: %v", err)
	// 	return nil, err
	// }
	//
	// output, err := os.ReadFile(file.Name())
	// if err != nil {
	// 	log.Printf("failed to read file: %v", err)
	// 	return nil, err
	// }
	buf := &bytes.Buffer{}
	err = png.Encode(buf, img)
	if err != nil {
		log.Printf("failed to encode png: %v", err)
		return nil, err
	}
	output := buf.Bytes()

	log.Printf("success to create png icon")
	return output, nil
}

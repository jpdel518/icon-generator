package service

import (
	"bytes"
	"github.com/mskrha/svg2png"
	"log"
	"os"
)

var _ Service = (*IconService)(nil)

type IconService struct {
}

func NewIconService() *IconService {
	return &IconService{}
}

func (s IconService) Generate(letter string, size int) ([]byte, error) {
	// pngの文字アイコンの生成（svgを生成してからpngへの変換）

	var buf bytes.Buffer
	// temp, err := os.CreateTemp(os.TempDir(), "tmp_icon")
	// if err != nil {
	// 	log.Fatalf("failed to create temp file: %v", err)
	// 	return err
	// }

	radius := size
	canvas := svg.New(&buf)
	canvas.Start(size, size)
	canvas.Circle(size/2, size/2, radius, "fill:rgb(77,177,135)")
	canvas.Text(size/2, size/2, letter, "text-anchor:middle;font-size:30px;fill:white")
	canvas.End()

	log.Println(buf.String())
	input := make([]byte, buf.Len())
	_, err := buf.Read(input)
	if err != nil {
		log.Printf("failed to read byte buffer: %v", err)
		return nil, err
	}

	// FIXME svg2pngは内部的にinkscapeを使用しているが、安定バージョンのinkscapeでは動かない。メンテナンスもされていない。
	converter := svg2png.New()
	output, err := converter.Convert(input)
	if err != nil {
		log.Printf("failed to convert svg to png: %v", err)
		return nil, err
	}
	file, err := os.Create("icon.png")
	if err != nil {
		log.Printf("failed to create file: %v", err)
		return nil, err
	}
	size, err := file.Write(output)
	if err != nil {
		log.Printf("failed to write icon to file: %v", err)
		return nil, err
	}

	log.Printf("success create icon. size: %v", size)
	return output, nil
}

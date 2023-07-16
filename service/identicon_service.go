package service

import (
	"bytes"
	"github.com/nullrocks/identicon"
	"log"
)

var _ Service = (*IdenticonService)(nil)

type IdenticonService struct {
}

func NewIdenticonService() *IdenticonService {
	return &IdenticonService{}
}

func (i IdenticonService) Generate(letter string, size int) ([]byte, error) {
	// identiconライブラリを使用したidenticonの生成

	// identiconインスタンスの生成
	generator, err := identicon.New(
		"icon_generator",
		10,
		10,
	)
	if err != nil {
		log.Printf("failed to generate identicon generator: %v", err)
		return nil, err
	}

	// 文字をシードにしてidenticonの描写
	ii, err := generator.Draw(letter)
	if err != nil {
		log.Printf("failed to draw identicon: %v", err)
		return nil, err
	}

	// 出力
	// file, err := os.Create("icon.png")
	// if err != nil {
	// 	log.Printf("failed to create file: %v", err)
	// 	return nil, err
	// }
	// defer file.Close()
	//
	// pix := width
	// if height > width {
	// 	pix = height
	// }
	// err = ii.Png(pix, file)
	// if err != nil {
	// 	log.Printf("failed to output to png: %v", err)
	// 	return nil, err
	// }
	//
	// output, err := os.ReadFile(file.Name())
	// if err != nil {
	// 	log.Printf("failed to read file: %v", err)
	// 	return nil, err
	// }

	buf := &bytes.Buffer{}
	err = ii.Png(size, buf)
	if err != nil {
		log.Printf("failed to output to png: %v", err)
		return nil, err
	}
	output := buf.Bytes()

	log.Printf("success to create identicon")
	return output, nil
}

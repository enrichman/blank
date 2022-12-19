package square

import (
	"image"
	"image/color"
	"io"
	"strconv"
	"strings"
)

const (
	Name  string = `square`
	Magic string = `square`
)

func init() {
	image.RegisterFormat(Name, Magic, decode, decodeConfig)
}

var _ image.Image = (*SquareImage)(nil)

type SquareImage struct {
	dimension int
	color     string
}

func (s *SquareImage) At(x, y int) color.Color {
	switch s.color {
	case "red":
		return color.RGBA{R: 255}
	case "green":
		return color.RGBA{G: 255}
	case "blue":
		return color.RGBA{B: 255}
	}
	return color.RGBA{A: 255}
}

func (s *SquareImage) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{},
		Max: image.Point{X: s.dimension, Y: s.dimension},
	}
}

func (s *SquareImage) ColorModel() color.Model {
	return &SquareColorModel{}
}

var _ color.Model = (*SquareColorModel)(nil)

type SquareColorModel struct{}

func (s *SquareColorModel) Convert(c color.Color) color.Color {
	return c
}

func decode(reader io.Reader) (image.Image, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	squareImageArr := strings.Split(string(b), ";")

	dim, err := strconv.Atoi(squareImageArr[1])
	if err != nil {
		return nil, err
	}

	return &SquareImage{
		color:     squareImageArr[2],
		dimension: dim,
	}, nil
}

func decodeConfig(reader io.Reader) (image.Config, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return image.Config{}, err
	}

	squareImageArr := strings.Split(string(b), ";")

	dim, err := strconv.Atoi(squareImageArr[1])
	if err != nil {
		return image.Config{}, err
	}

	return image.Config{
		ColorModel: &SquareColorModel{},
		Width:      dim,
		Height:     dim,
	}, nil
}

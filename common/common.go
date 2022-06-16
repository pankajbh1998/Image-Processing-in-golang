package common

import (
	"image"
	"image/color"
)


type Pixel struct {
	R,G,B,A uint32
}

func GetPixels(img image.Image) [][]Pixel {
	rec := img.Bounds()
	n := rec.Max.X
	m := rec.Max.Y
	matrix := MakeArray(n,m)
	for i:=0;i<n;i++ {
		for j:=0;j<m;j++ {
			matrix[i][j] = RGBAToPixel(img.At(i,j).RGBA())
		}

	}
	return matrix
}

func GetImage(matrix [][]Pixel,)image.Image{
	n,m := GetSize(matrix)
	img := image.NewRGBA(image.Rect(0,0,n,m))
	for i:=0;i<n;i++{
		for j:=0;j<m;j++ {
			img.Set(i,j,PixelTORGBA(matrix[i][j]))
		}
	}
	return img
}

func RGBAToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{r,g,b,a}
}

func PixelTORGBA(pixel Pixel) color.RGBA {
	return color.RGBA{uint8(pixel.R), uint8(pixel.G), uint8(pixel.B), uint8(pixel.A)}
}

func MakeArray(n int ,m int )[][]Pixel{
	ret := make([][]Pixel,n)
	for i:=0;i<n;i++ {
		ret[i]=make([]Pixel,m)
	}
	return ret
}

func GetSize(matrix [][]Pixel)(int, int){
	n := len(matrix)
	m := len(matrix[0])
	return n,m
}

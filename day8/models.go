package day8

import "fmt"

type Image struct {
	Layers []*Layer
	Rows int
	Cols int
	Decoded [][]int
}

type Layer struct {
	Rows []*ImageRow
}

type ImageRow struct {
	Pixels []int
}

func (layer *Layer) Pixels() (res []int) {
	for _, row := range layer.Rows {
		for _, p := range row.Pixels {
			res = append(res, p)
		}
	}
	return res
}

func (image *Image) ParseLayer(layer []int) {
	if len(layer) != image.Rows * image.Cols {
		fmt.Println("Error: invalid layer length")
		return
	}
	layerObj := &Layer{}
	for r := 0; r < len(layer); r += image.Cols {
		layerObj.Rows = append(layerObj.Rows, &ImageRow{Pixels:layer[r: r+image.Cols]})
	}
	image.Layers = append(image.Layers, layerObj)
}


func (image *Image) InitFromIntSlice(intSlice []int) {
	layers := [][]int{}
	for l := 0; l < len(intSlice); l += image.Cols * image.Rows {
		layers = append(layers, intSlice[l:l+image.Cols*image.Rows])
	}
	for _, la := range layers {
		image.ParseLayer(la)
	}
}

func (image *Image) FewestCountOfValueLayer(val int) (res *Layer) {
	lowestCount := 9999999
	for _, l := range image.Layers {
		m := IntSliceValueCountMap(l.Pixels())
		if m[val] < lowestCount {
			lowestCount = m[val]
			res = l
		}
	}
	return res
}

func (image *Image) DecodeThisBitch()  {
	res := [][]int{}
	for _, r := range image.Layers[len(image.Layers) - 1].Rows {
		res = append(res, r.Pixels)
	}
	for i := len(image.Layers) - 1; i >= 0; i-- {
		for t, r := range image.Layers[i].Rows {
			for pI, p := range r.Pixels {
				if p != 2 {
					res[t][pI] = p
				}
			}
		}
	}
	image.Decoded = res
}

func (image *Image) PrintThisBitch() {
	lines := [][]string{}
	for i, d := range image.Decoded {
		for _, p := range d {
			switch p {
			case 0:
				lines[i] = append(lines[i], " ")
			}

		}
	}
}
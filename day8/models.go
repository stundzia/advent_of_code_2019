package day8

import "fmt"

type Image struct {
	Layers []*Layer
	Rows int
	Cols int
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
	//fmt.Println(image.Layers)
	//for _, lay := range image.Layers {
	//	fmt.Println(lay.Rows[0])
	//}
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
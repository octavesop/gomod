package main

import (
	"fmt"
	"usepkg/cuspkg" //모듈 내 패키지 custompkg
	"github.com/guptarohit/asciigraph"           //외부 저장소 패키지 asciigraph
)

func main() {
	custompkg.PrintCustom()

	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}
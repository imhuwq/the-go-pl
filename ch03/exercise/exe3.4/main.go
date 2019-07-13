package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func corner(i, j int) (float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0.0, 0.0, false
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func genSVG(rsp http.ResponseWriter, req *http.Request) {
	rsp.Header().Set("Content-Type", "image/svg+xml")
	_, _ = fmt.Fprintf(rsp, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ao := corner(i+1, j)
			bx, by, bo := corner(i, j)
			cx, cy, co := corner(i, j+1)
			dx, dy, do := corner(i+1, j+1)
			if !(ao && bo && co && do) {
				continue
			}
			_, _ = fmt.Fprintf(rsp, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	_, _ = fmt.Fprintf(rsp, "</svg>")
}

func main() {
	http.HandleFunc("/", genSVG)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

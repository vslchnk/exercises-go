package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", surf)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func surf(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	width, err := strconv.Atoi(r.FormValue("width"))
	if err != nil {
		width = 600
	}

	height, err := strconv.Atoi(r.FormValue("height"))
	if err != nil {
		height = 320
	}

	var color string = r.FormValue("color")
	if len(color) == 0 {
		color = "grey"
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, width, height, color)
}

func surface(out io.Writer, width, height int, color string) {
	var s string
	s = fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", color, width, height)
	out.Write([]byte(s))

	xyscale := float64(width) / 2 / xyrange // pixels per x or y unit
	zscale := float64(height) * 0.4

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, width, height, xyscale, zscale)
			bx, by := corner(i, j, width, height, xyscale, zscale)
			cx, cy := corner(i, j+1, width, height, xyscale, zscale)
			dx, dy := corner(i+1, j+1, width, height, xyscale, zscale)
			s = fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
			out.Write([]byte(s))
		}
	}
	s = fmt.Sprintln("</svg>")
	out.Write([]byte(s))
}

func corner(i, j, width, height int, xyscale, zscale float64) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {

	return -math.Abs(0.25*math.Cos(0.5*x) + 0.25*math.Cos(0.5*y))
}

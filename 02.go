package gostudy

type Point struct {
	X, Y int
}

type Rectangle struct {
	Min, Max Point
}

func Rect(x0, y0, x1, y1 int) Rectangle {
	return Rectangle{Point{x0, y0}, Point{x1, y1}}
}

package vector

type Vector2 struct {
	x, y float64
}

func MakeVector2(x, y float64) *Vector2 {
	return &Vector2{
		x: x,
		y: y,
	}
}

func (v *Vector2) X() float64 {
	return v.x
}

func (v *Vector2) Y() float64 {
	return v.y
}

func (v *Vector2) SetX(newX float64) {
	v.x = newX
}

func (v *Vector2) SetY(newY float64) {
	v.y = newY
}
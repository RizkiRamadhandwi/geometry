package shape

type Ractangle struct {
	Width  float32
	Height float32
}

func (r *Ractangle) Luas() float32 {
	return r.Width * r.Height
}

func (r *Ractangle) Keliling() float32 {
	return 2 * (r.Width + r.Height)
}

package models

type Box struct {
	Length, Width, Height float64
}

func (b *Box) CalVolume() float64 {
	return b.Height * b.Length * b.Width
}
func (b *Box) CalSurface() float64 {
	return (2 * b.Length * b.Width) + (2 * b.Length * b.Height) + (2 * b.Width * b.Height)
}

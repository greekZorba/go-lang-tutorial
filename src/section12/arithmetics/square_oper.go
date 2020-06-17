package arithmetics

func (o *Numbers) SquarePlus() int {
	return (o.X * o.X) + (o.Y * o.Y)
}

func (o *Numbers) SquareMinus() int {
	return (o.X * o.X) - (o.Y * o.Y)
}

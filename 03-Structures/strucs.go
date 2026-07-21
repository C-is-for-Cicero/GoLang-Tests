package main

type Rectangle struct {
	width  float64
	height float64
}

type Counter struct {
	count int
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// A pointer receiver is used when the method needs to modify the receiver or when the receiver is a large struct and you want to avoid copying it.
func (c *Counter) Increment() {
	c.count++
}

func (c *Counter) Value() int {
	return c.count
}

package structsmi

import "math"

type Rectangle struct {
  width float64
  height float64
}

type Circle struct {
  radius float64
}

type Triangle struct {
  base float64
  height float64
}

type Shape interface {
  Area() float64
}

func (t Triangle) Area() float64 {
/*   result := 0.5 * (t.base * t.height) */
  return (0)
}

func (r Rectangle) Area() float64 {
  result := r.height * r.width
  return (result)
}

func (circle Circle) Area() float64 {
  result := math.Pi * (circle.radius * circle.radius)
  return (result)
}

func Perimeter(rectangle Rectangle) float64 {
  result := 2 *(rectangle.height + rectangle.width)

  return (result)
}

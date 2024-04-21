package vec

import (
	"fmt"
)

type Vec2 struct {
	X int
	Y int
}

func (v *Vec2) Add(a *Vec2) *Vec2 {
	return &Vec2{X: v.X + a.X, Y: v.Y + a.Y}
}

func (v *Vec2) IAdd(a *Vec2) {
	v.X += a.X
	v.Y += a.Y
}

func PrintVec(vec Vec2) {
	fmt.Printf("x=%d y=%d\n", vec.X, vec.Y)
}

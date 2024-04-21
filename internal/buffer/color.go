package buffer

import (
	"fmt"

	"github.com/jmmalqui/terminal-renderer/internal/vec"
)

type Color struct {
	R int
	G int
	B int
}

func (c *Color) ColorToString() string {
	// fmt.Println(c)
	return fmt.Sprintf("\033[48;2;%v;%v;%vm ", c.R, c.G, c.B)
}

func RGB256ToTrueColor216(color vec.Vec3) Color {
	var normalizedVec3 vec.Vec3
	{
		normalizedVec3 = *color.Normalize()
	}
	normalizedVec3.IMult(216)
	var trueColor Color
	{
		trueColor.R = int(normalizedVec3.X)
		trueColor.G = int(normalizedVec3.Y)
		trueColor.B = int(normalizedVec3.Z)
	}
	return trueColor
}
func Vec3ToColor(color vec.Vec3) Color {
	var c Color
	{
		c.R = int(color.X)
		c.G = int(color.Y)
		c.B = int(color.Z)

	}
	return c
}

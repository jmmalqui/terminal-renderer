package color

import (
	"fmt"

	"github.com/jmmalqui/terminal-renderer/internal/vec"
)

type Color struct {
	R int
	G int
	B int
}

func (c *Color) Diff(t Color) Color {
	var diff Color
	{
		diff.R = c.R - t.R
		diff.G = c.G - t.G
		diff.B = c.B - t.B
	}
	return diff
}
func (c *Color) Add(t Color) Color {
	var diff Color
	{
		diff.R = c.R + t.R
		diff.G = c.G + t.G
		diff.B = c.B + t.B
	}
	return diff
}
func (c *Color) Mult(t float64) Color {
	var diff Color
	{
		diff.R = int(float64(c.R) * t)
		diff.G = int(float64(c.G) * t)
		diff.B = int(float64(c.B) * t)
	}
	return diff
}
func (c *Color) ColorToString() string {
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

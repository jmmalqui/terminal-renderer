package main

import (
	"github.com/jmmalqui/terminal-renderer/internal/color"
	"github.com/jmmalqui/terminal-renderer/internal/geometry"
	"github.com/jmmalqui/terminal-renderer/internal/gfxdraw"
	"github.com/jmmalqui/terminal-renderer/internal/terminal"
	"github.com/jmmalqui/terminal-renderer/internal/vec"
)

func main() {
	term := terminal.MakeTerminal()
	run := true
	y := 70
	frame := 0
	if run {
		for {
			frame += 1
			// y += int(math.Sin(float64(frame)) * 5)
			term.Fill(vec.Vec3{X: 0, Y: 0, Z: 0})
			gfxdraw.DrawTriangle(
				&term,
				gfxdraw.Triangle{
					Node1: geometry.Node{Color: color.Color{R: 0, G: 0, B: 255}, Pos: vec.Vec2{X: 88, Y: 5}},
					Node2: geometry.Node{Color: color.Color{R: 255, G: 0, B: 0}, Pos: vec.Vec2{X: 145, Y: 25}},
					Node3: geometry.Node{Color: color.Color{R: 0, G: 255, B: 0}, Pos: vec.Vec2{X: 10, Y: y}},
				},
			)

			gfxdraw.DrawTriangle(
				&term,
				gfxdraw.Triangle{
					Node1: geometry.Node{Color: color.Color{R: 0, G: 255, B: 0}, Pos: vec.Vec2{X: 10, Y: y}},
					Node2: geometry.Node{Color: color.Color{R: 255, G: 0, B: 0}, Pos: vec.Vec2{X: 145, Y: 25}},
					Node3: geometry.Node{Color: color.Color{R: 0, G: 0, B: 255}, Pos: vec.Vec2{X: 160, Y: 45}},
				},
			)
			gfxdraw.DrawTriangle(
				&term,
				gfxdraw.Triangle{
					Node1: geometry.Node{Color: color.Color{R: 0, G: 0, B: 255}, Pos: vec.Vec2{X: 160, Y: 45}},
					Node2: geometry.Node{Color: color.Color{R: 255, G: 0, B: 0}, Pos: vec.Vec2{X: 145, Y: 25}},
					Node3: geometry.Node{Color: color.Color{R: 0, G: 255, B: 0}, Pos: vec.Vec2{X: 88, Y: 5}},
				},
			)
			term.Flip()
		}
	}
}

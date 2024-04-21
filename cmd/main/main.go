package main

import (
	"github.com/jmmalqui/terminal-renderer/internal/terminal"
	"github.com/jmmalqui/terminal-renderer/internal/vec"
)

func main() {
	term := terminal.MakeTerminal()
	for {
		term.Fill(vec.Vec3{X: 0, Y: 255, Z: 0})
		term.Flip()
	}
}

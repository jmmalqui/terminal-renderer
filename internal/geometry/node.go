package geometry

import (
	"github.com/jmmalqui/terminal-renderer/internal/color"
	"github.com/jmmalqui/terminal-renderer/internal/vec"
)

type Node struct {
	Color color.Color
	Pos   vec.Vec2
}

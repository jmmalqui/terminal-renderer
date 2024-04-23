package gfxdraw

import (
	"github.com/jmmalqui/terminal-renderer/internal/color"
	"github.com/jmmalqui/terminal-renderer/internal/geometry"
	"github.com/jmmalqui/terminal-renderer/internal/terminal"
)

type Triangle struct {
	Node1 geometry.Node
	Node2 geometry.Node
	Node3 geometry.Node
}

type Span struct {
	Color1 color.Color
	Color2 color.Color
	X1     int
	X2     int
}

func MakeSpan(color1 color.Color, x1 int, color2 color.Color, x2 int) Span {
	var span Span
	{
		if x1 < x2 {
			span.Color1 = color1
			span.X1 = x1
			span.Color2 = color2
			span.X2 = x2
		} else {
			span.Color1 = color2
			span.X1 = x2
			span.Color2 = color1
			span.X2 = x1
		}
	}
	return span
}

func DrawTriangle2D(terminal *terminal.Terminal, vertex1 []float64, color1 []int, vertex2 []float64, color2 []int, vertex3 []float64, color3 []int) {

}

func DrawTriangle(terminal *terminal.Terminal, triangle Triangle) {
	var edges []geometry.Edge
	{
		edges = append(edges, *geometry.MakeEdge(triangle.Node1, triangle.Node2))
		edges = append(edges, *geometry.MakeEdge(triangle.Node2, triangle.Node3))
		edges = append(edges, *geometry.MakeEdge(triangle.Node3, triangle.Node1))
	}

	maxlength := 0
	longEdge := 0

	for idx, edge := range edges {
		length := edge.Node2.Pos.Y - edge.Node1.Pos.Y
		if length > maxlength {
			maxlength = length
			longEdge = idx
		}
	}
	shortEdge1 := (longEdge + 1) % 3
	shortEdge2 := (longEdge + 2) % 3

	DrawSpansBetweenEdges(terminal, edges[longEdge], edges[shortEdge1])
	DrawSpansBetweenEdges(terminal, edges[longEdge], edges[shortEdge2])
	terminal.DrawPixel(triangle.Node1.Pos.X, triangle.Node1.Pos.Y, triangle.Node1.Color)
	terminal.DrawPixel(triangle.Node2.Pos.X, triangle.Node2.Pos.Y, triangle.Node2.Color)
	terminal.DrawPixel(triangle.Node3.Pos.X, triangle.Node3.Pos.Y, triangle.Node3.Color)
}

func DrawSpan(terminal *terminal.Terminal, span Span, y int) {
	xdiff := span.X2 - span.X1
	if xdiff == 0 {
		return
	}

	colorDiff := span.Color2.Diff(span.Color1)
	factor := 0.0
	factorStep := 1.0 / float64(xdiff)
	for x := span.X1; x < span.X2; x++ {
		color := span.Color1.Add(colorDiff.Mult(factor))
		terminal.DrawPixel(x, y, color)
		factor += factorStep
	}

}

func DrawSpansBetweenEdges(terminal *terminal.Terminal, edge1 geometry.Edge, edge2 geometry.Edge) {
	yDiff1 := float64(edge1.Node2.Pos.Y - edge1.Node1.Pos.Y)
	if yDiff1 == 0.0 {
		return
	}
	yDiff2 := float64(edge2.Node2.Pos.Y - edge2.Node1.Pos.Y)
	if yDiff2 == 0.0 {
		return
	}

	xDiff1 := edge1.Node2.Pos.X - edge1.Node1.Pos.X
	xDiff2 := edge2.Node2.Pos.X - edge2.Node1.Pos.X
	colorDiff1 := edge1.Node2.Color.Diff(edge1.Node1.Color)
	colorDiff2 := edge2.Node2.Color.Diff(edge2.Node1.Color)

	factor1 := float64(edge2.Node1.Pos.Y-edge1.Node1.Pos.Y) / yDiff1
	factorStep1 := 1.0 / float64(yDiff1)
	factor2 := 0.0
	factorStep2 := 1.0 / float64(yDiff2)

	for y := edge2.Node1.Pos.Y; y < edge2.Node2.Pos.Y; y++ {

		spanColor1 := edge1.Node1.Color.Add(colorDiff1.Mult(factor1))

		spanX1 := float64(edge1.Node1.Pos.X) + factor1*float64(xDiff1)

		spanColor2 := edge2.Node1.Color.Add(colorDiff2.Mult(factor2))

		spanX2 := float64(edge2.Node1.Pos.X) + factor2*float64(xDiff2)
		span := Span{
			spanColor1,
			spanColor2,
			int(spanX1),
			int(spanX2),
		}

		DrawSpan(terminal, span, y)
		factor1 += factorStep1
		factor2 += factorStep2
	}

}

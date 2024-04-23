package geometry

type Edge struct {
	Node1 Node
	Node2 Node
}

func MakeEdge(node1 Node, node2 Node) *Edge {
	var edge Edge
	{
		if node1.Pos.Y < node2.Pos.Y {
			edge.Node1.Color = node1.Color
			edge.Node1.Pos.X = node1.Pos.X
			edge.Node1.Pos.Y = node1.Pos.Y

			edge.Node2.Color = node2.Color
			edge.Node2.Pos.X = node2.Pos.X
			edge.Node2.Pos.Y = node2.Pos.Y
		} else {
			edge.Node1.Color = node2.Color
			edge.Node1.Pos.X = node2.Pos.X
			edge.Node1.Pos.Y = node2.Pos.Y

			edge.Node2.Color = node1.Color
			edge.Node2.Pos.X = node1.Pos.X
			edge.Node2.Pos.Y = node1.Pos.Y
		}
	}
	return &edge
}

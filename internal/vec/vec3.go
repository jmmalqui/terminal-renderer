package vec

import "math"

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (v3 *Vec3) Normalize() *Vec3 {
	var V float64
	{
		V = float64(v3.X*v3.X + v3.Y*v3.Y + v3.Z*v3.Z)
		V = math.Sqrt(V)
	}
	return &Vec3{v3.X / V, v3.Y / V, v3.Z / V}
}

func (v *Vec3) IMult(a float64) {
	v.X *= a
	v.Y *= a
	v.Z *= a
}

func (v *Vec3) IAdd(a *Vec3) {
	v.X += a.X
	v.Y += a.Y
	v.Z += a.Z
}

package main

type M struct {
	data []float64
	r, c int
}

func (m *M) Dims() (int, int) { return m.r, m.c }
func (m *M) At(i, j int) float64 {
	return m.data[i*m.c+j]
}

func main() {
	m := M{
		data: []float64{1, 2, 3, 4},
		r:    2,
		c:    2,
	}
	_ = m
}
